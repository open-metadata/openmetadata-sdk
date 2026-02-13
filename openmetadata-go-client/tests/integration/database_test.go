package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestDatabase(t *testing.T, ctx context.Context, name string) *ometa.Database {
	t.Helper()

	db, err := client.Databases.Create(ctx, &ometa.CreateDatabase{
		Name:    name,
		Service: testDatabaseService,
	})
	if err != nil {
		t.Fatalf("failed to create database '%s': %v", name, err)
	}

	t.Cleanup(func() {
		_ = client.Databases.DeleteByName(ctx, *db.FullyQualifiedName, &ometa.DeleteDatabaseByFQNParams{
			HardDelete: ometa.Bool(true),
			Recursive:  ometa.Bool(true),
		})
	})

	return db
}

func TestCreateDatabase(t *testing.T) {
	ctx := context.Background()
	db := createTestDatabase(t, ctx, "test_create_db")

	if db.Name != "test_create_db" {
		t.Errorf("expected name 'test_create_db', got '%s'", db.Name)
	}
}

func TestGetDatabaseByID(t *testing.T) {
	ctx := context.Background()
	db := createTestDatabase(t, ctx, "test_get_db_by_id")

	got, err := client.Databases.GetByID(ctx, db.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get database by ID: %v", err)
	}

	if got.Id != db.Id {
		t.Errorf("expected ID '%s', got '%s'", db.Id, got.Id)
	}
}

func TestGetDatabaseByName(t *testing.T) {
	ctx := context.Background()
	db := createTestDatabase(t, ctx, "test_get_db_by_name")

	got, err := client.Databases.GetByName(ctx, *db.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get database by name: %v", err)
	}

	if got.Id != db.Id {
		t.Errorf("expected ID '%s', got '%s'", db.Id, got.Id)
	}
}

func TestListDatabases(t *testing.T) {
	ctx := context.Background()
	db := createTestDatabase(t, ctx, "test_list_db")

	found := false
	for d, err := range client.Databases.List(ctx, &ometa.ListDatabasesParams{
		Limit: ometa.Int32(100),
	}) {
		if err != nil {
			t.Fatalf("error during list: %v", err)
		}
		if d.Id == db.Id {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("created database not found in list results")
	}
}

func TestDeleteDatabase(t *testing.T) {
	ctx := context.Background()

	db, err := client.Databases.Create(ctx, &ometa.CreateDatabase{
		Name:    "test_delete_db",
		Service: testDatabaseService,
	})
	if err != nil {
		t.Fatalf("failed to create database: %v", err)
	}

	err = client.Databases.Delete(ctx, db.Id.String(), &ometa.DeleteDatabaseParams{
		HardDelete: ometa.Bool(true),
		Recursive:  ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete database: %v", err)
	}

	_, err = client.Databases.GetByID(ctx, db.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
