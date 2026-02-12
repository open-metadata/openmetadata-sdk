package ometa

import (
	"context"
	"fmt"
	"iter"
)

const topicBasePath = "topics"

type TopicService struct {
	backend Backend
}

func (s *TopicService) List(
	ctx context.Context,
	params *ListTopicsParams,
) iter.Seq2[Topic, error] {
	return newEntityIterator[Topic](ctx, s.backend, topicBasePath, params)
}

func (s *TopicService) GetByID(ctx context.Context, id string, params *Get7Params) (*Topic, error) {
	return get[Topic](ctx, s.backend, fmt.Sprintf("%s/%s", topicBasePath, id), params)
}

func (s *TopicService) GetByName(ctx context.Context, fqn string, params *GetTopicByFQNParams) (*Topic, error) {
	return get[Topic](ctx, s.backend, fmt.Sprintf("%s/name/%s", topicBasePath, fqn), params)
}

func (s *TopicService) Create(ctx context.Context, body *CreateTopic) (*Topic, error) {
	return create[Topic](ctx, s.backend, topicBasePath, body)
}

func (s *TopicService) CreateOrUpdate(ctx context.Context, body *CreateTopic) (*Topic, error) {
	return put[Topic](ctx, s.backend, topicBasePath, body)
}

func (s *TopicService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Topic, error) {
	return patch[Topic](ctx, s.backend, fmt.Sprintf("%s/%s", topicBasePath, id), ops)
}

func (s *TopicService) Delete(ctx context.Context, id string, params *DeleteTopicParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", topicBasePath, id), params)
}

func (s *TopicService) DeleteByName(ctx context.Context, fqn string, params *DeleteTopicByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", topicBasePath, fqn), params)
}

func (s *TopicService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", topicBasePath, id), nil)
}

func (s *TopicService) GetVersion(ctx context.Context, id string, version string) (*Topic, error) {
	return get[Topic](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", topicBasePath, id, version), nil)
}

func (s *TopicService) Restore(ctx context.Context, body *RestoreEntity) (*Topic, error) {
	return put[Topic](ctx, s.backend, fmt.Sprintf("%s/restore", topicBasePath), body)
}
