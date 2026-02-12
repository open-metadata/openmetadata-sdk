package ometa

import (
	"context"
	"fmt"
	"iter"
)

const tableBasePath = "/tables"

type TableService struct {
	backend Backend
}

func (s *TableService) List(
	ctx context.Context,
	params *ListTablesParams,
) iter.Seq2[Table, error] {
	return newEntityIterator[Table](ctx, s.backend, tableBasePath, params)
}

func (s *TableService) GetByID(ctx context.Context, id string, params *GetTableByIDParams) (*Table, error) {
    return get[Table](ctx, s.backend, fmt.Sprintf("%s/%s", tableBasePath, id), params)
}

func (s *TableService) GetByName(ctx context.Context, fqn string, params *GetTableByFQNParams) (*Table, error) {
    return get[Table](ctx, s.backend, fmt.Sprintf("%s/name/%s", tableBasePath, fqn), params)
}

func (s *TableService) Create(ctx context.Context, body *CreateTable) (*Table, error) {
    return create[Table](ctx, s.backend, tableBasePath, body)
}

func (s *TableService) CreateOrUpdate(ctx context.Context, body *CreateTable) (*Table, error) {
    return put[Table](ctx, s.backend, tableBasePath, body)
}

func (s *TableService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Table, error) {
    return patch[Table](ctx, s.backend, fmt.Sprintf("%s/%s", tableBasePath, id), ops)
}

func (s *TableService) Delete(ctx context.Context, id string, params *DeleteTableParams) error {
    return delete(ctx, s.backend, fmt.Sprintf("%s/%s", tableBasePath, id))
}

func (s *TableService) DeleteByName(ctx context.Context, fqn string, params *DeleteTable1Params) error {
    return delete(ctx, s.backend, fmt.Sprintf("%s/name/%s", tableBasePath, fqn))
}

func (s *TableService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
    return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", tableBasePath, id), nil)
}

func (s *TableService) GetVersion(ctx context.Context, id string, version string) (*Table, error) {
    return get[Table](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", tableBasePath, id, version), nil)
}

func (s *TableService) Restore(ctx context.Context, body *RestoreEntity) (*Table, error) {
    return put[Table](ctx, s.backend, fmt.Sprintf("%s/restore", tableBasePath), body)
}