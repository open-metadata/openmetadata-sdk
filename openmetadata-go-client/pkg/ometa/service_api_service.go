package ometa

import (
	"context"
	"fmt"
	"iter"
)

const apiServiceBasePath = "services/apiServices"

type ApiServiceService struct {
	backend Backend
}

func (s *ApiServiceService) List(
	ctx context.Context,
	params *ListAPIServicesParams,
) iter.Seq2[ApiService, error] {
	return newEntityIterator[ApiService](ctx, s.backend, apiServiceBasePath, params)
}

func (s *ApiServiceService) GetByID(ctx context.Context, id string, params *GetAPIServiceByIDParams) (*ApiService, error) {
	return get[ApiService](ctx, s.backend, fmt.Sprintf("%s/%s", apiServiceBasePath, id), params)
}

func (s *ApiServiceService) GetByName(ctx context.Context, fqn string, params *GetAPIServiceByFQNParams) (*ApiService, error) {
	return get[ApiService](ctx, s.backend, fmt.Sprintf("%s/name/%s", apiServiceBasePath, fqn), params)
}

func (s *ApiServiceService) Create(ctx context.Context, body *CreateApiService) (*ApiService, error) {
	return create[ApiService](ctx, s.backend, apiServiceBasePath, body)
}

func (s *ApiServiceService) CreateOrUpdate(ctx context.Context, body *CreateApiService) (*ApiService, error) {
	return put[ApiService](ctx, s.backend, apiServiceBasePath, body)
}

func (s *ApiServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*ApiService, error) {
	return patch[ApiService](ctx, s.backend, fmt.Sprintf("%s/%s", apiServiceBasePath, id), ops)
}

func (s *ApiServiceService) Delete(ctx context.Context, id string, params *DeleteAPIServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", apiServiceBasePath, id), params)
}

func (s *ApiServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteAPIServiceByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", apiServiceBasePath, fqn), params)
}

func (s *ApiServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", apiServiceBasePath, id), nil)
}

func (s *ApiServiceService) GetVersion(ctx context.Context, id string, version string) (*ApiService, error) {
	return get[ApiService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", apiServiceBasePath, id, version), nil)
}

func (s *ApiServiceService) Restore(ctx context.Context, body *RestoreEntity) (*ApiService, error) {
	return put[ApiService](ctx, s.backend, fmt.Sprintf("%s/restore", apiServiceBasePath), body)
}
