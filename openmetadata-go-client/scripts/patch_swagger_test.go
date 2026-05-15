package main

import (
	"encoding/json"
	"testing"
)

func parseSchemas(t *testing.T, raw string) map[string]interface{} {
	t.Helper()
	var spec map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &spec); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	schemas := findSchemas(spec)
	if schemas == nil {
		t.Fatal("findSchemas returned nil")
	}
	return schemas
}

func getNestedString(t *testing.T, m map[string]interface{}, keys ...string) string {
	t.Helper()
	var cur interface{} = m
	for _, k := range keys {
		mm, ok := cur.(map[string]interface{})
		if !ok {
			t.Fatalf("expected map at key %q, got %T", k, cur)
		}
		cur = mm[k]
	}
	s, ok := cur.(string)
	if !ok {
		t.Fatalf("expected string at %v, got %T (%v)", keys, cur, cur)
	}
	return s
}

func TestPatchDateTimeToEpochMillis_FlatProperty(t *testing.T) {
	schemas := parseSchemas(t, `{
		"definitions": {
			"Entity": {
				"properties": {
					"updatedAt": {"type": "string", "format": "date-time"}
				}
			}
		}
	}`)

	n := patchDateTimeToEpochMillis(schemas)
	if n != 1 {
		t.Fatalf("want 1 patch, got %d", n)
	}
	if got := getNestedString(t, schemas, "Entity", "properties", "updatedAt", "type"); got != "integer" {
		t.Errorf("type: want integer, got %s", got)
	}
	if got := getNestedString(t, schemas, "Entity", "properties", "updatedAt", "format"); got != "int64" {
		t.Errorf("format: want int64, got %s", got)
	}
}

func TestPatchDateTimeToEpochMillis_NestedInlineObject(t *testing.T) {
	schemas := parseSchemas(t, `{
		"definitions": {
			"Entity": {
				"properties": {
					"schedule": {
						"type": "object",
						"properties": {
							"startDate": {"type": "string", "format": "date-time"},
							"endDate":   {"type": "string", "format": "date-time"}
						}
					}
				}
			}
		}
	}`)

	n := patchDateTimeToEpochMillis(schemas)
	if n != 2 {
		t.Fatalf("want 2 patches, got %d", n)
	}
	for _, field := range []string{"startDate", "endDate"} {
		if got := getNestedString(t, schemas, "Entity", "properties", "schedule", "properties", field, "type"); got != "integer" {
			t.Errorf("%s type: want integer, got %s", field, got)
		}
		if got := getNestedString(t, schemas, "Entity", "properties", "schedule", "properties", field, "format"); got != "int64" {
			t.Errorf("%s format: want int64, got %s", field, got)
		}
	}
}

func TestPatchDateTimeToEpochMillis_InsideAllOf(t *testing.T) {
	schemas := parseSchemas(t, `{
		"definitions": {
			"Entity": {
				"allOf": [
					{
						"properties": {
							"createdAt": {"type": "string", "format": "date-time"}
						}
					}
				]
			}
		}
	}`)

	n := patchDateTimeToEpochMillis(schemas)
	if n != 1 {
		t.Fatalf("want 1 patch, got %d", n)
	}
	allOf := schemas["Entity"].(map[string]interface{})["allOf"].([]interface{})
	entry := allOf[0].(map[string]interface{})["properties"].(map[string]interface{})["createdAt"].(map[string]interface{})
	if got := entry["type"].(string); got != "integer" {
		t.Errorf("type: want integer, got %s", got)
	}
	if got := entry["format"].(string); got != "int64" {
		t.Errorf("format: want int64, got %s", got)
	}
}

func TestPatchBareObjects_FlatProperty(t *testing.T) {
	schemas := parseSchemas(t, `{
		"definitions": {
			"Entity": {
				"properties": {
					"config": {"type": "object"}
				}
			}
		}
	}`)

	n := patchBareObjects(schemas)
	if n != 1 {
		t.Fatalf("want 1 patch, got %d", n)
	}
	prop := schemas["Entity"].(map[string]interface{})["properties"].(map[string]interface{})["config"].(map[string]interface{})
	if len(prop) != 0 {
		t.Errorf("want empty map, got %v", prop)
	}
}

func TestPatchBareObjects_ArrayItems(t *testing.T) {
	schemas := parseSchemas(t, `{
		"definitions": {
			"Entity": {
				"properties": {
					"tags": {"type": "array", "items": {"type": "object"}}
				}
			}
		}
	}`)

	n := patchBareObjects(schemas)
	if n != 1 {
		t.Fatalf("want 1 patch, got %d", n)
	}
	items := schemas["Entity"].(map[string]interface{})["properties"].(map[string]interface{})["tags"].(map[string]interface{})["items"].(map[string]interface{})
	if len(items) != 0 {
		t.Errorf("want empty items map, got %v", items)
	}
}

func TestPatchBareObjects_InsideAllOf(t *testing.T) {
	schemas := parseSchemas(t, `{
		"definitions": {
			"Entity": {
				"allOf": [
					{
						"properties": {
							"meta": {"type": "object"}
						}
					}
				]
			}
		}
	}`)

	n := patchBareObjects(schemas)
	if n != 1 {
		t.Fatalf("want 1 patch, got %d", n)
	}
	allOf := schemas["Entity"].(map[string]interface{})["allOf"].([]interface{})
	prop := allOf[0].(map[string]interface{})["properties"].(map[string]interface{})["meta"].(map[string]interface{})
	if len(prop) != 0 {
		t.Errorf("want empty map, got %v", prop)
	}
}

func TestPatchBareObjects_SkipsNonBareObjects(t *testing.T) {
	schemas := parseSchemas(t, `{
		"definitions": {
			"Entity": {
				"properties": {
					"rich": {"type": "object", "properties": {"x": {"type": "string"}}}
				}
			}
		}
	}`)

	n := patchBareObjects(schemas)
	if n != 0 {
		t.Fatalf("want 0 patches (rich object must not be touched), got %d", n)
	}
}
