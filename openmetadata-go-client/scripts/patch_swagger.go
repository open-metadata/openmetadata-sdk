// patch_swagger.go patches the OpenMetadata swagger.json before Go code generation.
//
// oapi-codegen maps {"type": "object"} (bare object with no properties) to
// map[string]interface{}, but the OpenMetadata API returns arbitrary JSON
// values (strings, bools, arrays, numbers) for many of these fields.
//
// This script removes the "type": "object" constraint from such fields so
// oapi-codegen generates interface{} (any) instead, which correctly accepts
// all JSON value types.
//
// Usage:
//
//	go run scripts/patch_swagger.go <input.json> <output.json>
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input.json> <output.json>\n", os.Args[0])
		os.Exit(1)
	}

	inputPath, outputPath := os.Args[1], os.Args[2]

	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read %s: %v\n", inputPath, err)
		os.Exit(1)
	}

	var spec map[string]interface{}
	if err := json.Unmarshal(data, &spec); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse JSON: %v\n", err)
		os.Exit(1)
	}

	schemas := findSchemas(spec)
	if schemas == nil {
		fmt.Fprintln(os.Stderr, "no definitions or components/schemas found in spec")
		os.Exit(1)
	}

	count := patchBareObjects(schemas)
	fmt.Printf("Patched %d bare-object fields in swagger spec\n", count)

	out, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal JSON: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write %s: %v\n", outputPath, err)
		os.Exit(1)
	}
}

// findSchemas returns the schema definitions map from either
// OpenAPI 2.x ("definitions") or 3.x ("components.schemas").
func findSchemas(spec map[string]interface{}) map[string]interface{} {
	if defs, ok := spec["definitions"].(map[string]interface{}); ok {
		return defs
	}
	if components, ok := spec["components"].(map[string]interface{}); ok {
		if schemas, ok := components["schemas"].(map[string]interface{}); ok {
			return schemas
		}
	}
	return nil
}

// isBareObject returns true if the value is {"type": "object"} with no other keys.
func isBareObject(v interface{}) bool {
	m, ok := v.(map[string]interface{})
	if !ok || len(m) != 1 {
		return false
	}
	t, ok := m["type"].(string)
	return ok && t == "object"
}

// patchBareObjects removes "type": "object" from properties that are bare objects
// (no "properties", "$ref", "additionalProperties", etc.). These represent free-form
// values in the OpenMetadata API.
//
// For array properties whose items are bare objects, the items are patched the same way.
func patchBareObjects(schemas map[string]interface{}) int {
	count := 0

	for _, schemaDef := range schemas {
		schemaMap, ok := schemaDef.(map[string]interface{})
		if !ok {
			continue
		}
		properties, ok := schemaMap["properties"].(map[string]interface{})
		if !ok {
			continue
		}

		for propName, propValue := range properties {
			propMap, ok := propValue.(map[string]interface{})
			if !ok {
				continue
			}

			// Case 1: property is exactly {"type": "object"}
			if isBareObject(propValue) {
				properties[propName] = map[string]interface{}{}
				count++
				continue
			}

			// Case 2: property is {"type": "array", "items": {"type": "object"}}
			if t, _ := propMap["type"].(string); t == "array" {
				if isBareObject(propMap["items"]) {
					propMap["items"] = map[string]interface{}{}
					count++
				}
			}
		}
	}

	return count
}
