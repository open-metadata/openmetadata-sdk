package ometa

import (
	"context"
	"fmt"
	"iter"
)

const pipelineServiceBasePath = "services/pipelineServices"

type PipelineServiceService struct {
	backend Backend
}

func (s *PipelineServiceService) List(
	ctx context.Context,
	params *ListPipelineServiceParams,
) iter.Seq2[PipelineService, error] {
	return newEntityIterator[PipelineService](ctx, s.backend, pipelineServiceBasePath, params)
}

func (s *PipelineServiceService) GetByID(ctx context.Context, id string, params *GetPipelineServiceByIDParams) (*PipelineService, error) {
	return get[PipelineService](ctx, s.backend, fmt.Sprintf("%s/%s", pipelineServiceBasePath, id), params)
}

func (s *PipelineServiceService) GetByName(ctx context.Context, fqn string, params *GetPipelineServiceByFQNParams) (*PipelineService, error) {
	return get[PipelineService](ctx, s.backend, fmt.Sprintf("%s/name/%s", pipelineServiceBasePath, fqn), params)
}

func (s *PipelineServiceService) Create(ctx context.Context, body *CreatePipelineService) (*PipelineService, error) {
	return create[PipelineService](ctx, s.backend, pipelineServiceBasePath, body)
}

func (s *PipelineServiceService) CreateOrUpdate(ctx context.Context, body *CreatePipelineService) (*PipelineService, error) {
	return put[PipelineService](ctx, s.backend, pipelineServiceBasePath, body)
}

func (s *PipelineServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*PipelineService, error) {
	return patch[PipelineService](ctx, s.backend, fmt.Sprintf("%s/%s", pipelineServiceBasePath, id), ops)
}

func (s *PipelineServiceService) Delete(ctx context.Context, id string, params *DeletePipelineServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", pipelineServiceBasePath, id), params)
}

func (s *PipelineServiceService) DeleteByName(ctx context.Context, fqn string, params *DeletePipelineServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", pipelineServiceBasePath, fqn), params)
}

func (s *PipelineServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", pipelineServiceBasePath, id), nil)
}

func (s *PipelineServiceService) GetVersion(ctx context.Context, id string, version string) (*PipelineService, error) {
	return get[PipelineService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", pipelineServiceBasePath, id, version), nil)
}

func (s *PipelineServiceService) Restore(ctx context.Context, body *RestoreEntity) (*PipelineService, error) {
	return put[PipelineService](ctx, s.backend, fmt.Sprintf("%s/restore", pipelineServiceBasePath), body)
}
