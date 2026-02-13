package ometa

import (
	"context"
	"fmt"
	"iter"
)

const workflowDefinitionBasePath = "governance/workflowDefinitions"

type WorkflowDefinitionService struct {
	backend Backend
}

func (s *WorkflowDefinitionService) List(
	ctx context.Context,
	params *ListWorkflowDefinitionsParams,
) iter.Seq2[WorkflowDefinition, error] {
	return newEntityIterator[WorkflowDefinition](ctx, s.backend, workflowDefinitionBasePath, params)
}

func (s *WorkflowDefinitionService) GetByID(ctx context.Context, id string, params *GetWorkflowDefinitionByIDParams) (*WorkflowDefinition, error) {
	return get[WorkflowDefinition](ctx, s.backend, fmt.Sprintf("%s/%s", workflowDefinitionBasePath, id), params)
}

func (s *WorkflowDefinitionService) GetByName(ctx context.Context, fqn string, params *GetWorkflowDefinitionByFQNParams) (*WorkflowDefinition, error) {
	return get[WorkflowDefinition](ctx, s.backend, fmt.Sprintf("%s/name/%s", workflowDefinitionBasePath, fqn), params)
}

func (s *WorkflowDefinitionService) Create(ctx context.Context, body *CreateWorkflowDefinition) (*WorkflowDefinition, error) {
	return create[WorkflowDefinition](ctx, s.backend, workflowDefinitionBasePath, body)
}

func (s *WorkflowDefinitionService) CreateOrUpdate(ctx context.Context, body *CreateWorkflowDefinition) (*WorkflowDefinition, error) {
	return put[WorkflowDefinition](ctx, s.backend, workflowDefinitionBasePath, body)
}

func (s *WorkflowDefinitionService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*WorkflowDefinition, error) {
	return patch[WorkflowDefinition](ctx, s.backend, fmt.Sprintf("%s/%s", workflowDefinitionBasePath, id), ops)
}

func (s *WorkflowDefinitionService) Delete(ctx context.Context, id string, params *DeleteWorkflowDefinitionParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", workflowDefinitionBasePath, id), params)
}

func (s *WorkflowDefinitionService) DeleteByName(ctx context.Context, fqn string, params *DeleteWorkflowDefinitionByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", workflowDefinitionBasePath, fqn), params)
}

func (s *WorkflowDefinitionService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", workflowDefinitionBasePath, id), nil)
}

func (s *WorkflowDefinitionService) GetVersion(ctx context.Context, id string, version string) (*WorkflowDefinition, error) {
	return get[WorkflowDefinition](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", workflowDefinitionBasePath, id, version), nil)
}

func (s *WorkflowDefinitionService) Restore(ctx context.Context, body *RestoreEntity) (*WorkflowDefinition, error) {
	return put[WorkflowDefinition](ctx, s.backend, fmt.Sprintf("%s/restore", workflowDefinitionBasePath), body)
}
