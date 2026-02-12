package ometa

import (
	"context"
	"fmt"
	"iter"
)

const apiCollectionBasePath = "apiCollections"

type APICollectionService struct {
	backend Backend
}

func (s *APICollectionService) List(
	ctx context.Context,
	params *ListAPICollectionsParams,
) iter.Seq2[APICollection, error] {
	return newEntityIterator[APICollection](ctx, s.backend, apiCollectionBasePath, params)
}

func (s *APICollectionService) GetByID(ctx context.Context, id string, params *GetAPICollectionByIDParams) (*APICollection, error) {
	return get[APICollection](ctx, s.backend, fmt.Sprintf("%s/%s", apiCollectionBasePath, id), params)
}

func (s *APICollectionService) GetByName(ctx context.Context, fqn string, params *GetAPICollectionByFQNParams) (*APICollection, error) {
	return get[APICollection](ctx, s.backend, fmt.Sprintf("%s/name/%s", apiCollectionBasePath, fqn), params)
}

func (s *APICollectionService) Create(ctx context.Context, body *CreateAPICollection) (*APICollection, error) {
	return create[APICollection](ctx, s.backend, apiCollectionBasePath, body)
}

func (s *APICollectionService) CreateOrUpdate(ctx context.Context, body *CreateAPICollection) (*APICollection, error) {
	return put[APICollection](ctx, s.backend, apiCollectionBasePath, body)
}

func (s *APICollectionService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*APICollection, error) {
	return patch[APICollection](ctx, s.backend, fmt.Sprintf("%s/%s", apiCollectionBasePath, id), ops)
}

func (s *APICollectionService) Delete(ctx context.Context, id string, params *DeleteAPICollectionParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", apiCollectionBasePath, id), params)
}

func (s *APICollectionService) DeleteByName(ctx context.Context, fqn string, params *DeleteAPICollectionByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", apiCollectionBasePath, fqn), params)
}

func (s *APICollectionService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", apiCollectionBasePath, id), nil)
}

func (s *APICollectionService) GetVersion(ctx context.Context, id string, version string) (*APICollection, error) {
	return get[APICollection](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", apiCollectionBasePath, id, version), nil)
}

func (s *APICollectionService) Restore(ctx context.Context, body *RestoreEntity) (*APICollection, error) {
	return put[APICollection](ctx, s.backend, fmt.Sprintf("%s/restore", apiCollectionBasePath), body)
}
