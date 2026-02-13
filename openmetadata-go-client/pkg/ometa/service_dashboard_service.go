package ometa

import (
	"context"
	"fmt"
	"iter"
)

const dashboardServiceBasePath = "services/dashboardServices"

type DashboardServiceService struct {
	backend Backend
}

func (s *DashboardServiceService) List(
	ctx context.Context,
	params *ListDashboardsServiceParams,
) iter.Seq2[DashboardService, error] {
	return newEntityIterator[DashboardService](ctx, s.backend, dashboardServiceBasePath, params)
}

func (s *DashboardServiceService) GetByID(ctx context.Context, id string, params *GetDashboardServiceByIDParams) (*DashboardService, error) {
	return get[DashboardService](ctx, s.backend, fmt.Sprintf("%s/%s", dashboardServiceBasePath, id), params)
}

func (s *DashboardServiceService) GetByName(ctx context.Context, fqn string, params *GetDashboardServiceByFQNParams) (*DashboardService, error) {
	return get[DashboardService](ctx, s.backend, fmt.Sprintf("%s/name/%s", dashboardServiceBasePath, fqn), params)
}

func (s *DashboardServiceService) Create(ctx context.Context, body *CreateDashboardService) (*DashboardService, error) {
	return create[DashboardService](ctx, s.backend, dashboardServiceBasePath, body)
}

func (s *DashboardServiceService) CreateOrUpdate(ctx context.Context, body *CreateDashboardService) (*DashboardService, error) {
	return put[DashboardService](ctx, s.backend, dashboardServiceBasePath, body)
}

func (s *DashboardServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*DashboardService, error) {
	return patch[DashboardService](ctx, s.backend, fmt.Sprintf("%s/%s", dashboardServiceBasePath, id), ops)
}

func (s *DashboardServiceService) Delete(ctx context.Context, id string, params *DeleteDashboardServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", dashboardServiceBasePath, id), params)
}

func (s *DashboardServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteDashboardServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", dashboardServiceBasePath, fqn), params)
}

func (s *DashboardServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", dashboardServiceBasePath, id), nil)
}

func (s *DashboardServiceService) GetVersion(ctx context.Context, id string, version string) (*DashboardService, error) {
	return get[DashboardService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", dashboardServiceBasePath, id, version), nil)
}

func (s *DashboardServiceService) Restore(ctx context.Context, body *RestoreEntity) (*DashboardService, error) {
	return put[DashboardService](ctx, s.backend, fmt.Sprintf("%s/restore", dashboardServiceBasePath), body)
}
