package ometa

import (
	"context"
	"fmt"
	"iter"
)

const messagingServiceBasePath = "services/messagingServices"

type MessagingServiceService struct {
	backend Backend
}

func (s *MessagingServiceService) List(
	ctx context.Context,
	params *ListMessagingServiceParams,
) iter.Seq2[MessagingService, error] {
	return newEntityIterator[MessagingService](ctx, s.backend, messagingServiceBasePath, params)
}

func (s *MessagingServiceService) GetByID(ctx context.Context, id string, params *GetMessagingServiceByIDParams) (*MessagingService, error) {
	return get[MessagingService](ctx, s.backend, fmt.Sprintf("%s/%s", messagingServiceBasePath, id), params)
}

func (s *MessagingServiceService) GetByName(ctx context.Context, fqn string, params *GetMessagingServiceByFQNParams) (*MessagingService, error) {
	return get[MessagingService](ctx, s.backend, fmt.Sprintf("%s/name/%s", messagingServiceBasePath, fqn), params)
}

func (s *MessagingServiceService) Create(ctx context.Context, body *CreateMessagingService) (*MessagingService, error) {
	return create[MessagingService](ctx, s.backend, messagingServiceBasePath, body)
}

func (s *MessagingServiceService) CreateOrUpdate(ctx context.Context, body *CreateMessagingService) (*MessagingService, error) {
	return put[MessagingService](ctx, s.backend, messagingServiceBasePath, body)
}

func (s *MessagingServiceService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*MessagingService, error) {
	return patch[MessagingService](ctx, s.backend, fmt.Sprintf("%s/%s", messagingServiceBasePath, id), ops)
}

func (s *MessagingServiceService) Delete(ctx context.Context, id string, params *DeleteMessagingServiceParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", messagingServiceBasePath, id), params)
}

func (s *MessagingServiceService) DeleteByName(ctx context.Context, fqn string, params *DeleteMessagingServiceByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", messagingServiceBasePath, fqn), params)
}

func (s *MessagingServiceService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", messagingServiceBasePath, id), nil)
}

func (s *MessagingServiceService) GetVersion(ctx context.Context, id string, version string) (*MessagingService, error) {
	return get[MessagingService](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", messagingServiceBasePath, id, version), nil)
}

func (s *MessagingServiceService) Restore(ctx context.Context, body *RestoreEntity) (*MessagingService, error) {
	return put[MessagingService](ctx, s.backend, fmt.Sprintf("%s/restore", messagingServiceBasePath), body)
}
