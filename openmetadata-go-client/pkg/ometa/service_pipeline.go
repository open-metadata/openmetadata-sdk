package ometa

import (
	"context"
	"fmt"
	"iter"
)

const pipelineBasePath = "pipelines"

type PipelineSvc struct {
	backend Backend
}

func (s *PipelineSvc) List(
	ctx context.Context,
	params *ListPipelinesParams,
) iter.Seq2[Pipeline, error] {
	return newEntityIterator[Pipeline](ctx, s.backend, pipelineBasePath, params)
}

func (s *PipelineSvc) GetByID(ctx context.Context, id string, params *GetPipelineWithIDParams) (*Pipeline, error) {
	return get[Pipeline](ctx, s.backend, fmt.Sprintf("%s/%s", pipelineBasePath, id), params)
}

func (s *PipelineSvc) GetByName(ctx context.Context, fqn string, params *GetPipelineByFQNParams) (*Pipeline, error) {
	return get[Pipeline](ctx, s.backend, fmt.Sprintf("%s/name/%s", pipelineBasePath, fqn), params)
}

func (s *PipelineSvc) Create(ctx context.Context, body *CreatePipeline) (*Pipeline, error) {
	return create[Pipeline](ctx, s.backend, pipelineBasePath, body)
}

func (s *PipelineSvc) CreateOrUpdate(ctx context.Context, body *CreatePipeline) (*Pipeline, error) {
	return put[Pipeline](ctx, s.backend, pipelineBasePath, body)
}

func (s *PipelineSvc) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Pipeline, error) {
	return patch[Pipeline](ctx, s.backend, fmt.Sprintf("%s/%s", pipelineBasePath, id), ops)
}

func (s *PipelineSvc) Delete(ctx context.Context, id string, params *DeletePipelineParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", pipelineBasePath, id), params)
}

func (s *PipelineSvc) DeleteByName(ctx context.Context, fqn string, params *DeletePipelineByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", pipelineBasePath, fqn), params)
}

func (s *PipelineSvc) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", pipelineBasePath, id), nil)
}

func (s *PipelineSvc) GetVersion(ctx context.Context, id string, version string) (*Pipeline, error) {
	return get[Pipeline](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", pipelineBasePath, id, version), nil)
}

func (s *PipelineSvc) Restore(ctx context.Context, body *RestoreEntity) (*Pipeline, error) {
	return put[Pipeline](ctx, s.backend, fmt.Sprintf("%s/restore", pipelineBasePath), body)
}
