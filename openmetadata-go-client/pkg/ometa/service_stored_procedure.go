package ometa

import (
	"context"
	"fmt"
	"iter"
)

const storedProcedureBasePath = "storedProcedures"

type StoredProcedureService struct {
	backend Backend
}

func (s *StoredProcedureService) List(
	ctx context.Context,
	params *ListStoredProceduresParams,
) iter.Seq2[StoredProcedure, error] {
	return newEntityIterator[StoredProcedure](ctx, s.backend, storedProcedureBasePath, params)
}

func (s *StoredProcedureService) GetByID(ctx context.Context, id string, params *GetStoredProcedureByIDParams) (*StoredProcedure, error) {
	return get[StoredProcedure](ctx, s.backend, fmt.Sprintf("%s/%s", storedProcedureBasePath, id), params)
}

func (s *StoredProcedureService) GetByName(ctx context.Context, fqn string, params *GetStoredProcedureByFQNParams) (*StoredProcedure, error) {
	return get[StoredProcedure](ctx, s.backend, fmt.Sprintf("%s/name/%s", storedProcedureBasePath, fqn), params)
}

func (s *StoredProcedureService) Create(ctx context.Context, body *CreateStoredProcedure) (*StoredProcedure, error) {
	return create[StoredProcedure](ctx, s.backend, storedProcedureBasePath, body)
}

func (s *StoredProcedureService) CreateOrUpdate(ctx context.Context, body *CreateStoredProcedure) (*StoredProcedure, error) {
	return put[StoredProcedure](ctx, s.backend, storedProcedureBasePath, body)
}

func (s *StoredProcedureService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*StoredProcedure, error) {
	return patch[StoredProcedure](ctx, s.backend, fmt.Sprintf("%s/%s", storedProcedureBasePath, id), ops)
}

func (s *StoredProcedureService) Delete(ctx context.Context, id string, params *DeleteStoredProcedureParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", storedProcedureBasePath, id), params)
}

func (s *StoredProcedureService) DeleteByName(ctx context.Context, fqn string, params *DeleteDBSchemaByFQN1Params) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", storedProcedureBasePath, fqn), params)
}

func (s *StoredProcedureService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", storedProcedureBasePath, id), nil)
}

func (s *StoredProcedureService) GetVersion(ctx context.Context, id string, version string) (*StoredProcedure, error) {
	return get[StoredProcedure](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", storedProcedureBasePath, id, version), nil)
}

func (s *StoredProcedureService) Restore(ctx context.Context, body *RestoreEntity) (*StoredProcedure, error) {
	return put[StoredProcedure](ctx, s.backend, fmt.Sprintf("%s/restore", storedProcedureBasePath), body)
}
