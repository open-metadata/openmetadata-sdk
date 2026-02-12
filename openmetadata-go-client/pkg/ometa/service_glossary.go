package ometa

import (
	"context"
	"fmt"
	"iter"
)

const glossaryBasePath = "glossaries"

type GlossaryService struct {
	backend Backend
}

func (s *GlossaryService) List(
	ctx context.Context,
	params *ListGlossariesParams,
) iter.Seq2[Glossary, error] {
	return newEntityIterator[Glossary](ctx, s.backend, glossaryBasePath, params)
}

func (s *GlossaryService) GetByID(ctx context.Context, id string, params *GetGlossaryByIDParams) (*Glossary, error) {
	return get[Glossary](ctx, s.backend, fmt.Sprintf("%s/%s", glossaryBasePath, id), params)
}

func (s *GlossaryService) GetByName(ctx context.Context, fqn string, params *GetGlossaryByFQNParams) (*Glossary, error) {
	return get[Glossary](ctx, s.backend, fmt.Sprintf("%s/name/%s", glossaryBasePath, fqn), params)
}

func (s *GlossaryService) Create(ctx context.Context, body *CreateGlossary) (*Glossary, error) {
	return create[Glossary](ctx, s.backend, glossaryBasePath, body)
}

func (s *GlossaryService) CreateOrUpdate(ctx context.Context, body *CreateGlossary) (*Glossary, error) {
	return put[Glossary](ctx, s.backend, glossaryBasePath, body)
}

func (s *GlossaryService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Glossary, error) {
	return patch[Glossary](ctx, s.backend, fmt.Sprintf("%s/%s", glossaryBasePath, id), ops)
}

func (s *GlossaryService) Delete(ctx context.Context, id string, params *DeleteGlossaryParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", glossaryBasePath, id), params)
}

func (s *GlossaryService) DeleteByName(ctx context.Context, fqn string, params *DeleteGlossaryByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", glossaryBasePath, fqn), params)
}

func (s *GlossaryService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", glossaryBasePath, id), nil)
}

func (s *GlossaryService) GetVersion(ctx context.Context, id string, version string) (*Glossary, error) {
	return get[Glossary](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", glossaryBasePath, id, version), nil)
}

func (s *GlossaryService) Restore(ctx context.Context, body *RestoreEntity) (*Glossary, error) {
	return put[Glossary](ctx, s.backend, fmt.Sprintf("%s/restore", glossaryBasePath), body)
}
