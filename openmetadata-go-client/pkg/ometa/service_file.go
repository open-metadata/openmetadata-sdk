package ometa

import (
	"context"
	"fmt"
	"iter"
)

const fileBasePath = "drives/files"

type FileService struct {
	backend Backend
}

func (s *FileService) List(
	ctx context.Context,
	params *ListFilesParams,
) iter.Seq2[File, error] {
	return newEntityIterator[File](ctx, s.backend, fileBasePath, params)
}

func (s *FileService) GetByID(ctx context.Context, id string, params *GetFileByIDParams) (*File, error) {
	return get[File](ctx, s.backend, fmt.Sprintf("%s/%s", fileBasePath, id), params)
}

func (s *FileService) GetByName(ctx context.Context, fqn string, params *GetFileByFQNParams) (*File, error) {
	return get[File](ctx, s.backend, fmt.Sprintf("%s/name/%s", fileBasePath, fqn), params)
}

func (s *FileService) Create(ctx context.Context, body *CreateFile) (*File, error) {
	return create[File](ctx, s.backend, fileBasePath, body)
}

func (s *FileService) CreateOrUpdate(ctx context.Context, body *CreateFile) (*File, error) {
	return put[File](ctx, s.backend, fileBasePath, body)
}

func (s *FileService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*File, error) {
	return patch[File](ctx, s.backend, fmt.Sprintf("%s/%s", fileBasePath, id), ops)
}

func (s *FileService) Delete(ctx context.Context, id string, params *DeleteFileParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", fileBasePath, id), params)
}

func (s *FileService) DeleteByName(ctx context.Context, fqn string, params *DeleteFileByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", fileBasePath, fqn), params)
}

func (s *FileService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", fileBasePath, id), nil)
}

func (s *FileService) GetVersion(ctx context.Context, id string, version string) (*File, error) {
	return get[File](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", fileBasePath, id, version), nil)
}

func (s *FileService) Restore(ctx context.Context, body *RestoreEntity) (*File, error) {
	return put[File](ctx, s.backend, fmt.Sprintf("%s/restore", fileBasePath), body)
}
