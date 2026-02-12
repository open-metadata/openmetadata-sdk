package ometa

import (
	"context"
	"fmt"
	"iter"
)

const dashboardBasePath = "dashboards"

type DashboardSvc struct {
	backend Backend
}

func (s *DashboardSvc) List(
	ctx context.Context,
	params *ListDashboardsParams,
) iter.Seq2[Dashboard, error] {
	return newEntityIterator[Dashboard](ctx, s.backend, dashboardBasePath, params)
}

func (s *DashboardSvc) GetByID(ctx context.Context, id string, params *GetDashboardByIDParams) (*Dashboard, error) {
	return get[Dashboard](ctx, s.backend, fmt.Sprintf("%s/%s", dashboardBasePath, id), params)
}

func (s *DashboardSvc) GetByName(ctx context.Context, fqn string, params *GetDashboardByFQNParams) (*Dashboard, error) {
	return get[Dashboard](ctx, s.backend, fmt.Sprintf("%s/name/%s", dashboardBasePath, fqn), params)
}

func (s *DashboardSvc) Create(ctx context.Context, body *CreateDashboard) (*Dashboard, error) {
	return create[Dashboard](ctx, s.backend, dashboardBasePath, body)
}

func (s *DashboardSvc) CreateOrUpdate(ctx context.Context, body *CreateDashboard) (*Dashboard, error) {
	return put[Dashboard](ctx, s.backend, dashboardBasePath, body)
}

func (s *DashboardSvc) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Dashboard, error) {
	return patch[Dashboard](ctx, s.backend, fmt.Sprintf("%s/%s", dashboardBasePath, id), ops)
}

func (s *DashboardSvc) Delete(ctx context.Context, id string, params *DeleteDashboardParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", dashboardBasePath, id), params)
}

func (s *DashboardSvc) DeleteByName(ctx context.Context, fqn string, params *DeleteDashboardByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", dashboardBasePath, fqn), params)
}

func (s *DashboardSvc) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", dashboardBasePath, id), nil)
}

func (s *DashboardSvc) GetVersion(ctx context.Context, id string, version string) (*Dashboard, error) {
	return get[Dashboard](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", dashboardBasePath, id, version), nil)
}

func (s *DashboardSvc) Restore(ctx context.Context, body *RestoreEntity) (*Dashboard, error) {
	return put[Dashboard](ctx, s.backend, fmt.Sprintf("%s/restore", dashboardBasePath), body)
}
