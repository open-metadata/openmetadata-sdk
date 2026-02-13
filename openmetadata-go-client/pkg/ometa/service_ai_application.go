package ometa

import (
	"context"
	"fmt"
	"iter"
)

const aiApplicationBasePath = "aiApplications"

type AIApplicationService struct {
	backend Backend
}

func (s *AIApplicationService) List(
	ctx context.Context,
	params *ListAIApplicationsParams,
) iter.Seq2[AIApplication, error] {
	return newEntityIterator[AIApplication](ctx, s.backend, aiApplicationBasePath, params)
}

func (s *AIApplicationService) GetByID(ctx context.Context, id string, params *GetAIApplicationByIDParams) (*AIApplication, error) {
	return get[AIApplication](ctx, s.backend, fmt.Sprintf("%s/%s", aiApplicationBasePath, id), params)
}

func (s *AIApplicationService) GetByName(ctx context.Context, fqn string, params *GetAIApplicationByFQNParams) (*AIApplication, error) {
	return get[AIApplication](ctx, s.backend, fmt.Sprintf("%s/name/%s", aiApplicationBasePath, fqn), params)
}

func (s *AIApplicationService) Create(ctx context.Context, body *CreateAIApplication) (*AIApplication, error) {
	return create[AIApplication](ctx, s.backend, aiApplicationBasePath, body)
}

func (s *AIApplicationService) CreateOrUpdate(ctx context.Context, body *CreateAIApplication) (*AIApplication, error) {
	return put[AIApplication](ctx, s.backend, aiApplicationBasePath, body)
}

func (s *AIApplicationService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*AIApplication, error) {
	return patch[AIApplication](ctx, s.backend, fmt.Sprintf("%s/%s", aiApplicationBasePath, id), ops)
}

func (s *AIApplicationService) Delete(ctx context.Context, id string, params *DeleteAIApplicationParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", aiApplicationBasePath, id), params)
}

func (s *AIApplicationService) DeleteByName(ctx context.Context, fqn string, params *DeleteAIApplicationByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", aiApplicationBasePath, fqn), params)
}

func (s *AIApplicationService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", aiApplicationBasePath, id), nil)
}

func (s *AIApplicationService) GetVersion(ctx context.Context, id string, version string) (*AIApplication, error) {
	return get[AIApplication](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", aiApplicationBasePath, id, version), nil)
}

func (s *AIApplicationService) Restore(ctx context.Context, body *RestoreEntity) (*AIApplication, error) {
	return put[AIApplication](ctx, s.backend, fmt.Sprintf("%s/restore", aiApplicationBasePath), body)
}
