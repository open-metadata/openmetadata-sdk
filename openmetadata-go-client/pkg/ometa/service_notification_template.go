package ometa

import (
	"context"
	"fmt"
	"iter"
)

const notificationTemplateBasePath = "notificationTemplates"

type NotificationTemplateService struct {
	backend Backend
}

func (s *NotificationTemplateService) List(
	ctx context.Context,
	params *ListNotificationTemplatesParams,
) iter.Seq2[NotificationTemplate, error] {
	return newEntityIterator[NotificationTemplate](ctx, s.backend, notificationTemplateBasePath, params)
}

func (s *NotificationTemplateService) GetByID(ctx context.Context, id string, params *GetNotificationTemplateByIdParams) (*NotificationTemplate, error) {
	return get[NotificationTemplate](ctx, s.backend, fmt.Sprintf("%s/%s", notificationTemplateBasePath, id), params)
}

func (s *NotificationTemplateService) GetByName(ctx context.Context, fqn string, params *GetNotificationTemplateByFQNParams) (*NotificationTemplate, error) {
	return get[NotificationTemplate](ctx, s.backend, fmt.Sprintf("%s/name/%s", notificationTemplateBasePath, fqn), params)
}

func (s *NotificationTemplateService) Create(ctx context.Context, body *CreateNotificationTemplate) (*NotificationTemplate, error) {
	return create[NotificationTemplate](ctx, s.backend, notificationTemplateBasePath, body)
}

func (s *NotificationTemplateService) CreateOrUpdate(ctx context.Context, body *CreateNotificationTemplate) (*NotificationTemplate, error) {
	return put[NotificationTemplate](ctx, s.backend, notificationTemplateBasePath, body)
}

func (s *NotificationTemplateService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*NotificationTemplate, error) {
	return patch[NotificationTemplate](ctx, s.backend, fmt.Sprintf("%s/%s", notificationTemplateBasePath, id), ops)
}

func (s *NotificationTemplateService) Delete(ctx context.Context, id string, params *DeleteNotificationTemplateParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", notificationTemplateBasePath, id), params)
}

func (s *NotificationTemplateService) DeleteByName(ctx context.Context, fqn string, params *DeleteNotificationTemplateByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", notificationTemplateBasePath, fqn), params)
}

func (s *NotificationTemplateService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", notificationTemplateBasePath, id), nil)
}

func (s *NotificationTemplateService) GetVersion(ctx context.Context, id string, version string) (*NotificationTemplate, error) {
	return get[NotificationTemplate](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", notificationTemplateBasePath, id, version), nil)
}

func (s *NotificationTemplateService) Restore(ctx context.Context, body *RestoreEntity) (*NotificationTemplate, error) {
	return put[NotificationTemplate](ctx, s.backend, fmt.Sprintf("%s/restore", notificationTemplateBasePath), body)
}
