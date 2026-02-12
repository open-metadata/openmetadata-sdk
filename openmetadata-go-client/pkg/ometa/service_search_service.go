package ometa

import (
	"context"
	"fmt"
	"iter"
)

const searchServiceBasePath = "services/searchServices"

type SearchServiceService struct {
	backend Backend
}

func (s *SearchServiceService) List(
	ctx context.Context,
	params *ListSearchServicesParams,
) iter.Seq2[SearchService, error] {
	return newEntityIterator[SearchService](ctx, s.backend, searchServiceBasePath, params)
}

func (s *SearchServiceService) GetByID(ctx context.Context, id string, params *GetSearchServiceByIDParams) (*SearchService, error) {
	return get[SearchService](ctx, s.backend, fmt.Sprintf("%s/%s", searchServiceBasePath, id), params)
}

func (s *SearchServiceService) GetByName(ctx context.Context, fqn string, params *GetSearchServiceByFQNParams) (*SearchService, error) {
	return get[SearchService](ctx, s.backend, fmt.Sprintf("%s/name/%s", searchServiceBasePath, fqn), params)
}

func (s *SearchServiceService) Create(ctx context.Context, body *CreateSearchService) (*SearchService, error) {
	return create[SearchService](ctx, s.backend, searchServiceBasePath, body)
}

func (s *SearchServiceService) CreateOrUpdate(ctx context.Context, body *CreateSearchService) (*SearchService, error) {
	return put[SearchService](ctx, s.backend, searchServiceBasePath, body)
}

func (s *SearchServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*SearchService, error) {
	return patch[SearchService](ctx, s.backend, fmt.Sprintf("%s/%s", searchServiceBasePath, id), ops)
}

func (s *SearchServiceService) Delete(ctx context.Context, id string, params *DeleteSearchServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", searchServiceBasePath, id), params)
}

func (s *SearchServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteSearchServiceByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", searchServiceBasePath, fqn), params)
}

func (s *SearchServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", searchServiceBasePath, id), nil)
}

func (s *SearchServiceService) GetVersion(ctx context.Context, id string, version string) (*SearchService, error) {
	return get[SearchService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", searchServiceBasePath, id, version), nil)
}

func (s *SearchServiceService) Restore(ctx context.Context, body *RestoreEntity) (*SearchService, error) {
	return put[SearchService](ctx, s.backend, fmt.Sprintf("%s/restore", searchServiceBasePath), body)
}
