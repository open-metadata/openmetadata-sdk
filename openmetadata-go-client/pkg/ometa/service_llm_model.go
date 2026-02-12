package ometa

import (
	"context"
	"fmt"
	"iter"
)

const llmModelBasePath = "llmModels"

type LLMModelService struct {
	backend Backend
}

func (s *LLMModelService) List(
	ctx context.Context,
	params *ListLLMModelsParams,
) iter.Seq2[LLMModel, error] {
	return newEntityIterator[LLMModel](ctx, s.backend, llmModelBasePath, params)
}

func (s *LLMModelService) GetByID(ctx context.Context, id string, params *GetLLMModelByIDParams) (*LLMModel, error) {
	return get[LLMModel](ctx, s.backend, fmt.Sprintf("%s/%s", llmModelBasePath, id), params)
}

func (s *LLMModelService) GetByName(ctx context.Context, fqn string, params *GetLLMModelByFQNParams) (*LLMModel, error) {
	return get[LLMModel](ctx, s.backend, fmt.Sprintf("%s/name/%s", llmModelBasePath, fqn), params)
}

func (s *LLMModelService) Create(ctx context.Context, body *CreateLLMModel) (*LLMModel, error) {
	return create[LLMModel](ctx, s.backend, llmModelBasePath, body)
}

func (s *LLMModelService) CreateOrUpdate(ctx context.Context, body *CreateLLMModel) (*LLMModel, error) {
	return put[LLMModel](ctx, s.backend, llmModelBasePath, body)
}

func (s *LLMModelService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*LLMModel, error) {
	return patch[LLMModel](ctx, s.backend, fmt.Sprintf("%s/%s", llmModelBasePath, id), ops)
}

func (s *LLMModelService) Delete(ctx context.Context, id string, params *DeleteLLMModelParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", llmModelBasePath, id), params)
}

func (s *LLMModelService) DeleteByName(ctx context.Context, fqn string, params *DeleteLLMModelByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", llmModelBasePath, fqn), params)
}

func (s *LLMModelService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", llmModelBasePath, id), nil)
}

func (s *LLMModelService) GetVersion(ctx context.Context, id string, version string) (*LLMModel, error) {
	return get[LLMModel](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", llmModelBasePath, id, version), nil)
}

func (s *LLMModelService) Restore(ctx context.Context, body *RestoreEntity) (*LLMModel, error) {
	return put[LLMModel](ctx, s.backend, fmt.Sprintf("%s/restore", llmModelBasePath), body)
}
