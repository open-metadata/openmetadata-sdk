package ometa

import (
	"context"
	"fmt"
	"iter"
)

const driveServiceBasePath = "services/driveServices"

type DriveServiceService struct {
	backend Backend
}

func (s *DriveServiceService) List(
	ctx context.Context,
	params *ListDriveServicesParams,
) iter.Seq2[DriveService, error] {
	return newEntityIterator[DriveService](ctx, s.backend, driveServiceBasePath, params)
}

func (s *DriveServiceService) GetByID(ctx context.Context, id string, params *GetDriveServiceByIDParams) (*DriveService, error) {
	return get[DriveService](ctx, s.backend, fmt.Sprintf("%s/%s", driveServiceBasePath, id), params)
}

func (s *DriveServiceService) GetByName(ctx context.Context, fqn string, params *GetDriveServiceByFQNParams) (*DriveService, error) {
	return get[DriveService](ctx, s.backend, fmt.Sprintf("%s/name/%s", driveServiceBasePath, fqn), params)
}

func (s *DriveServiceService) Create(ctx context.Context, body *CreateDriveService) (*DriveService, error) {
	return create[DriveService](ctx, s.backend, driveServiceBasePath, body)
}

func (s *DriveServiceService) CreateOrUpdate(ctx context.Context, body *CreateDriveService) (*DriveService, error) {
	return put[DriveService](ctx, s.backend, driveServiceBasePath, body)
}

func (s *DriveServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*DriveService, error) {
	return patch[DriveService](ctx, s.backend, fmt.Sprintf("%s/%s", driveServiceBasePath, id), ops)
}

func (s *DriveServiceService) Delete(ctx context.Context, id string, params *DeleteDriveServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", driveServiceBasePath, id), params)
}

func (s *DriveServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteDriveServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", driveServiceBasePath, fqn), params)
}

func (s *DriveServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", driveServiceBasePath, id), nil)
}

func (s *DriveServiceService) GetVersion(ctx context.Context, id string, version string) (*DriveService, error) {
	return get[DriveService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", driveServiceBasePath, id, version), nil)
}

func (s *DriveServiceService) Restore(ctx context.Context, body *RestoreEntity) (*DriveService, error) {
	return put[DriveService](ctx, s.backend, fmt.Sprintf("%s/restore", driveServiceBasePath), body)
}
