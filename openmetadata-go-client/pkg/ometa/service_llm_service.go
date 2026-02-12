package ometa

import (
	"context"
	"fmt"
	"iter"
)

const llmServiceBasePath = "services/llmServices"

type LLMServiceService struct {
	backend Backend
}

func (s *LLMServiceService) List(
	ctx context.Context,
	params *ListLLMServicesParams,
) iter.Seq2[LLMService, error] {
	return newEntityIterator[LLMService](ctx, s.backend, llmServiceBasePath, params)
}

func (s *LLMServiceService) GetByID(ctx context.Context, id string, params *GetLLMServiceByIDParams) (*LLMService, error) {
	return get[LLMService](ctx, s.backend, fmt.Sprintf("%s/%s", llmServiceBasePath, id), params)
}

func (s *LLMServiceService) GetByName(ctx context.Context, fqn string, params *GetLLMServiceByFQNParams) (*LLMService, error) {
	return get[LLMService](ctx, s.backend, fmt.Sprintf("%s/name/%s", llmServiceBasePath, fqn), params)
}

func (s *LLMServiceService) Create(ctx context.Context, body *CreateLLMService) (*LLMService, error) {
	return create[LLMService](ctx, s.backend, llmServiceBasePath, body)
}

func (s *LLMServiceService) CreateOrUpdate(ctx context.Context, body *CreateLLMService) (*LLMService, error) {
	return put[LLMService](ctx, s.backend, llmServiceBasePath, body)
}

func (s *LLMServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*LLMService, error) {
	return patch[LLMService](ctx, s.backend, fmt.Sprintf("%s/%s", llmServiceBasePath, id), ops)
}

func (s *LLMServiceService) Delete(ctx context.Context, id string, params *DeleteLLMServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", llmServiceBasePath, id), params)
}

func (s *LLMServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteLLMServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", llmServiceBasePath, fqn), params)
}

func (s *LLMServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", llmServiceBasePath, id), nil)
}

func (s *LLMServiceService) GetVersion(ctx context.Context, id string, version string) (*LLMService, error) {
	return get[LLMService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", llmServiceBasePath, id, version), nil)
}

func (s *LLMServiceService) Restore(ctx context.Context, body *RestoreEntity) (*LLMService, error) {
	return put[LLMService](ctx, s.backend, fmt.Sprintf("%s/restore", llmServiceBasePath), body)
}
