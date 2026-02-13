package ometa

import (
	"context"
	"fmt"
	"iter"
)

const apiEndpointBasePath = "apiEndpoints"

type APIEndpointService struct {
	backend Backend
}

func (s *APIEndpointService) List(
	ctx context.Context,
	params *ListAPIEndpointsParams,
) iter.Seq2[APIEndpoint, error] {
	return newEntityIterator[APIEndpoint](ctx, s.backend, apiEndpointBasePath, params)
}

func (s *APIEndpointService) GetByID(ctx context.Context, id string, params *GetEndpointByIdParams) (*APIEndpoint, error) {
	return get[APIEndpoint](ctx, s.backend, fmt.Sprintf("%s/%s", apiEndpointBasePath, id), params)
}

func (s *APIEndpointService) GetByName(ctx context.Context, fqn string, params *GetEndpointByFQNParams) (*APIEndpoint, error) {
	return get[APIEndpoint](ctx, s.backend, fmt.Sprintf("%s/name/%s", apiEndpointBasePath, fqn), params)
}

func (s *APIEndpointService) Create(ctx context.Context, body *CreateAPIEndpoint) (*APIEndpoint, error) {
	return create[APIEndpoint](ctx, s.backend, apiEndpointBasePath, body)
}

func (s *APIEndpointService) CreateOrUpdate(ctx context.Context, body *CreateAPIEndpoint) (*APIEndpoint, error) {
	return put[APIEndpoint](ctx, s.backend, apiEndpointBasePath, body)
}

func (s *APIEndpointService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*APIEndpoint, error) {
	return patch[APIEndpoint](ctx, s.backend, fmt.Sprintf("%s/%s", apiEndpointBasePath, id), ops)
}

func (s *APIEndpointService) Delete(ctx context.Context, id string, params *DeleteAPIEndpointParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", apiEndpointBasePath, id), params)
}

func (s *APIEndpointService) DeleteByName(ctx context.Context, fqn string, params *DeleteAPIEndpointByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", apiEndpointBasePath, fqn), params)
}

func (s *APIEndpointService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", apiEndpointBasePath, id), nil)
}

func (s *APIEndpointService) GetVersion(ctx context.Context, id string, version string) (*APIEndpoint, error) {
	return get[APIEndpoint](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", apiEndpointBasePath, id, version), nil)
}

func (s *APIEndpointService) Restore(ctx context.Context, body *RestoreEntity) (*APIEndpoint, error) {
	return put[APIEndpoint](ctx, s.backend, fmt.Sprintf("%s/restore", apiEndpointBasePath), body)
}
