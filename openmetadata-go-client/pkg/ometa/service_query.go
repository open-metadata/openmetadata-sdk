package ometa

import (
	"context"
	"fmt"
	"iter"
)

const queryBasePath = "queries"

type QueryService struct {
	backend Backend
}

func (s *QueryService) List(
	ctx context.Context,
	params *ListQueriesParams,
) iter.Seq2[Query, error] {
	return newEntityIterator[Query](ctx, s.backend, queryBasePath, params)
}

func (s *QueryService) GetByID(ctx context.Context, id string, params *GetQueryByIdParams) (*Query, error) {
	return get[Query](ctx, s.backend, fmt.Sprintf("%s/%s", queryBasePath, id), params)
}

func (s *QueryService) GetByName(ctx context.Context, fqn string, params *GetQueryFqnParams) (*Query, error) {
	return get[Query](ctx, s.backend, fmt.Sprintf("%s/name/%s", queryBasePath, fqn), params)
}

func (s *QueryService) Create(ctx context.Context, body *CreateQuery) (*Query, error) {
	return create[Query](ctx, s.backend, queryBasePath, body)
}

func (s *QueryService) CreateOrUpdate(ctx context.Context, body *CreateQuery) (*Query, error) {
	return put[Query](ctx, s.backend, queryBasePath, body)
}

func (s *QueryService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Query, error) {
	return patch[Query](ctx, s.backend, fmt.Sprintf("%s/%s", queryBasePath, id), ops)
}

func (s *QueryService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", queryBasePath, id), nil)
}

func (s *QueryService) GetVersion(ctx context.Context, id string, version string) (*Query, error) {
	return get[Query](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", queryBasePath, id, version), nil)
}

func (s *QueryService) Restore(ctx context.Context, body *RestoreEntity) (*Query, error) {
	return put[Query](ctx, s.backend, fmt.Sprintf("%s/restore", queryBasePath), body)
}
