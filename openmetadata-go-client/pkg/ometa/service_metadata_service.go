package ometa

import (
	"context"
	"fmt"
	"iter"
)

const metadataServiceBasePath = "services/metadataServices"

type MetadataServiceService struct {
	backend Backend
}

func (s *MetadataServiceService) List(
	ctx context.Context,
	params *ListMetadataServicesParams,
) iter.Seq2[MetadataService, error] {
	return newEntityIterator[MetadataService](ctx, s.backend, metadataServiceBasePath, params)
}

func (s *MetadataServiceService) GetByID(ctx context.Context, id string, params *GetMetadataServiceByIDParams) (*MetadataService, error) {
	return get[MetadataService](ctx, s.backend, fmt.Sprintf("%s/%s", metadataServiceBasePath, id), params)
}

func (s *MetadataServiceService) GetByName(ctx context.Context, fqn string, params *GetMetadataServiceByFQNParams) (*MetadataService, error) {
	return get[MetadataService](ctx, s.backend, fmt.Sprintf("%s/name/%s", metadataServiceBasePath, fqn), params)
}

func (s *MetadataServiceService) Create(ctx context.Context, body *CreateMetadataService) (*MetadataService, error) {
	return create[MetadataService](ctx, s.backend, metadataServiceBasePath, body)
}

func (s *MetadataServiceService) CreateOrUpdate(ctx context.Context, body *CreateMetadataService) (*MetadataService, error) {
	return put[MetadataService](ctx, s.backend, metadataServiceBasePath, body)
}

func (s *MetadataServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*MetadataService, error) {
	return patch[MetadataService](ctx, s.backend, fmt.Sprintf("%s/%s", metadataServiceBasePath, id), ops)
}

func (s *MetadataServiceService) Delete(ctx context.Context, id string, params *DeleteMetadataServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", metadataServiceBasePath, id), params)
}

func (s *MetadataServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteMetadataServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", metadataServiceBasePath, fqn), params)
}

func (s *MetadataServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", metadataServiceBasePath, id), nil)
}

func (s *MetadataServiceService) GetVersion(ctx context.Context, id string, version string) (*MetadataService, error) {
	return get[MetadataService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", metadataServiceBasePath, id, version), nil)
}

func (s *MetadataServiceService) Restore(ctx context.Context, body *RestoreEntity) (*MetadataService, error) {
	return put[MetadataService](ctx, s.backend, fmt.Sprintf("%s/restore", metadataServiceBasePath), body)
}
