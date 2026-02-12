package ometa

import (
	"context"
	"fmt"
	"iter"
)

const storageServiceBasePath = "services/storageServices"

type StorageServiceService struct {
	backend Backend
}

func (s *StorageServiceService) List(
	ctx context.Context,
	params *ListStorageServicesParams,
) iter.Seq2[StorageService, error] {
	return newEntityIterator[StorageService](ctx, s.backend, storageServiceBasePath, params)
}

func (s *StorageServiceService) GetByID(ctx context.Context, id string, params *GetStorageServiceByIDParams) (*StorageService, error) {
	return get[StorageService](ctx, s.backend, fmt.Sprintf("%s/%s", storageServiceBasePath, id), params)
}

func (s *StorageServiceService) GetByName(ctx context.Context, fqn string, params *GetStorageServiceByFQNParams) (*StorageService, error) {
	return get[StorageService](ctx, s.backend, fmt.Sprintf("%s/name/%s", storageServiceBasePath, fqn), params)
}

func (s *StorageServiceService) Create(ctx context.Context, body *CreateStorageService) (*StorageService, error) {
	return create[StorageService](ctx, s.backend, storageServiceBasePath, body)
}

func (s *StorageServiceService) CreateOrUpdate(ctx context.Context, body *CreateStorageService) (*StorageService, error) {
	return put[StorageService](ctx, s.backend, storageServiceBasePath, body)
}

func (s *StorageServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*StorageService, error) {
	return patch[StorageService](ctx, s.backend, fmt.Sprintf("%s/%s", storageServiceBasePath, id), ops)
}

func (s *StorageServiceService) Delete(ctx context.Context, id string, params *DeleteStorageServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", storageServiceBasePath, id), params)
}

func (s *StorageServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteStorageServiceByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", storageServiceBasePath, fqn), params)
}

func (s *StorageServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", storageServiceBasePath, id), nil)
}

func (s *StorageServiceService) GetVersion(ctx context.Context, id string, version string) (*StorageService, error) {
	return get[StorageService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", storageServiceBasePath, id, version), nil)
}

func (s *StorageServiceService) Restore(ctx context.Context, body *RestoreEntity) (*StorageService, error) {
	return put[StorageService](ctx, s.backend, fmt.Sprintf("%s/restore", storageServiceBasePath), body)
}
