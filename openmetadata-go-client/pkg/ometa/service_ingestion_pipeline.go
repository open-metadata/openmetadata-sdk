package ometa

import (
	"context"
	"fmt"
	"iter"
)

const ingestionPipelineBasePath = "services/ingestionPipelines"

type IngestionPipelineService struct {
	backend Backend
}

func (s *IngestionPipelineService) List(
	ctx context.Context,
	params *ListIngestionPipelinesParams,
) iter.Seq2[IngestionPipeline, error] {
	return newEntityIterator[IngestionPipeline](ctx, s.backend, ingestionPipelineBasePath, params)
}

func (s *IngestionPipelineService) GetByID(ctx context.Context, id string, params *GetIngestionPipelineByIDParams) (*IngestionPipeline, error) {
	return get[IngestionPipeline](ctx, s.backend, fmt.Sprintf("%s/%s", ingestionPipelineBasePath, id), params)
}

func (s *IngestionPipelineService) GetByName(ctx context.Context, fqn string, params *GetSpecificIngestionPipelineByFQNParams) (*IngestionPipeline, error) {
	return get[IngestionPipeline](ctx, s.backend, fmt.Sprintf("%s/name/%s", ingestionPipelineBasePath, fqn), params)
}

func (s *IngestionPipelineService) Create(ctx context.Context, body *CreateIngestionPipeline) (*IngestionPipeline, error) {
	return create[IngestionPipeline](ctx, s.backend, ingestionPipelineBasePath, body)
}

func (s *IngestionPipelineService) CreateOrUpdate(ctx context.Context, body *CreateIngestionPipeline) (*IngestionPipeline, error) {
	return put[IngestionPipeline](ctx, s.backend, ingestionPipelineBasePath, body)
}

func (s *IngestionPipelineService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*IngestionPipeline, error) {
	return patch[IngestionPipeline](ctx, s.backend, fmt.Sprintf("%s/%s", ingestionPipelineBasePath, id), ops)
}

func (s *IngestionPipelineService) Delete(ctx context.Context, id string, params *DeleteIngestionPipelineParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", ingestionPipelineBasePath, id), params)
}

func (s *IngestionPipelineService) DeleteByName(ctx context.Context, fqn string, params *DeleteIngestionPipelineByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", ingestionPipelineBasePath, fqn), params)
}

func (s *IngestionPipelineService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", ingestionPipelineBasePath, id), nil)
}

func (s *IngestionPipelineService) GetVersion(ctx context.Context, id string, version string) (*IngestionPipeline, error) {
	return get[IngestionPipeline](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", ingestionPipelineBasePath, id, version), nil)
}

func (s *IngestionPipelineService) Restore(ctx context.Context, body *RestoreEntity) (*IngestionPipeline, error) {
	return put[IngestionPipeline](ctx, s.backend, fmt.Sprintf("%s/restore", ingestionPipelineBasePath), body)
}
