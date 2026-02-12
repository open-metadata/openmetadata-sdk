package ometa

import (
	"context"
	"fmt"
	"iter"
)

const tagBasePath = "tags"

type TagService struct {
	backend Backend
}

func (s *TagService) List(
	ctx context.Context,
	params *ListTagsParams,
) iter.Seq2[Tag, error] {
	return newEntityIterator[Tag](ctx, s.backend, tagBasePath, params)
}

func (s *TagService) GetByID(ctx context.Context, id string, params *GetTagByIDParams) (*Tag, error) {
	return get[Tag](ctx, s.backend, fmt.Sprintf("%s/%s", tagBasePath, id), params)
}

func (s *TagService) GetByName(ctx context.Context, fqn string, params *GetTagByFQNParams) (*Tag, error) {
	return get[Tag](ctx, s.backend, fmt.Sprintf("%s/name/%s", tagBasePath, fqn), params)
}

func (s *TagService) Create(ctx context.Context, body *CreateTag) (*Tag, error) {
	return create[Tag](ctx, s.backend, tagBasePath, body)
}

func (s *TagService) CreateOrUpdate(ctx context.Context, body *CreateTag) (*Tag, error) {
	return put[Tag](ctx, s.backend, tagBasePath, body)
}

func (s *TagService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Tag, error) {
	return patch[Tag](ctx, s.backend, fmt.Sprintf("%s/%s", tagBasePath, id), ops)
}

func (s *TagService) Delete(ctx context.Context, id string, params *DeleteTagParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", tagBasePath, id), params)
}

func (s *TagService) DeleteByName(ctx context.Context, fqn string, params *DeleteTagByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", tagBasePath, fqn), params)
}

func (s *TagService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", tagBasePath, id), nil)
}

func (s *TagService) GetVersion(ctx context.Context, id string, version string) (*Tag, error) {
	return get[Tag](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", tagBasePath, id, version), nil)
}

func (s *TagService) Restore(ctx context.Context, body *RestoreEntity) (*Tag, error) {
	return put[Tag](ctx, s.backend, fmt.Sprintf("%s/restore", tagBasePath), body)
}
