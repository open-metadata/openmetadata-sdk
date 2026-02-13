package ometa

import (
	"context"
	"fmt"
	"iter"
)

const directoryBasePath = "drives/directories"

type DirectoryService struct {
	backend Backend
}

func (s *DirectoryService) List(
	ctx context.Context,
	params *ListDirectoriesParams,
) iter.Seq2[Directory, error] {
	return newEntityIterator[Directory](ctx, s.backend, directoryBasePath, params)
}

func (s *DirectoryService) GetByID(ctx context.Context, id string, params *GetDirectoryByIDParams) (*Directory, error) {
	return get[Directory](ctx, s.backend, fmt.Sprintf("%s/%s", directoryBasePath, id), params)
}

func (s *DirectoryService) GetByName(ctx context.Context, fqn string, params *GetDirectoryByFQNParams) (*Directory, error) {
	return get[Directory](ctx, s.backend, fmt.Sprintf("%s/name/%s", directoryBasePath, fqn), params)
}

func (s *DirectoryService) Create(ctx context.Context, body *CreateDirectory) (*Directory, error) {
	return create[Directory](ctx, s.backend, directoryBasePath, body)
}

func (s *DirectoryService) CreateOrUpdate(ctx context.Context, body *CreateDirectory) (*Directory, error) {
	return put[Directory](ctx, s.backend, directoryBasePath, body)
}

func (s *DirectoryService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Directory, error) {
	return patch[Directory](ctx, s.backend, fmt.Sprintf("%s/%s", directoryBasePath, id), ops)
}

func (s *DirectoryService) Delete(ctx context.Context, id string, params *DeleteDirectoryParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", directoryBasePath, id), params)
}

func (s *DirectoryService) DeleteByName(ctx context.Context, fqn string, params *DeleteDirectoryByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", directoryBasePath, fqn), params)
}

func (s *DirectoryService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", directoryBasePath, id), nil)
}

func (s *DirectoryService) GetVersion(ctx context.Context, id string, version string) (*Directory, error) {
	return get[Directory](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", directoryBasePath, id, version), nil)
}

func (s *DirectoryService) Restore(ctx context.Context, body *RestoreEntity) (*Directory, error) {
	return put[Directory](ctx, s.backend, fmt.Sprintf("%s/restore", directoryBasePath), body)
}
