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

	dtCount := patchDateTimeToEpochMillis(schemas)
	fmt.Printf("Patched %d date-time fields to int64 (epoch millis) in swagger spec\n", dtCount)

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

// walkSchema recursively visits every schema node in the tree, calling fn on
// each map[string]interface{} encountered. This covers properties, allOf/anyOf/oneOf
// entries, array items, and any other inline schema.
func walkSchema(node interface{}, fn func(m map[string]interface{})) {
	switch v := node.(type) {
	case map[string]interface{}:
		fn(v)
		for _, child := range v {
			walkSchema(child, fn)
		}
	case []interface{}:
		for _, elem := range v {
			walkSchema(elem, fn)
		}
	}
}

// patchBareObjects removes "type": "object" from properties that are bare objects
// (no "properties", "$ref", "additionalProperties", etc.). These represent free-form
// values in the OpenMetadata API.
//
// For array properties whose items are bare objects, the items are patched the same way.
func patchBareObjects(schemas map[string]interface{}) int {
	count := 0

	for _, schemaDef := range schemas {
		walkSchema(schemaDef, func(m map[string]interface{}) {
			props, ok := m["properties"].(map[string]interface{})
			if !ok {
				return
			}
			for propName, propValue := range props {
				propMap, ok := propValue.(map[string]interface{})
				if !ok {
					continue
				}

				if isBareObject(propValue) {
					props[propName] = map[string]interface{}{}
					count++
					continue
				}

				if t, _ := propMap["type"].(string); t == "array" {
					if isBareObject(propMap["items"]) {
						propMap["items"] = map[string]interface{}{}
						count++
					}
				}
			}
		})
	}

	return count
}

// patchDateTimeToEpochMillis converts all date-time properties to
// integer/int64 so oapi-codegen generates *int64 instead of *time.Time.
// The OpenMetadata API returns epoch milliseconds for these fields, which
// causes time.Time.UnmarshalJSON to fail.
//
// Walks the entire schema tree so fields nested inside allOf/anyOf/oneOf or
// inline object properties are also patched.
func patchDateTimeToEpochMillis(schemas map[string]interface{}) int {
	count := 0

	for _, schemaDef := range schemas {
		walkSchema(schemaDef, func(m map[string]interface{}) {
			if fmtVal, ok := m["format"].(string); ok && fmtVal == "date-time" {
				m["type"] = "integer"
				m["format"] = "int64"
				count++
			}
		})
	}

	return count
}
