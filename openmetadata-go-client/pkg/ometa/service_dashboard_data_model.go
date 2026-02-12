package ometa

import (
	"context"
	"fmt"
	"iter"
)

const dashboardDataModelBasePath = "dashboard/datamodels"

type DashboardDataModelService struct {
	backend Backend
}

func (s *DashboardDataModelService) List(
	ctx context.Context,
	params *ListDashboardDataModelsParams,
) iter.Seq2[DashboardDataModel, error] {
	return newEntityIterator[DashboardDataModel](ctx, s.backend, dashboardDataModelBasePath, params)
}

func (s *DashboardDataModelService) GetByID(ctx context.Context, id string, params *GetDataModelByIDParams) (*DashboardDataModel, error) {
	return get[DashboardDataModel](ctx, s.backend, fmt.Sprintf("%s/%s", dashboardDataModelBasePath, id), params)
}

func (s *DashboardDataModelService) GetByName(ctx context.Context, fqn string, params *GetDataModelByFQNParams) (*DashboardDataModel, error) {
	return get[DashboardDataModel](ctx, s.backend, fmt.Sprintf("%s/name/%s", dashboardDataModelBasePath, fqn), params)
}

func (s *DashboardDataModelService) Create(ctx context.Context, body *CreateDashboardDataModel) (*DashboardDataModel, error) {
	return create[DashboardDataModel](ctx, s.backend, dashboardDataModelBasePath, body)
}

func (s *DashboardDataModelService) CreateOrUpdate(ctx context.Context, body *CreateDashboardDataModel) (*DashboardDataModel, error) {
	return put[DashboardDataModel](ctx, s.backend, dashboardDataModelBasePath, body)
}

func (s *DashboardDataModelService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*DashboardDataModel, error) {
	return patch[DashboardDataModel](ctx, s.backend, fmt.Sprintf("%s/%s", dashboardDataModelBasePath, id), ops)
}

func (s *DashboardDataModelService) Delete(ctx context.Context, id string, params *DeleteDataModelParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", dashboardDataModelBasePath, id), params)
}

func (s *DashboardDataModelService) DeleteByName(ctx context.Context, fqn string, params *DeleteDataModelByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", dashboardDataModelBasePath, fqn), params)
}

func (s *DashboardDataModelService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", dashboardDataModelBasePath, id), nil)
}

func (s *DashboardDataModelService) GetVersion(ctx context.Context, id string, version string) (*DashboardDataModel, error) {
	return get[DashboardDataModel](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", dashboardDataModelBasePath, id, version), nil)
}

func (s *DashboardDataModelService) Restore(ctx context.Context, body *RestoreEntity) (*DashboardDataModel, error) {
	return put[DashboardDataModel](ctx, s.backend, fmt.Sprintf("%s/restore", dashboardDataModelBasePath), body)
}
