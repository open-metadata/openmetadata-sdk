package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestTag(t *testing.T, ctx context.Context, name string) (*ometa.Tag, *ometa.Classification) {
	t.Helper()

	cls := createTestClassification(t, ctx, name+"_parent_cls")

	tag, err := client.Tags.Create(ctx, &ometa.CreateTag{
		Name:           name,
		Description:    "test tag",
		Classification: ometa.Str(*cls.FullyQualifiedName),
	})
	if err != nil {
		t.Fatalf("failed to create tag '%s': %v", name, err)
	}

	t.Cleanup(func() {
		_ = client.Tags.DeleteByName(ctx, *tag.FullyQualifiedName, &ometa.DeleteTagByNameParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return tag, cls
}

func TestCreateTag(t *testing.T) {
	ctx := context.Background()
	tag, _ := createTestTag(t, ctx, "test_create_tag")

	if tag.Name != "test_create_tag" {
		t.Errorf("expected name 'test_create_tag', got '%s'", tag.Name)
	}
}

func TestGetTagByID(t *testing.T) {
	ctx := context.Background()
	tag, _ := createTestTag(t, ctx, "test_get_tag_by_id")

	got, err := client.Tags.GetByID(ctx, tag.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get tag by ID: %v", err)
	}

	if got.Id != tag.Id {
		t.Errorf("expected ID '%s', got '%s'", tag.Id, got.Id)
	}
}

func TestGetTagByName(t *testing.T) {
	ctx := context.Background()
	tag, _ := createTestTag(t, ctx, "test_get_tag_by_name")

	got, err := client.Tags.GetByName(ctx, *tag.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get tag by name: %v", err)
	}

	if got.Id != tag.Id {
		t.Errorf("expected ID '%s', got '%s'", tag.Id, got.Id)
	}
}

func TestDeleteTag(t *testing.T) {
	ctx := context.Background()

	cls := createTestClassification(t, ctx, "test_delete_tag_cls")
	tag, err := client.Tags.Create(ctx, &ometa.CreateTag{
		Name:           "test_delete_tag",
		Description:    "to be deleted",
		Classification: ometa.Str(*cls.FullyQualifiedName),
	})
	if err != nil {
		t.Fatalf("failed to create tag: %v", err)
	}

	err = client.Tags.Delete(ctx, tag.Id.String(), &ometa.DeleteTagParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete tag: %v", err)
	}

	_, err = client.Tags.GetByID(ctx, tag.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
