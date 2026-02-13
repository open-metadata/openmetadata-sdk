package ometa

import (
	"context"
	"fmt"
	"iter"
)

const glossaryTermBasePath = "glossaryTerms"

type GlossaryTermService struct {
	backend Backend
}

func (s *GlossaryTermService) List(
	ctx context.Context,
	params *ListGlossaryTermParams,
) iter.Seq2[GlossaryTerm, error] {
	return newEntityIterator[GlossaryTerm](ctx, s.backend, glossaryTermBasePath, params)
}

func (s *GlossaryTermService) GetByID(ctx context.Context, id string, params *GetGlossaryTermByIDParams) (*GlossaryTerm, error) {
	return get[GlossaryTerm](ctx, s.backend, fmt.Sprintf("%s/%s", glossaryTermBasePath, id), params)
}

func (s *GlossaryTermService) GetByName(ctx context.Context, fqn string, params *GetGlossaryTermByFQNParams) (*GlossaryTerm, error) {
	return get[GlossaryTerm](ctx, s.backend, fmt.Sprintf("%s/name/%s", glossaryTermBasePath, fqn), params)
}

func (s *GlossaryTermService) Create(ctx context.Context, body *CreateGlossaryTerm) (*GlossaryTerm, error) {
	return create[GlossaryTerm](ctx, s.backend, glossaryTermBasePath, body)
}

func (s *GlossaryTermService) CreateOrUpdate(ctx context.Context, body *CreateGlossaryTerm) (*GlossaryTerm, error) {
	return put[GlossaryTerm](ctx, s.backend, glossaryTermBasePath, body)
}

func (s *GlossaryTermService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*GlossaryTerm, error) {
	return patch[GlossaryTerm](ctx, s.backend, fmt.Sprintf("%s/%s", glossaryTermBasePath, id), ops)
}

func (s *GlossaryTermService) Delete(ctx context.Context, id string, params *DeleteParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", glossaryTermBasePath, id), params)
}

func (s *GlossaryTermService) DeleteByName(ctx context.Context, fqn string, params *DeleteGlossaryTermByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", glossaryTermBasePath, fqn), params)
}

func (s *GlossaryTermService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", glossaryTermBasePath, id), nil)
}

func (s *GlossaryTermService) GetVersion(ctx context.Context, id string, version string) (*GlossaryTerm, error) {
	return get[GlossaryTerm](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", glossaryTermBasePath, id, version), nil)
}

func (s *GlossaryTermService) Restore(ctx context.Context, body *RestoreEntity) (*GlossaryTerm, error) {
	return put[GlossaryTerm](ctx, s.backend, fmt.Sprintf("%s/restore", glossaryTermBasePath), body)
}
