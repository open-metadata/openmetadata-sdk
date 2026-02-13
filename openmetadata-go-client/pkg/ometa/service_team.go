package ometa

import (
	"context"
	"fmt"
	"iter"
)

const teamBasePath = "teams"

type TeamService struct {
	backend Backend
}

func (s *TeamService) List(
	ctx context.Context,
	params *ListTeamsParams,
) iter.Seq2[Team, error] {
	return newEntityIterator[Team](ctx, s.backend, teamBasePath, params)
}

func (s *TeamService) GetByID(ctx context.Context, id string, params *GetTeamByIDParams) (*Team, error) {
	return get[Team](ctx, s.backend, fmt.Sprintf("%s/%s", teamBasePath, id), params)
}

func (s *TeamService) GetByName(ctx context.Context, fqn string, params *GetTeamByFQNParams) (*Team, error) {
	return get[Team](ctx, s.backend, fmt.Sprintf("%s/name/%s", teamBasePath, fqn), params)
}

func (s *TeamService) Create(ctx context.Context, body *CreateTeam) (*Team, error) {
	return create[Team](ctx, s.backend, teamBasePath, body)
}

func (s *TeamService) CreateOrUpdate(ctx context.Context, body *CreateTeam) (*Team, error) {
	return put[Team](ctx, s.backend, teamBasePath, body)
}

func (s *TeamService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Team, error) {
	return patch[Team](ctx, s.backend, fmt.Sprintf("%s/%s", teamBasePath, id), ops)
}

func (s *TeamService) Delete(ctx context.Context, id string, params *DeleteTeamParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", teamBasePath, id), params)
}

func (s *TeamService) DeleteByName(ctx context.Context, fqn string, params *DeleteTeamByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", teamBasePath, fqn), params)
}

func (s *TeamService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", teamBasePath, id), nil)
}

func (s *TeamService) GetVersion(ctx context.Context, id string, version string) (*Team, error) {
	return get[Team](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", teamBasePath, id, version), nil)
}

func (s *TeamService) Restore(ctx context.Context, body *RestoreEntity) (*Team, error) {
	return put[Team](ctx, s.backend, fmt.Sprintf("%s/restore", teamBasePath), body)
}
