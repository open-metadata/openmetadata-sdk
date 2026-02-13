package ometa

import (
	"context"
	"fmt"
	"iter"
)

const databaseSchemaBasePath = "databaseSchemas"

type DatabaseSchemaService struct {
	backend Backend
}

func (s *DatabaseSchemaService) List(
	ctx context.Context,
	params *ListDBSchemasParams,
) iter.Seq2[DatabaseSchema, error] {
	return newEntityIterator[DatabaseSchema](ctx, s.backend, databaseSchemaBasePath, params)
}

func (s *DatabaseSchemaService) GetByID(ctx context.Context, id string, params *GetDBSchemaByIDParams) (*DatabaseSchema, error) {
	return get[DatabaseSchema](ctx, s.backend, fmt.Sprintf("%s/%s", databaseSchemaBasePath, id), params)
}

func (s *DatabaseSchemaService) GetByName(ctx context.Context, fqn string, params *GetDBSchemaByFQNParams) (*DatabaseSchema, error) {
	return get[DatabaseSchema](ctx, s.backend, fmt.Sprintf("%s/name/%s", databaseSchemaBasePath, fqn), params)
}

func (s *DatabaseSchemaService) Create(ctx context.Context, body *CreateDatabaseSchema) (*DatabaseSchema, error) {
	return create[DatabaseSchema](ctx, s.backend, databaseSchemaBasePath, body)
}

func (s *DatabaseSchemaService) CreateOrUpdate(ctx context.Context, body *CreateDatabaseSchema) (*DatabaseSchema, error) {
	return put[DatabaseSchema](ctx, s.backend, databaseSchemaBasePath, body)
}

func (s *DatabaseSchemaService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*DatabaseSchema, error) {
	return patch[DatabaseSchema](ctx, s.backend, fmt.Sprintf("%s/%s", databaseSchemaBasePath, id), ops)
}

func (s *DatabaseSchemaService) Delete(ctx context.Context, id string, params *DeleteDBSchemaParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", databaseSchemaBasePath, id), params)
}

func (s *DatabaseSchemaService) DeleteByName(ctx context.Context, fqn string, params *DeleteDBSchemaByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", databaseSchemaBasePath, fqn), params)
}

func (s *DatabaseSchemaService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", databaseSchemaBasePath, id), nil)
}

func (s *DatabaseSchemaService) GetVersion(ctx context.Context, id string, version string) (*DatabaseSchema, error) {
	return get[DatabaseSchema](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", databaseSchemaBasePath, id, version), nil)
}

func (s *DatabaseSchemaService) Restore(ctx context.Context, body *RestoreEntity) (*DatabaseSchema, error) {
	return put[DatabaseSchema](ctx, s.backend, fmt.Sprintf("%s/restore", databaseSchemaBasePath), body)
}
