package ometa

import (
	"context"
	"fmt"
	"iter"
)

const eventSubscriptionBasePath = "events/subscriptions"

type EventSubscriptionService struct {
	backend Backend
}

func (s *EventSubscriptionService) List(
	ctx context.Context,
	params *ListEventSubscriptionsParams,
) iter.Seq2[EventSubscription, error] {
	return newEntityIterator[EventSubscription](ctx, s.backend, eventSubscriptionBasePath, params)
}

func (s *EventSubscriptionService) GetByID(ctx context.Context, id string, params *GetEventSubscriptionByIDParams) (*EventSubscription, error) {
	return get[EventSubscription](ctx, s.backend, fmt.Sprintf("%s/%s", eventSubscriptionBasePath, id), params)
}

func (s *EventSubscriptionService) GetByName(ctx context.Context, fqn string, params *GetEventSubscriptionByNameParams) (*EventSubscription, error) {
	return get[EventSubscription](ctx, s.backend, fmt.Sprintf("%s/name/%s", eventSubscriptionBasePath, fqn), params)
}

func (s *EventSubscriptionService) Create(ctx context.Context, body *CreateEventSubscription) (*EventSubscription, error) {
	return create[EventSubscription](ctx, s.backend, eventSubscriptionBasePath, body)
}

func (s *EventSubscriptionService) CreateOrUpdate(ctx context.Context, body *CreateEventSubscription) (*EventSubscription, error) {
	return put[EventSubscription](ctx, s.backend, eventSubscriptionBasePath, body)
}

func (s *EventSubscriptionService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*EventSubscription, error) {
	return patch[EventSubscription](ctx, s.backend, fmt.Sprintf("%s/%s", eventSubscriptionBasePath, id), ops)
}

func (s *EventSubscriptionService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", eventSubscriptionBasePath, id), nil)
}

func (s *EventSubscriptionService) GetVersion(ctx context.Context, id string, version string) (*EventSubscription, error) {
	return get[EventSubscription](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", eventSubscriptionBasePath, id, version), nil)
}

func (s *EventSubscriptionService) Restore(ctx context.Context, body *RestoreEntity) (*EventSubscription, error) {
	return put[EventSubscription](ctx, s.backend, fmt.Sprintf("%s/restore", eventSubscriptionBasePath), body)
}
