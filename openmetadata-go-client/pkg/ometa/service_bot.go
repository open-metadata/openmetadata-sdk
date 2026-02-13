package ometa

import (
	"context"
	"fmt"
	"iter"
)

const botBasePath = "bots"

type BotService struct {
	backend Backend
}

func (s *BotService) List(
	ctx context.Context,
	params *ListBotsParams,
) iter.Seq2[Bot, error] {
	return newEntityIterator[Bot](ctx, s.backend, botBasePath, params)
}

func (s *BotService) GetByID(ctx context.Context, id string, params *GetBotByIDParams) (*Bot, error) {
	return get[Bot](ctx, s.backend, fmt.Sprintf("%s/%s", botBasePath, id), params)
}

func (s *BotService) GetByName(ctx context.Context, fqn string, params *GetBotByFQNParams) (*Bot, error) {
	return get[Bot](ctx, s.backend, fmt.Sprintf("%s/name/%s", botBasePath, fqn), params)
}

func (s *BotService) Create(ctx context.Context, body *CreateBot) (*Bot, error) {
	return create[Bot](ctx, s.backend, botBasePath, body)
}

func (s *BotService) CreateOrUpdate(ctx context.Context, body *CreateBot) (*Bot, error) {
	return put[Bot](ctx, s.backend, botBasePath, body)
}

func (s *BotService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Bot, error) {
	return patch[Bot](ctx, s.backend, fmt.Sprintf("%s/%s", botBasePath, id), ops)
}

func (s *BotService) Delete(ctx context.Context, id string, params *DeleteBotParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", botBasePath, id), params)
}

func (s *BotService) DeleteByName(ctx context.Context, fqn string, params *DeleteBotByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", botBasePath, fqn), params)
}

func (s *BotService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", botBasePath, id), nil)
}

func (s *BotService) GetVersion(ctx context.Context, id string, version string) (*Bot, error) {
	return get[Bot](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", botBasePath, id, version), nil)
}

func (s *BotService) Restore(ctx context.Context, body *RestoreEntity) (*Bot, error) {
	return put[Bot](ctx, s.backend, fmt.Sprintf("%s/restore", botBasePath), body)
}
