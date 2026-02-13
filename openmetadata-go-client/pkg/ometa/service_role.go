package ometa

import (
	"context"
	"fmt"
	"iter"
)

const roleBasePath = "roles"

type RoleService struct {
	backend Backend
}

func (s *RoleService) List(
	ctx context.Context,
	params *ListRolesParams,
) iter.Seq2[Role, error] {
	return newEntityIterator[Role](ctx, s.backend, roleBasePath, params)
}

func (s *RoleService) GetByID(ctx context.Context, id string, params *GetRoleByIDParams) (*Role, error) {
	return get[Role](ctx, s.backend, fmt.Sprintf("%s/%s", roleBasePath, id), params)
}

func (s *RoleService) GetByName(ctx context.Context, fqn string, params *GetRoleByFQNParams) (*Role, error) {
	return get[Role](ctx, s.backend, fmt.Sprintf("%s/name/%s", roleBasePath, fqn), params)
}

func (s *RoleService) Create(ctx context.Context, body *CreateRole) (*Role, error) {
	return create[Role](ctx, s.backend, roleBasePath, body)
}

func (s *RoleService) CreateOrUpdate(ctx context.Context, body *CreateRole) (*Role, error) {
	return put[Role](ctx, s.backend, roleBasePath, body)
}

func (s *RoleService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Role, error) {
	return patch[Role](ctx, s.backend, fmt.Sprintf("%s/%s", roleBasePath, id), ops)
}

func (s *RoleService) Delete(ctx context.Context, id string, params *DeleteRoleParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", roleBasePath, id), params)
}

func (s *RoleService) DeleteByName(ctx context.Context, fqn string, params *DeleteRoleByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", roleBasePath, fqn), params)
}

func (s *RoleService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", roleBasePath, id), nil)
}

func (s *RoleService) GetVersion(ctx context.Context, id string, version string) (*Role, error) {
	return get[Role](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", roleBasePath, id, version), nil)
}

func (s *RoleService) Restore(ctx context.Context, body *RestoreEntity) (*Role, error) {
	return put[Role](ctx, s.backend, fmt.Sprintf("%s/restore", roleBasePath), body)
}
