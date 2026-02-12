package ometa

import (
	"context"
	"fmt"
	"iter"
)

const chartBasePath = "charts"

type ChartService struct {
	backend Backend
}

func (s *ChartService) List(
	ctx context.Context,
	params *ListChartsParams,
) iter.Seq2[Chart, error] {
	return newEntityIterator[Chart](ctx, s.backend, chartBasePath, params)
}

func (s *ChartService) GetByID(ctx context.Context, id string, params *GetChartByIDParams) (*Chart, error) {
	return get[Chart](ctx, s.backend, fmt.Sprintf("%s/%s", chartBasePath, id), params)
}

func (s *ChartService) GetByName(ctx context.Context, fqn string, params *GetChartByFQNParams) (*Chart, error) {
	return get[Chart](ctx, s.backend, fmt.Sprintf("%s/name/%s", chartBasePath, fqn), params)
}

func (s *ChartService) Create(ctx context.Context, body *CreateChart) (*Chart, error) {
	return create[Chart](ctx, s.backend, chartBasePath, body)
}

func (s *ChartService) CreateOrUpdate(ctx context.Context, body *CreateChart) (*Chart, error) {
	return put[Chart](ctx, s.backend, chartBasePath, body)
}

func (s *ChartService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Chart, error) {
	return patch[Chart](ctx, s.backend, fmt.Sprintf("%s/%s", chartBasePath, id), ops)
}

func (s *ChartService) Delete(ctx context.Context, id string, params *DeleteChartParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", chartBasePath, id), params)
}

func (s *ChartService) DeleteByName(ctx context.Context, fqn string, params *DeleteChartByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", chartBasePath, fqn), params)
}

func (s *ChartService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", chartBasePath, id), nil)
}

func (s *ChartService) GetVersion(ctx context.Context, id string, version string) (*Chart, error) {
	return get[Chart](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", chartBasePath, id, version), nil)
}

func (s *ChartService) Restore(ctx context.Context, body *RestoreEntity) (*Chart, error) {
	return put[Chart](ctx, s.backend, fmt.Sprintf("%s/restore", chartBasePath), body)
}
