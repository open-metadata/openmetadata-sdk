package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestDatabaseSchema(t *testing.T, ctx context.Context, name string) *ometa.DatabaseSchema {
	t.Helper()

	db := createTestDatabase(t, ctx, name+"_parent_db")

	schema, err := client.DatabaseSchemas.Create(ctx, &ometa.CreateDatabaseSchema{
		Name:     name,
		Database: *db.FullyQualifiedName,
	})
	if err != nil {
		t.Fatalf("failed to create database schema '%s': %v", name, err)
	}

	t.Cleanup(func() {
		client.DatabaseSchemas.DeleteByName(ctx, *schema.FullyQualifiedName, &ometa.DeleteDBSchemaByFQNParams{
			HardDelete: ometa.Bool(true),
			Recursive:  ometa.Bool(true),
		})
	})

	return schema
}

func TestCreateDatabaseSchema(t *testing.T) {
	ctx := context.Background()
	schema := createTestDatabaseSchema(t, ctx, "test_create_schema")

	if schema.Name != "test_create_schema" {
		t.Errorf("expected name 'test_create_schema', got '%s'", schema.Name)
	}
}

func TestGetDatabaseSchemaByID(t *testing.T) {
	ctx := context.Background()
	schema := createTestDatabaseSchema(t, ctx, "test_get_schema_by_id")

	got, err := client.DatabaseSchemas.GetByID(ctx, schema.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get schema by ID: %v", err)
	}

	if got.Id != schema.Id {
		t.Errorf("expected ID '%s', got '%s'", schema.Id, got.Id)
	}
}

func TestGetDatabaseSchemaByName(t *testing.T) {
	ctx := context.Background()
	schema := createTestDatabaseSchema(t, ctx, "test_get_schema_by_name")

	got, err := client.DatabaseSchemas.GetByName(ctx, *schema.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get schema by name: %v", err)
	}

	if got.Id != schema.Id {
		t.Errorf("expected ID '%s', got '%s'", schema.Id, got.Id)
	}
}
