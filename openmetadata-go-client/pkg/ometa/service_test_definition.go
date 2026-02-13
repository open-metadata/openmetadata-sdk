package ometa

import (
	"context"
	"fmt"
	"iter"
)

const testDefinitionBasePath = "dataQuality/testDefinitions"

type TestDefinitionService struct {
	backend Backend
}

func (s *TestDefinitionService) List(
	ctx context.Context,
	params *ListTestDefinitionsParams,
) iter.Seq2[TestDefinition, error] {
	return newEntityIterator[TestDefinition](ctx, s.backend, testDefinitionBasePath, params)
}

func (s *TestDefinitionService) GetByID(ctx context.Context, id string, params *Get4Params) (*TestDefinition, error) {
	return get[TestDefinition](ctx, s.backend, fmt.Sprintf("%s/%s", testDefinitionBasePath, id), params)
}

func (s *TestDefinitionService) GetByName(ctx context.Context, fqn string, params *GetTestDefinitionByNameParams) (*TestDefinition, error) {
	return get[TestDefinition](ctx, s.backend, fmt.Sprintf("%s/name/%s", testDefinitionBasePath, fqn), params)
}

func (s *TestDefinitionService) Create(ctx context.Context, body *CreateTestDefinition) (*TestDefinition, error) {
	return create[TestDefinition](ctx, s.backend, testDefinitionBasePath, body)
}

func (s *TestDefinitionService) CreateOrUpdate(ctx context.Context, body *CreateTestDefinition) (*TestDefinition, error) {
	return put[TestDefinition](ctx, s.backend, testDefinitionBasePath, body)
}

func (s *TestDefinitionService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*TestDefinition, error) {
	return patch[TestDefinition](ctx, s.backend, fmt.Sprintf("%s/%s", testDefinitionBasePath, id), ops)
}

func (s *TestDefinitionService) Delete(ctx context.Context, id string, params *DeleteTestDefinitionParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", testDefinitionBasePath, id), params)
}

func (s *TestDefinitionService) DeleteByName(ctx context.Context, fqn string, params *DeleteTestDefinitionByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", testDefinitionBasePath, fqn), params)
}

func (s *TestDefinitionService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", testDefinitionBasePath, id), nil)
}

func (s *TestDefinitionService) GetVersion(ctx context.Context, id string, version string) (*TestDefinition, error) {
	return get[TestDefinition](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", testDefinitionBasePath, id, version), nil)
}

func (s *TestDefinitionService) Restore(ctx context.Context, body *RestoreEntity) (*TestDefinition, error) {
	return put[TestDefinition](ctx, s.backend, fmt.Sprintf("%s/restore", testDefinitionBasePath), body)
}
