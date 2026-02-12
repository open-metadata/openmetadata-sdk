package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestClassification(t *testing.T, ctx context.Context, name string) *ometa.Classification {
	t.Helper()

	cls, err := client.Classifications.Create(ctx, &ometa.CreateClassification{
		Name:        name,
		Description: "test classification",
	})
	if err != nil {
		t.Fatalf("failed to create classification '%s': %v", name, err)
	}

	t.Cleanup(func() {
		client.Classifications.DeleteByName(ctx, *cls.FullyQualifiedName, &ometa.DeleteClassificationByNameParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return cls
}

func TestCreateClassification(t *testing.T) {
	ctx := context.Background()
	cls := createTestClassification(t, ctx, "test_create_cls")

	if cls.Name != "test_create_cls" {
		t.Errorf("expected name 'test_create_cls', got '%s'", cls.Name)
	}
}

func TestGetClassificationByID(t *testing.T) {
	ctx := context.Background()
	cls := createTestClassification(t, ctx, "test_get_cls_by_id")

	got, err := client.Classifications.GetByID(ctx, cls.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get classification by ID: %v", err)
	}

	if got.Id != cls.Id {
		t.Errorf("expected ID '%s', got '%s'", cls.Id, got.Id)
	}
}

func TestGetClassificationByName(t *testing.T) {
	ctx := context.Background()
	cls := createTestClassification(t, ctx, "test_get_cls_by_name")

	got, err := client.Classifications.GetByName(ctx, *cls.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get classification by name: %v", err)
	}

	if got.Id != cls.Id {
		t.Errorf("expected ID '%s', got '%s'", cls.Id, got.Id)
	}
}

func TestDeleteClassification(t *testing.T) {
	ctx := context.Background()

	cls, err := client.Classifications.Create(ctx, &ometa.CreateClassification{
		Name:        "test_delete_cls",
		Description: "to be deleted",
	})
	if err != nil {
		t.Fatalf("failed to create classification: %v", err)
	}

	err = client.Classifications.Delete(ctx, cls.Id.String(), &ometa.DeleteClassificationParams{
		HardDelete: ometa.Bool(true),
		Recursive:  ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete classification: %v", err)
	}

	_, err = client.Classifications.GetByID(ctx, cls.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
