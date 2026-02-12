package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestGlossary(t *testing.T, ctx context.Context, name string) *ometa.Glossary {
	t.Helper()

	glossary, err := client.Glossaries.Create(ctx, &ometa.CreateGlossary{
		Name:        name,
		Description: "test glossary",
	})
	if err != nil {
		t.Fatalf("failed to create glossary '%s': %v", name, err)
	}

	t.Cleanup(func() {
		_ = client.Glossaries.DeleteByName(ctx, *glossary.FullyQualifiedName, &ometa.DeleteGlossaryByNameParams{
			HardDelete: ometa.Bool(true),
			Recursive:  ometa.Bool(true),
		})
	})

	return glossary
}

func createTestGlossaryTerm(t *testing.T, ctx context.Context, name string) (*ometa.GlossaryTerm, *ometa.Glossary) {
	t.Helper()

	glossary := createTestGlossary(t, ctx, name+"_parent_glossary")

	term, err := client.GlossaryTerms.Create(ctx, &ometa.CreateGlossaryTerm{
		Name:        name,
		Description: "test glossary term",
		Glossary:    *glossary.FullyQualifiedName,
	})
	if err != nil {
		t.Fatalf("failed to create glossary term '%s': %v", name, err)
	}

	t.Cleanup(func() {
		_ = client.GlossaryTerms.DeleteByName(ctx, *term.FullyQualifiedName, &ometa.DeleteGlossaryTermByNameParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return term, glossary
}

func TestCreateGlossary(t *testing.T) {
	ctx := context.Background()
	glossary := createTestGlossary(t, ctx, "test_create_glossary")

	if glossary.Name != "test_create_glossary" {
		t.Errorf("expected name 'test_create_glossary', got '%s'", glossary.Name)
	}
}

func TestGetGlossaryByID(t *testing.T) {
	ctx := context.Background()
	glossary := createTestGlossary(t, ctx, "test_get_glossary_by_id")

	got, err := client.Glossaries.GetByID(ctx, glossary.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get glossary by ID: %v", err)
	}

	if got.Id != glossary.Id {
		t.Errorf("expected ID '%s', got '%s'", glossary.Id, got.Id)
	}
}

func TestGetGlossaryByName(t *testing.T) {
	ctx := context.Background()
	glossary := createTestGlossary(t, ctx, "test_get_glossary_by_name")

	got, err := client.Glossaries.GetByName(ctx, *glossary.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get glossary by name: %v", err)
	}

	if got.Id != glossary.Id {
		t.Errorf("expected ID '%s', got '%s'", glossary.Id, got.Id)
	}
}

func TestDeleteGlossary(t *testing.T) {
	ctx := context.Background()

	glossary, err := client.Glossaries.Create(ctx, &ometa.CreateGlossary{
		Name:        "test_delete_glossary",
		Description: "to be deleted",
	})
	if err != nil {
		t.Fatalf("failed to create glossary: %v", err)
	}

	err = client.Glossaries.Delete(ctx, glossary.Id.String(), &ometa.DeleteGlossaryParams{
		HardDelete: ometa.Bool(true),
		Recursive:  ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete glossary: %v", err)
	}

	_, err = client.Glossaries.GetByID(ctx, glossary.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}

func TestCreateGlossaryTerm(t *testing.T) {
	ctx := context.Background()
	term, _ := createTestGlossaryTerm(t, ctx, "test_create_term")

	if term.Name != "test_create_term" {
		t.Errorf("expected name 'test_create_term', got '%s'", term.Name)
	}
}

func TestGetGlossaryTermByID(t *testing.T) {
	ctx := context.Background()
	term, _ := createTestGlossaryTerm(t, ctx, "test_get_term_by_id")

	got, err := client.GlossaryTerms.GetByID(ctx, term.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get glossary term by ID: %v", err)
	}

	if got.Id != term.Id {
		t.Errorf("expected ID '%s', got '%s'", term.Id, got.Id)
	}
}

func TestGetGlossaryTermByName(t *testing.T) {
	ctx := context.Background()
	term, _ := createTestGlossaryTerm(t, ctx, "test_get_term_by_name")

	got, err := client.GlossaryTerms.GetByName(ctx, *term.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get glossary term by name: %v", err)
	}

	if got.Id != term.Id {
		t.Errorf("expected ID '%s', got '%s'", term.Id, got.Id)
	}
}
