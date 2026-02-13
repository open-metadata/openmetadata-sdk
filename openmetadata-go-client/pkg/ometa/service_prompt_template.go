package ometa

import (
	"context"
	"fmt"
	"iter"
)

const promptTemplateBasePath = "promptTemplates"

type PromptTemplateService struct {
	backend Backend
}

func (s *PromptTemplateService) List(
	ctx context.Context,
	params *ListPromptTemplatesParams,
) iter.Seq2[PromptTemplate, error] {
	return newEntityIterator[PromptTemplate](ctx, s.backend, promptTemplateBasePath, params)
}

func (s *PromptTemplateService) GetByID(ctx context.Context, id string, params *GetPromptTemplateByIDParams) (*PromptTemplate, error) {
	return get[PromptTemplate](ctx, s.backend, fmt.Sprintf("%s/%s", promptTemplateBasePath, id), params)
}

func (s *PromptTemplateService) GetByName(ctx context.Context, fqn string, params *GetPromptTemplateByFQNParams) (*PromptTemplate, error) {
	return get[PromptTemplate](ctx, s.backend, fmt.Sprintf("%s/name/%s", promptTemplateBasePath, fqn), params)
}

func (s *PromptTemplateService) Create(ctx context.Context, body *CreatePromptTemplate) (*PromptTemplate, error) {
	return create[PromptTemplate](ctx, s.backend, promptTemplateBasePath, body)
}

func (s *PromptTemplateService) CreateOrUpdate(ctx context.Context, body *CreatePromptTemplate) (*PromptTemplate, error) {
	return put[PromptTemplate](ctx, s.backend, promptTemplateBasePath, body)
}

func (s *PromptTemplateService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*PromptTemplate, error) {
	return patch[PromptTemplate](ctx, s.backend, fmt.Sprintf("%s/%s", promptTemplateBasePath, id), ops)
}

func (s *PromptTemplateService) Delete(ctx context.Context, id string, params *DeletePromptTemplateParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", promptTemplateBasePath, id), params)
}

func (s *PromptTemplateService) DeleteByName(ctx context.Context, fqn string, params *DeletePromptTemplateByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", promptTemplateBasePath, fqn), params)
}

func (s *PromptTemplateService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", promptTemplateBasePath, id), nil)
}

func (s *PromptTemplateService) GetVersion(ctx context.Context, id string, version string) (*PromptTemplate, error) {
	return get[PromptTemplate](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", promptTemplateBasePath, id, version), nil)
}

func (s *PromptTemplateService) Restore(ctx context.Context, body *RestoreEntity) (*PromptTemplate, error) {
	return put[PromptTemplate](ctx, s.backend, fmt.Sprintf("%s/restore", promptTemplateBasePath), body)
}
