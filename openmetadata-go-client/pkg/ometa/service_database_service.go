package ometa

import (
	"context"
	"fmt"
	"iter"
)

const databaseServiceBasePath = "services/databaseServices"

type DatabaseServiceService struct {
	backend Backend
}

func (s *DatabaseServiceService) List(
	ctx context.Context,
	params *ListDatabaseServicesParams,
) iter.Seq2[DatabaseService, error] {
	return newEntityIterator[DatabaseService](ctx, s.backend, databaseServiceBasePath, params)
}

func (s *DatabaseServiceService) GetByID(ctx context.Context, id string, params *GetDatabaseServiceByIDParams) (*DatabaseService, error) {
	return get[DatabaseService](ctx, s.backend, fmt.Sprintf("%s/%s", databaseServiceBasePath, id), params)
}

func (s *DatabaseServiceService) GetByName(ctx context.Context, fqn string, params *GetDatabaseServiceByFQNParams) (*DatabaseService, error) {
	return get[DatabaseService](ctx, s.backend, fmt.Sprintf("%s/name/%s", databaseServiceBasePath, fqn), params)
}

func (s *DatabaseServiceService) Create(ctx context.Context, body *CreateDatabaseService) (*DatabaseService, error) {
	return create[DatabaseService](ctx, s.backend, databaseServiceBasePath, body)
}

func (s *DatabaseServiceService) CreateOrUpdate(ctx context.Context, body *CreateDatabaseService) (*DatabaseService, error) {
	return put[DatabaseService](ctx, s.backend, databaseServiceBasePath, body)
}

func (s *DatabaseServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*DatabaseService, error) {
	return patch[DatabaseService](ctx, s.backend, fmt.Sprintf("%s/%s", databaseServiceBasePath, id), ops)
}

func (s *DatabaseServiceService) Delete(ctx context.Context, id string, params *DeleteDatabaseServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", databaseServiceBasePath, id), params)
}

func (s *DatabaseServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteDatabaseServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", databaseServiceBasePath, fqn), params)
}

func (s *DatabaseServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", databaseServiceBasePath, id), nil)
}

func (s *DatabaseServiceService) GetVersion(ctx context.Context, id string, version string) (*DatabaseService, error) {
	return get[DatabaseService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", databaseServiceBasePath, id, version), nil)
}

func (s *DatabaseServiceService) Restore(ctx context.Context, body *RestoreEntity) (*DatabaseService, error) {
	return put[DatabaseService](ctx, s.backend, fmt.Sprintf("%s/restore", databaseServiceBasePath), body)
}
