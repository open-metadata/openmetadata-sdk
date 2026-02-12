package ometa

import (
	"context"
	"fmt"
	"iter"
)

const searchIndexBasePath = "searchIndexes"

type SearchIndexService struct {
	backend Backend
}

func (s *SearchIndexService) List(
	ctx context.Context,
	params *ListSearchIndexesParams,
) iter.Seq2[SearchIndex, error] {
	return newEntityIterator[SearchIndex](ctx, s.backend, searchIndexBasePath, params)
}

func (s *SearchIndexService) GetByID(ctx context.Context, id string, params *Get6Params) (*SearchIndex, error) {
	return get[SearchIndex](ctx, s.backend, fmt.Sprintf("%s/%s", searchIndexBasePath, id), params)
}

func (s *SearchIndexService) GetByName(ctx context.Context, fqn string, params *GetSearchIndexByFQNParams) (*SearchIndex, error) {
	return get[SearchIndex](ctx, s.backend, fmt.Sprintf("%s/name/%s", searchIndexBasePath, fqn), params)
}

func (s *SearchIndexService) Create(ctx context.Context, body *CreateSearchIndex) (*SearchIndex, error) {
	return create[SearchIndex](ctx, s.backend, searchIndexBasePath, body)
}

func (s *SearchIndexService) CreateOrUpdate(ctx context.Context, body *CreateSearchIndex) (*SearchIndex, error) {
	return put[SearchIndex](ctx, s.backend, searchIndexBasePath, body)
}

func (s *SearchIndexService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*SearchIndex, error) {
	return patch[SearchIndex](ctx, s.backend, fmt.Sprintf("%s/%s", searchIndexBasePath, id), ops)
}

func (s *SearchIndexService) Delete(ctx context.Context, id string, params *DeleteSearchIndexParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", searchIndexBasePath, id), params)
}

func (s *SearchIndexService) DeleteByName(ctx context.Context, fqn string, params *DeleteSearchIndexByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", searchIndexBasePath, fqn), params)
}

func (s *SearchIndexService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", searchIndexBasePath, id), nil)
}

func (s *SearchIndexService) GetVersion(ctx context.Context, id string, version string) (*SearchIndex, error) {
	return get[SearchIndex](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", searchIndexBasePath, id, version), nil)
}

func (s *SearchIndexService) Restore(ctx context.Context, body *RestoreEntity) (*SearchIndex, error) {
	return put[SearchIndex](ctx, s.backend, fmt.Sprintf("%s/restore", searchIndexBasePath), body)
}
