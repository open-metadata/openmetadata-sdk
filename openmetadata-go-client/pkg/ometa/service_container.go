package ometa

import (
	"context"
	"fmt"
	"iter"
)

const containerBasePath = "containers"

type ContainerService struct {
	backend Backend
}

func (s *ContainerService) List(
	ctx context.Context,
	params *ListContainersParams,
) iter.Seq2[Container, error] {
	return newEntityIterator[Container](ctx, s.backend, containerBasePath, params)
}

func (s *ContainerService) GetByID(ctx context.Context, id string, params *GetContainerByIDParams) (*Container, error) {
	return get[Container](ctx, s.backend, fmt.Sprintf("%s/%s", containerBasePath, id), params)
}

func (s *ContainerService) GetByName(ctx context.Context, fqn string, params *GetContainerByFQNParams) (*Container, error) {
	return get[Container](ctx, s.backend, fmt.Sprintf("%s/name/%s", containerBasePath, fqn), params)
}

func (s *ContainerService) Create(ctx context.Context, body *CreateContainer) (*Container, error) {
	return create[Container](ctx, s.backend, containerBasePath, body)
}

func (s *ContainerService) CreateOrUpdate(ctx context.Context, body *CreateContainer) (*Container, error) {
	return put[Container](ctx, s.backend, containerBasePath, body)
}

func (s *ContainerService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Container, error) {
	return patch[Container](ctx, s.backend, fmt.Sprintf("%s/%s", containerBasePath, id), ops)
}

func (s *ContainerService) Delete(ctx context.Context, id string, params *DeleteContainerParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", containerBasePath, id), params)
}

func (s *ContainerService) DeleteByName(ctx context.Context, fqn string, params *DeleteContainerByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", containerBasePath, fqn), params)
}

func (s *ContainerService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", containerBasePath, id), nil)
}

func (s *ContainerService) GetVersion(ctx context.Context, id string, version string) (*Container, error) {
	return get[Container](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", containerBasePath, id, version), nil)
}

func (s *ContainerService) Restore(ctx context.Context, body *RestoreEntity) (*Container, error) {
	return put[Container](ctx, s.backend, fmt.Sprintf("%s/restore", containerBasePath), body)
}
