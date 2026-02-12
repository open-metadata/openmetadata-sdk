package ometa

import (
	"context"
	"fmt"
	"iter"
)

const mlModelServiceBasePath = "services/mlmodelServices"

type MlModelServiceService struct {
	backend Backend
}

func (s *MlModelServiceService) List(
	ctx context.Context,
	params *ListMlModelServiceParams,
) iter.Seq2[MlModelService, error] {
	return newEntityIterator[MlModelService](ctx, s.backend, mlModelServiceBasePath, params)
}

func (s *MlModelServiceService) GetByID(ctx context.Context, id string, params *GetMlModelServiceByIDParams) (*MlModelService, error) {
	return get[MlModelService](ctx, s.backend, fmt.Sprintf("%s/%s", mlModelServiceBasePath, id), params)
}

func (s *MlModelServiceService) GetByName(ctx context.Context, fqn string, params *GetMlModelServiceByFQNParams) (*MlModelService, error) {
	return get[MlModelService](ctx, s.backend, fmt.Sprintf("%s/name/%s", mlModelServiceBasePath, fqn), params)
}

func (s *MlModelServiceService) Create(ctx context.Context, body *CreateMlModelService) (*MlModelService, error) {
	return create[MlModelService](ctx, s.backend, mlModelServiceBasePath, body)
}

func (s *MlModelServiceService) CreateOrUpdate(ctx context.Context, body *CreateMlModelService) (*MlModelService, error) {
	return put[MlModelService](ctx, s.backend, mlModelServiceBasePath, body)
}

func (s *MlModelServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*MlModelService, error) {
	return patch[MlModelService](ctx, s.backend, fmt.Sprintf("%s/%s", mlModelServiceBasePath, id), ops)
}

func (s *MlModelServiceService) Delete(ctx context.Context, id string, params *DeleteMlModelServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", mlModelServiceBasePath, id), params)
}

func (s *MlModelServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteMlModelServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", mlModelServiceBasePath, fqn), params)
}

func (s *MlModelServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", mlModelServiceBasePath, id), nil)
}

func (s *MlModelServiceService) GetVersion(ctx context.Context, id string, version string) (*MlModelService, error) {
	return get[MlModelService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", mlModelServiceBasePath, id, version), nil)
}

func (s *MlModelServiceService) Restore(ctx context.Context, body *RestoreEntity) (*MlModelService, error) {
	return put[MlModelService](ctx, s.backend, fmt.Sprintf("%s/restore", mlModelServiceBasePath), body)
}
