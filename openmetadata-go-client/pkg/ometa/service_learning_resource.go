package ometa

import (
	"context"
	"fmt"
	"iter"
)

const learningResourceBasePath = "learning/resources"

type LearningResourceService struct {
	backend Backend
}

func (s *LearningResourceService) List(
	ctx context.Context,
	params *ListLearningResourcesParams,
) iter.Seq2[LearningResource, error] {
	return newEntityIterator[LearningResource](ctx, s.backend, learningResourceBasePath, params)
}

func (s *LearningResourceService) GetByID(ctx context.Context, id string, params *GetLearningResourceParams) (*LearningResource, error) {
	return get[LearningResource](ctx, s.backend, fmt.Sprintf("%s/%s", learningResourceBasePath, id), params)
}

func (s *LearningResourceService) GetByName(ctx context.Context, fqn string, params *GetLearningResourceByNameParams) (*LearningResource, error) {
	return get[LearningResource](ctx, s.backend, fmt.Sprintf("%s/name/%s", learningResourceBasePath, fqn), params)
}

func (s *LearningResourceService) Create(ctx context.Context, body *CreateLearningResource) (*LearningResource, error) {
	return create[LearningResource](ctx, s.backend, learningResourceBasePath, body)
}

func (s *LearningResourceService) CreateOrUpdate(ctx context.Context, body *CreateLearningResource) (*LearningResource, error) {
	return put[LearningResource](ctx, s.backend, learningResourceBasePath, body)
}

func (s *LearningResourceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*LearningResource, error) {
	return patch[LearningResource](ctx, s.backend, fmt.Sprintf("%s/%s", learningResourceBasePath, id), ops)
}

func (s *LearningResourceService) Delete(ctx context.Context, id string, params *DeleteLearningResourceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", learningResourceBasePath, id), params)
}

func (s *LearningResourceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", learningResourceBasePath, id), nil)
}

func (s *LearningResourceService) GetVersion(ctx context.Context, id string, version string) (*LearningResource, error) {
	return get[LearningResource](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", learningResourceBasePath, id, version), nil)
}

func (s *LearningResourceService) Restore(ctx context.Context, body *RestoreEntity) (*LearningResource, error) {
	return put[LearningResource](ctx, s.backend, fmt.Sprintf("%s/restore", learningResourceBasePath), body)
}
