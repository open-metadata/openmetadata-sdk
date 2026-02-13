package ometa

import (
	"context"
	"fmt"
	"iter"
)

const classificationBasePath = "classifications"

type ClassificationService struct {
	backend Backend
}

func (s *ClassificationService) List(
	ctx context.Context,
	params *ListClassificationsParams,
) iter.Seq2[Classification, error] {
	return newEntityIterator[Classification](ctx, s.backend, classificationBasePath, params)
}

func (s *ClassificationService) GetByID(ctx context.Context, id string, params *GetClassificationByIDParams) (*Classification, error) {
	return get[Classification](ctx, s.backend, fmt.Sprintf("%s/%s", classificationBasePath, id), params)
}

func (s *ClassificationService) GetByName(ctx context.Context, fqn string, params *GetClassificationByNameParams) (*Classification, error) {
	return get[Classification](ctx, s.backend, fmt.Sprintf("%s/name/%s", classificationBasePath, fqn), params)
}

func (s *ClassificationService) Create(ctx context.Context, body *CreateClassification) (*Classification, error) {
	return create[Classification](ctx, s.backend, classificationBasePath, body)
}

func (s *ClassificationService) CreateOrUpdate(ctx context.Context, body *CreateClassification) (*Classification, error) {
	return put[Classification](ctx, s.backend, classificationBasePath, body)
}

func (s *ClassificationService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Classification, error) {
	return patch[Classification](ctx, s.backend, fmt.Sprintf("%s/%s", classificationBasePath, id), ops)
}

func (s *ClassificationService) Delete(ctx context.Context, id string, params *DeleteClassificationParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", classificationBasePath, id), params)
}

func (s *ClassificationService) DeleteByName(ctx context.Context, fqn string, params *DeleteClassificationByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", classificationBasePath, fqn), params)
}

func (s *ClassificationService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", classificationBasePath, id), nil)
}

func (s *ClassificationService) GetVersion(ctx context.Context, id string, version string) (*Classification, error) {
	return get[Classification](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", classificationBasePath, id, version), nil)
}

func (s *ClassificationService) Restore(ctx context.Context, body *RestoreEntity) (*Classification, error) {
	return put[Classification](ctx, s.backend, fmt.Sprintf("%s/restore", classificationBasePath), body)
}
