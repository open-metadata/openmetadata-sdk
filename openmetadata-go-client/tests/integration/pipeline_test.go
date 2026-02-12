package integration

import (
	"context"
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
)

func createTestPipeline(t *testing.T, ctx context.Context, name string) *ometa.Pipeline {
	t.Helper()

	pipeline, err := client.Pipelines.Create(ctx, &ometa.CreatePipeline{
		Name:    name,
		Service: testPipelineService,
	})
	if err != nil {
		t.Fatalf("failed to create pipeline '%s': %v", name, err)
	}

	t.Cleanup(func() {
		client.Pipelines.DeleteByName(ctx, *pipeline.FullyQualifiedName, &ometa.DeletePipelineByFQNParams{
			HardDelete: ometa.Bool(true),
		})
	})

	return pipeline
}

func TestCreatePipeline(t *testing.T) {
	ctx := context.Background()
	pipeline := createTestPipeline(t, ctx, "test_create_pipeline")

	if pipeline.Name != "test_create_pipeline" {
		t.Errorf("expected name 'test_create_pipeline', got '%s'", pipeline.Name)
	}
}

func TestGetPipelineByID(t *testing.T) {
	ctx := context.Background()
	pipeline := createTestPipeline(t, ctx, "test_get_pipeline_by_id")

	got, err := client.Pipelines.GetByID(ctx, pipeline.Id.String(), nil)
	if err != nil {
		t.Fatalf("failed to get pipeline by ID: %v", err)
	}

	if got.Id != pipeline.Id {
		t.Errorf("expected ID '%s', got '%s'", pipeline.Id, got.Id)
	}
}

func TestGetPipelineByName(t *testing.T) {
	ctx := context.Background()
	pipeline := createTestPipeline(t, ctx, "test_get_pipeline_by_name")

	got, err := client.Pipelines.GetByName(ctx, *pipeline.FullyQualifiedName, nil)
	if err != nil {
		t.Fatalf("failed to get pipeline by name: %v", err)
	}

	if got.Id != pipeline.Id {
		t.Errorf("expected ID '%s', got '%s'", pipeline.Id, got.Id)
	}
}

func TestDeletePipeline(t *testing.T) {
	ctx := context.Background()

	pipeline, err := client.Pipelines.Create(ctx, &ometa.CreatePipeline{
		Name:    "test_delete_pipeline",
		Service: testPipelineService,
	})
	if err != nil {
		t.Fatalf("failed to create pipeline: %v", err)
	}

	err = client.Pipelines.Delete(ctx, pipeline.Id.String(), &ometa.DeletePipelineParams{
		HardDelete: ometa.Bool(true),
	})
	if err != nil {
		t.Fatalf("failed to delete pipeline: %v", err)
	}

	_, err = client.Pipelines.GetByID(ctx, pipeline.Id.String(), nil)
	if !ometa.IsNotFound(err) {
		t.Errorf("expected 404 after delete, got: %v", err)
	}
}
