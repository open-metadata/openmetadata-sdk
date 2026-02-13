package ometa

import (
	"context"
	"fmt"
	"iter"
)

const policyBasePath = "policies"

type PolicyService struct {
	backend Backend
}

func (s *PolicyService) List(
	ctx context.Context,
	params *ListPoliciesParams,
) iter.Seq2[Policy, error] {
	return newEntityIterator[Policy](ctx, s.backend, policyBasePath, params)
}

func (s *PolicyService) GetByID(ctx context.Context, id string, params *GetPolicyByIDParams) (*Policy, error) {
	return get[Policy](ctx, s.backend, fmt.Sprintf("%s/%s", policyBasePath, id), params)
}

func (s *PolicyService) GetByName(ctx context.Context, fqn string, params *GetPolicyByFQNParams) (*Policy, error) {
	return get[Policy](ctx, s.backend, fmt.Sprintf("%s/name/%s", policyBasePath, fqn), params)
}

func (s *PolicyService) Create(ctx context.Context, body *CreatePolicy) (*Policy, error) {
	return create[Policy](ctx, s.backend, policyBasePath, body)
}

func (s *PolicyService) CreateOrUpdate(ctx context.Context, body *CreatePolicy) (*Policy, error) {
	return put[Policy](ctx, s.backend, policyBasePath, body)
}

func (s *PolicyService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Policy, error) {
	return patch[Policy](ctx, s.backend, fmt.Sprintf("%s/%s", policyBasePath, id), ops)
}

func (s *PolicyService) Delete(ctx context.Context, id string, params *DeletePolicyParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", policyBasePath, id), params)
}

func (s *PolicyService) DeleteByName(ctx context.Context, fqn string, params *DeletePolicyByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", policyBasePath, fqn), params)
}

func (s *PolicyService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", policyBasePath, id), nil)
}

func (s *PolicyService) GetVersion(ctx context.Context, id string, version string) (*Policy, error) {
	return get[Policy](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", policyBasePath, id, version), nil)
}

func (s *PolicyService) Restore(ctx context.Context, body *RestoreEntity) (*Policy, error) {
	return put[Policy](ctx, s.backend, fmt.Sprintf("%s/restore", policyBasePath), body)
}
