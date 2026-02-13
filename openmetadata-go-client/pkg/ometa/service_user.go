package ometa

import (
	"context"
	"fmt"
	"iter"
)

const userBasePath = "users"

type UserService struct {
	backend Backend
}

func (s *UserService) List(
	ctx context.Context,
	params *ListUsersParams,
) iter.Seq2[User, error] {
	return newEntityIterator[User](ctx, s.backend, userBasePath, params)
}

func (s *UserService) GetByID(ctx context.Context, id string, params *GetUserByIDParams) (*User, error) {
	return get[User](ctx, s.backend, fmt.Sprintf("%s/%s", userBasePath, id), params)
}

func (s *UserService) GetByName(ctx context.Context, fqn string, params *GetUserByFQNParams) (*User, error) {
	return get[User](ctx, s.backend, fmt.Sprintf("%s/name/%s", userBasePath, fqn), params)
}

func (s *UserService) Create(ctx context.Context, body *CreateUser) (*User, error) {
	return create[User](ctx, s.backend, userBasePath, body)
}

func (s *UserService) CreateOrUpdate(ctx context.Context, body *CreateUser) (*User, error) {
	return put[User](ctx, s.backend, userBasePath, body)
}

func (s *UserService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*User, error) {
	return patch[User](ctx, s.backend, fmt.Sprintf("%s/%s", userBasePath, id), ops)
}

func (s *UserService) Delete(ctx context.Context, id string, params *DeleteUserParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", userBasePath, id), params)
}

func (s *UserService) DeleteByName(ctx context.Context, fqn string, params *DeleteUserByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", userBasePath, fqn), params)
}

func (s *UserService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", userBasePath, id), nil)
}

func (s *UserService) GetVersion(ctx context.Context, id string, version string) (*User, error) {
	return get[User](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", userBasePath, id, version), nil)
}

func (s *UserService) Restore(ctx context.Context, body *RestoreEntity) (*User, error) {
	return put[User](ctx, s.backend, fmt.Sprintf("%s/restore", userBasePath), body)
}
