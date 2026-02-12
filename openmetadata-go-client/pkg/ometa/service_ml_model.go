package ometa

import (
	"context"
	"fmt"
	"iter"
)

const mlModelBasePath = "mlmodels"

type MlModelSvc struct {
	backend Backend
}

func (s *MlModelSvc) List(
	ctx context.Context,
	params *ListMlModelsParams,
) iter.Seq2[MlModel, error] {
	return newEntityIterator[MlModel](ctx, s.backend, mlModelBasePath, params)
}

func (s *MlModelSvc) GetByID(ctx context.Context, id string, params *GetMlModelByIDParams) (*MlModel, error) {
	return get[MlModel](ctx, s.backend, fmt.Sprintf("%s/%s", mlModelBasePath, id), params)
}

func (s *MlModelSvc) GetByName(ctx context.Context, fqn string, params *GetMlModelByFQNParams) (*MlModel, error) {
	return get[MlModel](ctx, s.backend, fmt.Sprintf("%s/name/%s", mlModelBasePath, fqn), params)
}

func (s *MlModelSvc) Create(ctx context.Context, body *CreateMlModel) (*MlModel, error) {
	return create[MlModel](ctx, s.backend, mlModelBasePath, body)
}

func (s *MlModelSvc) CreateOrUpdate(ctx context.Context, body *CreateMlModel) (*MlModel, error) {
	return put[MlModel](ctx, s.backend, mlModelBasePath, body)
}

func (s *MlModelSvc) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*MlModel, error) {
	return patch[MlModel](ctx, s.backend, fmt.Sprintf("%s/%s", mlModelBasePath, id), ops)
}

func (s *MlModelSvc) Delete(ctx context.Context, id string, params *DeleteMlModelParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", mlModelBasePath, id), params)
}

func (s *MlModelSvc) DeleteByName(ctx context.Context, fqn string, params *DeleteMlModelByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", mlModelBasePath, fqn), params)
}

func (s *MlModelSvc) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", mlModelBasePath, id), nil)
}

func (s *MlModelSvc) GetVersion(ctx context.Context, id string, version string) (*MlModel, error) {
	return get[MlModel](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", mlModelBasePath, id, version), nil)
}

func (s *MlModelSvc) Restore(ctx context.Context, body *RestoreEntity) (*MlModel, error) {
	return put[MlModel](ctx, s.backend, fmt.Sprintf("%s/restore", mlModelBasePath), body)
}
