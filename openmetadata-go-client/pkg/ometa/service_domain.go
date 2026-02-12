package ometa

import (
	"context"
	"fmt"
	"iter"
)

const domainBasePath = "domains"

type DomainService struct {
	backend Backend
}

func (s *DomainService) List(
	ctx context.Context,
	params *ListDomainsParams,
) iter.Seq2[Domain, error] {
	return newEntityIterator[Domain](ctx, s.backend, domainBasePath, params)
}

func (s *DomainService) GetByID(ctx context.Context, id string, params *GetDomainByIDParams) (*Domain, error) {
	return get[Domain](ctx, s.backend, fmt.Sprintf("%s/%s", domainBasePath, id), params)
}

func (s *DomainService) GetByName(ctx context.Context, fqn string, params *GetDomainByFQNParams) (*Domain, error) {
	return get[Domain](ctx, s.backend, fmt.Sprintf("%s/name/%s", domainBasePath, fqn), params)
}

func (s *DomainService) Create(ctx context.Context, body *CreateDomain) (*Domain, error) {
	return create[Domain](ctx, s.backend, domainBasePath, body)
}

func (s *DomainService) CreateOrUpdate(ctx context.Context, body *CreateDomain) (*Domain, error) {
	return put[Domain](ctx, s.backend, domainBasePath, body)
}

func (s *DomainService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Domain, error) {
	return patch[Domain](ctx, s.backend, fmt.Sprintf("%s/%s", domainBasePath, id), ops)
}

func (s *DomainService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", domainBasePath, id), nil)
}

func (s *DomainService) GetVersion(ctx context.Context, id string, version string) (*Domain, error) {
	return get[Domain](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", domainBasePath, id, version), nil)
}

func (s *DomainService) Restore(ctx context.Context, body *RestoreEntity) (*Domain, error) {
	return put[Domain](ctx, s.backend, fmt.Sprintf("%s/restore", domainBasePath), body)
}
