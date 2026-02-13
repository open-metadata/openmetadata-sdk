package ometa

import (
	"context"
	"fmt"
	"iter"
)

const metricBasePath = "metrics"

type MetricService struct {
	backend Backend
}

func (s *MetricService) List(
	ctx context.Context,
	params *ListMetricsParams,
) iter.Seq2[Metric, error] {
	return newEntityIterator[Metric](ctx, s.backend, metricBasePath, params)
}

func (s *MetricService) GetByID(ctx context.Context, id string, params *GetMetricByIDParams) (*Metric, error) {
	return get[Metric](ctx, s.backend, fmt.Sprintf("%s/%s", metricBasePath, id), params)
}

func (s *MetricService) GetByName(ctx context.Context, fqn string, params *GetMetricByFQNParams) (*Metric, error) {
	return get[Metric](ctx, s.backend, fmt.Sprintf("%s/name/%s", metricBasePath, fqn), params)
}

func (s *MetricService) Create(ctx context.Context, body *CreateMetric) (*Metric, error) {
	return create[Metric](ctx, s.backend, metricBasePath, body)
}

func (s *MetricService) CreateOrUpdate(ctx context.Context, body *CreateMetric) (*Metric, error) {
	return put[Metric](ctx, s.backend, metricBasePath, body)
}

func (s *MetricService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Metric, error) {
	return patch[Metric](ctx, s.backend, fmt.Sprintf("%s/%s", metricBasePath, id), ops)
}

func (s *MetricService) Delete(ctx context.Context, id string, params *DeleteMetricParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", metricBasePath, id), params)
}

func (s *MetricService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", metricBasePath, id), nil)
}

func (s *MetricService) GetVersion(ctx context.Context, id string, version string) (*Metric, error) {
	return get[Metric](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", metricBasePath, id, version), nil)
}

func (s *MetricService) Restore(ctx context.Context, body *RestoreEntity) (*Metric, error) {
	return put[Metric](ctx, s.backend, fmt.Sprintf("%s/restore", metricBasePath), body)
}
