package ometa

import (
	"context"
	"fmt"
	"iter"
)

const databaseBasePath = "databases"

type DatabaseSvc struct {
	backend Backend
}

func (s *DatabaseSvc) List(
	ctx context.Context,
	params *ListDatabasesParams,
) iter.Seq2[Database, error] {
	return newEntityIterator[Database](ctx, s.backend, databaseBasePath, params)
}

func (s *DatabaseSvc) GetByID(ctx context.Context, id string, params *GetDatabaseByIDParams) (*Database, error) {
	return get[Database](ctx, s.backend, fmt.Sprintf("%s/%s", databaseBasePath, id), params)
}

func (s *DatabaseSvc) GetByName(ctx context.Context, fqn string, params *GetDatabaseByFQNParams) (*Database, error) {
	return get[Database](ctx, s.backend, fmt.Sprintf("%s/name/%s", databaseBasePath, fqn), params)
}

func (s *DatabaseSvc) Create(ctx context.Context, body *CreateDatabase) (*Database, error) {
	return create[Database](ctx, s.backend, databaseBasePath, body)
}

func (s *DatabaseSvc) CreateOrUpdate(ctx context.Context, body *CreateDatabase) (*Database, error) {
	return put[Database](ctx, s.backend, databaseBasePath, body)
}

func (s *DatabaseSvc) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Database, error) {
	return patch[Database](ctx, s.backend, fmt.Sprintf("%s/%s", databaseBasePath, id), ops)
}

func (s *DatabaseSvc) Delete(ctx context.Context, id string, params *DeleteDatabaseParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", databaseBasePath, id), params)
}

func (s *DatabaseSvc) DeleteByName(ctx context.Context, fqn string, params *DeleteDatabaseByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", databaseBasePath, fqn), params)
}

func (s *DatabaseSvc) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", databaseBasePath, id), nil)
}

func (s *DatabaseSvc) GetVersion(ctx context.Context, id string, version string) (*Database, error) {
	return get[Database](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", databaseBasePath, id, version), nil)
}

func (s *DatabaseSvc) Restore(ctx context.Context, body *RestoreEntity) (*Database, error) {
	return put[Database](ctx, s.backend, fmt.Sprintf("%s/restore", databaseBasePath), body)
}
