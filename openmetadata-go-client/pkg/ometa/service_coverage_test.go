package ometa

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

var swaggerFilePath = flag.String("swagger-path", "", "path to swagger.json for service coverage validation")

func TestServiceCoverage(t *testing.T) {
	swaggerPath := *swaggerFilePath
	if swaggerPath == "" {
		swaggerPath = os.Getenv("SWAGGER_PATH")
	}
	if swaggerPath == "" {
		t.Skip("skipping: provide -swagger-path flag or SWAGGER_PATH env var")
	}

	swaggerPaths := parseSwaggerPaths(t, swaggerPath)
	serviceBasePaths := scanServiceBasePaths(t)

	// Service → Swagger: every service basePath must appear in the swagger.
	// A service is "found" if at least one swagger path equals or starts with its basePath.
	var staleServices []string
	for basePath := range serviceBasePaths {
		found := false
		for _, sp := range swaggerPaths {
			if sp == basePath || strings.HasPrefix(sp, basePath+"/") {
				found = true
				break
			}
		}
		if !found {
			staleServices = append(staleServices, basePath)
		}
	}
	sort.Strings(staleServices)
	for _, s := range staleServices {
		t.Errorf("service with basePath %q has no matching swagger resource", s)
	}

	// Swagger → Service: check which swagger paths are covered by a service.
	// A swagger path is "covered" if any service basePath equals it or is a prefix of it.
	// Uncovered paths are grouped by root resource and logged (informational only).
	uncoveredRoots := make(map[string]bool)
	for _, sp := range swaggerPaths {
		covered := false
		for basePath := range serviceBasePaths {
			if sp == basePath || strings.HasPrefix(sp, basePath+"/") {
				covered = true
				break
			}
		}
		if !covered {
			root := extractResourceBase(sp)
			uncoveredRoots[root] = true
		}
	}
	var roots []string
	for r := range uncoveredRoots {
		roots = append(roots, r)
	}
	sort.Strings(roots)
	if len(roots) > 0 {
		t.Logf("swagger resource groups without a service (%d): %s", len(roots), strings.Join(roots, ", "))
	}

	t.Logf("swagger paths: %d, service base paths: %d, stale services: %d, uncovered resource groups: %d",
		len(swaggerPaths), len(serviceBasePaths), len(staleServices), len(roots))
}

// parseSwaggerPaths reads a swagger.json and returns all paths with the API
// version prefix stripped (e.g., "/api/v1/tables/{id}" → "tables/{id}").
func parseSwaggerPaths(t *testing.T, path string) []string {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read swagger file: %v", err)
	}

	var spec struct {
		Paths map[string]json.RawMessage `json:"paths"`
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatalf("failed to parse swagger JSON: %v", err)
	}

	var paths []string
	for rawPath := range spec.Paths {
		if stripped := stripVersionPrefix(rawPath); stripped != "" {
			paths = append(paths, stripped)
		}
	}
	sort.Strings(paths)
	return paths
}

// stripVersionPrefix removes the API version prefix from a swagger path.
//
//	/api/v1/tables/{id} → tables/{id}
//	/v1/tables          → tables
func stripVersionPrefix(rawPath string) string {
	path := rawPath
	for i := 0; i < len(path)-3; i++ {
		if path[i] == '/' && path[i+1] == 'v' && path[i+2] >= '0' && path[i+2] <= '9' {
			if end := strings.Index(path[i+1:], "/"); end != -1 {
				return path[i+1+end+1:]
			}
		}
	}
	return strings.TrimPrefix(path, "/")
}

// extractResourceBase derives the base resource path from a version-stripped
// swagger path by collecting segments before the first {param} and stripping
// trailing action keywords. Used only for grouping uncovered paths.
//
//	tables/{id}/versions/{version} → tables
//	tables/async                   → tables
//	services/databaseServices/{id} → services/databaseServices
func extractResourceBase(path string) string {
	segments := strings.Split(path, "/")
	var base []string
	for _, seg := range segments {
		if strings.HasPrefix(seg, "{") {
			break
		}
		base = append(base, seg)
	}
	if len(base) == 0 {
		return ""
	}

	trailingActions := map[string]bool{
		"name": true, "restore": true, "versions": true,
		"export": true, "import": true, "csv": true,
		"async": true, "bulk": true, "history": true,
	}
	for len(base) > 1 && trailingActions[base[len(base)-1]] {
		base = base[:len(base)-1]
	}

	return strings.Join(base, "/")
}

// scanServiceBasePaths reads all service_*.go files in the package directory
// and extracts their basePath constants.
func scanServiceBasePaths(t *testing.T) map[string]bool {
	t.Helper()

	matches, err := filepath.Glob("service_*.go")
	if err != nil {
		t.Fatalf("failed to glob service files: %v", err)
	}

	re := regexp.MustCompile(`const \w+BasePath\s*=\s*"([^"]+)"`)
	result := make(map[string]bool)

	for _, f := range matches {
		name := filepath.Base(f)
		if strings.HasSuffix(name, "_test.go") || name == "service_base.go" {
			continue
		}
		data, err := os.ReadFile(f)
		if err != nil {
			t.Errorf("failed to read %s: %v", f, err)
			continue
		}
		if m := re.FindSubmatch(data); m != nil {
			result[string(m[1])] = true
		}
	}
	return result
}

func TestStripVersionPrefix(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"/api/v1/tables", "tables"},
		{"/api/v1/tables/{id}", "tables/{id}"},
		{"/api/v1/tables/{id}/versions/{version}", "tables/{id}/versions/{version}"},
		{"/api/v1/services/databaseServices/{id}", "services/databaseServices/{id}"},
		{"/v1/tables", "tables"},
		{"/api/v2/tables", "tables"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := stripVersionPrefix(tt.input)
			if got != tt.want {
				t.Errorf("stripVersionPrefix(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestExtractResourceBase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"tables", "tables"},
		{"tables/{id}", "tables"},
		{"tables/{id}/versions/{version}", "tables"},
		{"tables/name/{fqn}", "tables"},
		{"tables/restore", "tables"},
		{"tables/export", "tables"},
		{"tables/async", "tables"},
		{"tables/bulk", "tables"},
		{"tables/history", "tables"},
		{"services/databaseServices", "services/databaseServices"},
		{"services/databaseServices/{id}", "services/databaseServices"},
		{"services/databaseServices/async", "services/databaseServices"},
		{"services/databaseServices/history", "services/databaseServices"},
		{"dataQuality/testCases", "dataQuality/testCases"},
		{"dataQuality/testCases/{id}", "dataQuality/testCases"},
		{"events/subscriptions", "events/subscriptions"},
		{"dashboard/datamodels", "dashboard/datamodels"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := extractResourceBase(tt.input)
			if got != tt.want {
				t.Errorf("extractResourceBase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
