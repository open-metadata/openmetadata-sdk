package ometa

import (
	"context"
	"fmt"
	"iter"
)

const personaBasePath = "personas"

type PersonaService struct {
	backend Backend
}

func (s *PersonaService) List(
	ctx context.Context,
	params *ListPersonasParams,
) iter.Seq2[Persona, error] {
	return newEntityIterator[Persona](ctx, s.backend, personaBasePath, params)
}

func (s *PersonaService) GetByID(ctx context.Context, id string, params *GetPersonaByIDParams) (*Persona, error) {
	return get[Persona](ctx, s.backend, fmt.Sprintf("%s/%s", personaBasePath, id), params)
}

func (s *PersonaService) GetByName(ctx context.Context, fqn string, params *GetPersonaByFQNParams) (*Persona, error) {
	return get[Persona](ctx, s.backend, fmt.Sprintf("%s/name/%s", personaBasePath, fqn), params)
}

func (s *PersonaService) Create(ctx context.Context, body *CreatePersona) (*Persona, error) {
	return create[Persona](ctx, s.backend, personaBasePath, body)
}

func (s *PersonaService) CreateOrUpdate(ctx context.Context, body *CreatePersona) (*Persona, error) {
	return put[Persona](ctx, s.backend, personaBasePath, body)
}

func (s *PersonaService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Persona, error) {
	return patch[Persona](ctx, s.backend, fmt.Sprintf("%s/%s", personaBasePath, id), ops)
}

func (s *PersonaService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", personaBasePath, id), nil)
}

func (s *PersonaService) GetVersion(ctx context.Context, id string, version string) (*Persona, error) {
	return get[Persona](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", personaBasePath, id, version), nil)
}

func (s *PersonaService) Restore(ctx context.Context, body *RestoreEntity) (*Persona, error) {
	return put[Persona](ctx, s.backend, fmt.Sprintf("%s/restore", personaBasePath), body)
}
