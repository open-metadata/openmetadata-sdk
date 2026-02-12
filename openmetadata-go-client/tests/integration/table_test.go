package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestTable(t *testing.T, ctx context.Context, name string) *ometa.Table {
	t.Helper()

	table, err := client.Tables.Create(ctx, &ometa.CreateTable{
		Name:           name,
		DatabaseSchema: testSchema,
		Columns: []ometa.Column{
			{Name: "id", DataType: "INT"},
			{Name: "name", DataType: "STRING"},
		},
	})
	if err != nil {
		t.Fatalf("failed to create table '%s': %v", name, err)
	}

	t.Cleanup(func() {
		_ = client.Tables.DeleteByName(
			ctx,
			*table.FullyQualifiedName,
			&ometa.DeleteTable1Params{
				HardDelete: ometa.Bool(true),
			},
		)
	})

	return table
}

func TestCreateTable(t *testing.T) {
	ctx := context.Background()
	table := createTestTable(t, ctx, "test_create_table")

	if table.Name != "test_create_table" {
		t.Errorf("expected table name 'test_create_table', got '%s'", table.Name)
	}
}

func TestGetTableByID(t *testing.T) {
	ctx := context.Background()
	table := createTestTable(t, ctx, "test_get_by_id")

	got, err := client.Tables.GetByID(ctx, table.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get table by ID: %v", err)
	}

	if got.Id != table.Id {
		t.Errorf("expected ID '%s', got '%s'", table.Id, got.Id)
	}
}

func TestGetTableByName(t *testing.T) {
	ctx := context.Background()
	table := createTestTable(t, ctx, "test_get_by_name")

	got, err := client.Tables.GetByName(ctx, *table.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get table by name: %v", err)
	}

	if got.Id != table.Id {
		t.Errorf("expected ID '%s', got '%s'", table.Id, got.Id)
	}
}

func TestCreateOrUpdateTable(t *testing.T) {
	ctx := context.Background()
	table := createTestTable(t, ctx, "test_create_or_update")

	desc := "updated description"
	updated, err := client.Tables.CreateOrUpdate(ctx, &ometa.CreateTable{
		Name:           "test_create_or_update",
		DatabaseSchema: testSchema,
		Columns: []ometa.Column{
			{Name: "id", DataType: "INT"},
			{Name: "name", DataType: "STRING"},
		},
		Description: &desc,
	})
	if err != nil {
		t.Fatalf("failed to create or update table: %v", err)
	}

	if updated.Id != table.Id {
		t.Errorf("expected same ID after upsert, got different: '%s' vs '%s'", table.Id, updated.Id)
	}
	if updated.Description == nil || *updated.Description != desc {
		t.Errorf("expected description '%s', got '%v'", desc, updated.Description)
	}
}

func TestPatchTable(t *testing.T) {
	ctx := context.Background()
	table := createTestTable(t, ctx, "test_patch")

	patched, err := client.Tables.Patch(ctx, table.Id.String(), []ometa.JSONPatchOp{
		{Op: "add", Path: "/description", Value: "patched description"},
	})
	if err != nil {
		t.Fatalf("failed to patch table: %v", err)
	}

	if patched.Description == nil || *patched.Description != "patched description" {
		t.Errorf("expected description 'patched description', got '%v'", patched.Description)
	}
}

func TestDeleteTable(t *testing.T) {
	ctx := context.Background()

	table, err := client.Tables.Create(ctx, &ometa.CreateTable{
		Name:           "test_delete",
		DatabaseSchema: testSchema,
		Columns: []ometa.Column{
			{Name: "id", DataType: "INT"},
		},
	})
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	err = client.Tables.Delete(ctx, table.Id.String(), &ometa.DeleteTableParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete table: %v", err)
	}

	_, err = client.Tables.GetByID(ctx, table.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}

func TestDeleteTableByName(t *testing.T) {
	ctx := context.Background()

	table, err := client.Tables.Create(ctx, &ometa.CreateTable{
		Name:           "test_delete_by_name",
		DatabaseSchema: testSchema,
		Columns: []ometa.Column{
			{Name: "id", DataType: "INT"},
		},
	})
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	err = client.Tables.DeleteByName(ctx, *table.FullyQualifiedName, &ometa.DeleteTable1Params{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete table by name: %v", err)
	}

	_, err = client.Tables.GetByName(ctx, *table.FullyQualifiedName, nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}

func TestListTables(t *testing.T) {
	ctx := context.Background()
	table := createTestTable(t, ctx, "test_list")

	found := false
	for tbl, err := range client.Tables.List(ctx, &ometa.ListTablesParams{
		Limit: ometa.Int32(100),
	}) {
		if err != nil {
			t.Fatalf("error during list: %v", err)
		}
		if tbl.Id == table.Id {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("created table not found in list results")
	}
}

func TestListTableVersions(t *testing.T) {
	ctx := context.Background()
	table := createTestTable(t, ctx, "test_versions")

	_, _ = client.Tables.Patch(ctx, table.Id.String(), []ometa.JSONPatchOp{
		{Op: "add", Path: "/description", Value: "v2"},
	})

	history, err := client.Tables.ListVersions(ctx, table.Id.String())
	if err != nil {
		t.Fatalf("failed to list versions: %v", err)
	}

	if len(history.Versions) < 2 {
		t.Errorf("expected at least 2 versions, got %d", len(history.Versions))
	}
}

func TestGetTableVersion(t *testing.T) {
	ctx := context.Background()
	table := createTestTable(t, ctx, "test_get_version")

	got, err := client.Tables.GetVersion(ctx, table.Id.String(), "0.1")
	if err != nil {
		t.Fatalf("failed to get version 0.1: %v", err)
	}

	if got.Id != table.Id {
		t.Errorf("expected ID '%s', got '%s'", table.Id, got.Id)
	}
}

func TestRestoreTable(t *testing.T) {
	ctx := context.Background()

	table, err := client.Tables.Create(ctx, &ometa.CreateTable{
		Name:           "test_restore",
		DatabaseSchema: testSchema,
		Columns: []ometa.Column{
			{Name: "id", DataType: "INT"},
		},
	})
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	t.Cleanup(func() {
		_ = client.Tables.DeleteByName(
			ctx,
			fmt.Sprintf("%s.test_restore", testSchema),
			&ometa.DeleteTable1Params{HardDelete: ometa.Bool(true)},
		)
	})

	err = client.Tables.Delete(ctx, table.Id.String(), &ometa.DeleteTableParams{
		HardDelete: ometa.Bool(false),
	})
	if err != nil {
		t.Fatalf("failed to soft delete table: %v", err)
	}

	restored, err := client.Tables.Restore(ctx, &ometa.RestoreEntity{
		Id: table.Id,
	})
	if err != nil {
		t.Fatalf("failed to restore table: %v", err)
	}

	if restored.Id != table.Id {
		t.Errorf("expected restored ID '%s', got '%s'", table.Id, restored.Id)
	}
}