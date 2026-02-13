package ometa

import (
	"context"
	"fmt"
	"iter"
)

const aiGovernancePolicyBasePath = "aiGovernancePolicies"

type AIGovernancePolicyService struct {
	backend Backend
}

func (s *AIGovernancePolicyService) List(
	ctx context.Context,
	params *ListAIGovernancePoliciesParams,
) iter.Seq2[AIGovernancePolicy, error] {
	return newEntityIterator[AIGovernancePolicy](ctx, s.backend, aiGovernancePolicyBasePath, params)
}

func (s *AIGovernancePolicyService) GetByID(ctx context.Context, id string, params *GetAIGovernancePolicyByIDParams) (*AIGovernancePolicy, error) {
	return get[AIGovernancePolicy](ctx, s.backend, fmt.Sprintf("%s/%s", aiGovernancePolicyBasePath, id), params)
}

func (s *AIGovernancePolicyService) GetByName(ctx context.Context, fqn string, params *GetAIGovernancePolicyByFQNParams) (*AIGovernancePolicy, error) {
	return get[AIGovernancePolicy](ctx, s.backend, fmt.Sprintf("%s/name/%s", aiGovernancePolicyBasePath, fqn), params)
}

func (s *AIGovernancePolicyService) Create(ctx context.Context, body *CreateAIGovernancePolicy) (*AIGovernancePolicy, error) {
	return create[AIGovernancePolicy](ctx, s.backend, aiGovernancePolicyBasePath, body)
}

func (s *AIGovernancePolicyService) CreateOrUpdate(ctx context.Context, body *CreateAIGovernancePolicy) (*AIGovernancePolicy, error) {
	return put[AIGovernancePolicy](ctx, s.backend, aiGovernancePolicyBasePath, body)
}

func (s *AIGovernancePolicyService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*AIGovernancePolicy, error) {
	return patch[AIGovernancePolicy](ctx, s.backend, fmt.Sprintf("%s/%s", aiGovernancePolicyBasePath, id), ops)
}

func (s *AIGovernancePolicyService) Delete(ctx context.Context, id string, params *DeleteAIGovernancePolicyParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", aiGovernancePolicyBasePath, id), params)
}

func (s *AIGovernancePolicyService) DeleteByName(ctx context.Context, fqn string, params *DeleteAIGovernancePolicyByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", aiGovernancePolicyBasePath, fqn), params)
}

func (s *AIGovernancePolicyService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", aiGovernancePolicyBasePath, id), nil)
}

func (s *AIGovernancePolicyService) GetVersion(ctx context.Context, id string, version string) (*AIGovernancePolicy, error) {
	return get[AIGovernancePolicy](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", aiGovernancePolicyBasePath, id, version), nil)
}

func (s *AIGovernancePolicyService) Restore(ctx context.Context, body *RestoreEntity) (*AIGovernancePolicy, error) {
	return put[AIGovernancePolicy](ctx, s.backend, fmt.Sprintf("%s/restore", aiGovernancePolicyBasePath), body)
}
