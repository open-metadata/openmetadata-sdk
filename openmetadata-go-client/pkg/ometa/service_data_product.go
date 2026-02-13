package ometa

import (
	"context"
	"fmt"
	"iter"
)

const dataProductBasePath = "dataProducts"

type DataProductService struct {
	backend Backend
}

func (s *DataProductService) List(
	ctx context.Context,
	params *ListDataProductsParams,
) iter.Seq2[DataProduct, error] {
	return newEntityIterator[DataProduct](ctx, s.backend, dataProductBasePath, params)
}

func (s *DataProductService) GetByID(ctx context.Context, id string, params *GetDataProductByIDParams) (*DataProduct, error) {
	return get[DataProduct](ctx, s.backend, fmt.Sprintf("%s/%s", dataProductBasePath, id), params)
}

func (s *DataProductService) GetByName(ctx context.Context, fqn string, params *GetDataProductByFQNParams) (*DataProduct, error) {
	return get[DataProduct](ctx, s.backend, fmt.Sprintf("%s/name/%s", dataProductBasePath, fqn), params)
}

func (s *DataProductService) Create(ctx context.Context, body *CreateDataProduct) (*DataProduct, error) {
	return create[DataProduct](ctx, s.backend, dataProductBasePath, body)
}

func (s *DataProductService) CreateOrUpdate(ctx context.Context, body *CreateDataProduct) (*DataProduct, error) {
	return put[DataProduct](ctx, s.backend, dataProductBasePath, body)
}

func (s *DataProductService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*DataProduct, error) {
	return patch[DataProduct](ctx, s.backend, fmt.Sprintf("%s/%s", dataProductBasePath, id), ops)
}

func (s *DataProductService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", dataProductBasePath, id), nil)
}

func (s *DataProductService) GetVersion(ctx context.Context, id string, version string) (*DataProduct, error) {
	return get[DataProduct](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", dataProductBasePath, id, version), nil)
}

func (s *DataProductService) Restore(ctx context.Context, body *RestoreEntity) (*DataProduct, error) {
	return put[DataProduct](ctx, s.backend, fmt.Sprintf("%s/restore", dataProductBasePath), body)
}
