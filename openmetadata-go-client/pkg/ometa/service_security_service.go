package ometa

import (
	"context"
	"fmt"
	"iter"
)

const securityServiceBasePath = "services/securityServices"

type SecurityServiceService struct {
	backend Backend
}

func (s *SecurityServiceService) List(
	ctx context.Context,
	params *ListSecurityServicesParams,
) iter.Seq2[SecurityService, error] {
	return newEntityIterator[SecurityService](ctx, s.backend, securityServiceBasePath, params)
}

func (s *SecurityServiceService) GetByID(ctx context.Context, id string, params *GetSecurityServiceByIDParams) (*SecurityService, error) {
	return get[SecurityService](ctx, s.backend, fmt.Sprintf("%s/%s", securityServiceBasePath, id), params)
}

func (s *SecurityServiceService) GetByName(ctx context.Context, fqn string, params *GetSecurityServiceByFQNParams) (*SecurityService, error) {
	return get[SecurityService](ctx, s.backend, fmt.Sprintf("%s/name/%s", securityServiceBasePath, fqn), params)
}

func (s *SecurityServiceService) Create(ctx context.Context, body *CreateSecurityService) (*SecurityService, error) {
	return create[SecurityService](ctx, s.backend, securityServiceBasePath, body)
}

func (s *SecurityServiceService) CreateOrUpdate(ctx context.Context, body *CreateSecurityService) (*SecurityService, error) {
	return put[SecurityService](ctx, s.backend, securityServiceBasePath, body)
}

func (s *SecurityServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*SecurityService, error) {
	return patch[SecurityService](ctx, s.backend, fmt.Sprintf("%s/%s", securityServiceBasePath, id), ops)
}

func (s *SecurityServiceService) Delete(ctx context.Context, id string, params *DeleteSecurityServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", securityServiceBasePath, id), params)
}

func (s *SecurityServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteSecurityServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", securityServiceBasePath, fqn), params)
}

func (s *SecurityServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", securityServiceBasePath, id), nil)
}

func (s *SecurityServiceService) GetVersion(ctx context.Context, id string, version string) (*SecurityService, error) {
	return get[SecurityService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", securityServiceBasePath, id, version), nil)
}

func (s *SecurityServiceService) Restore(ctx context.Context, body *RestoreEntity) (*SecurityService, error) {
	return put[SecurityService](ctx, s.backend, fmt.Sprintf("%s/restore", securityServiceBasePath), body)
}
