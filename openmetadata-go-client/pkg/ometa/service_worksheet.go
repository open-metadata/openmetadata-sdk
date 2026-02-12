package ometa

import (
	"context"
	"fmt"
	"iter"
)

const worksheetBasePath = "drives/worksheets"

type WorksheetService struct {
	backend Backend
}

func (s *WorksheetService) List(
	ctx context.Context,
	params *ListWorksheetsParams,
) iter.Seq2[Worksheet, error] {
	return newEntityIterator[Worksheet](ctx, s.backend, worksheetBasePath, params)
}

func (s *WorksheetService) GetByID(ctx context.Context, id string, params *GetWorksheetByIDParams) (*Worksheet, error) {
	return get[Worksheet](ctx, s.backend, fmt.Sprintf("%s/%s", worksheetBasePath, id), params)
}

func (s *WorksheetService) GetByName(ctx context.Context, fqn string, params *GetWorksheetByFQNParams) (*Worksheet, error) {
	return get[Worksheet](ctx, s.backend, fmt.Sprintf("%s/name/%s", worksheetBasePath, fqn), params)
}

func (s *WorksheetService) Create(ctx context.Context, body *CreateWorksheet) (*Worksheet, error) {
	return create[Worksheet](ctx, s.backend, worksheetBasePath, body)
}

func (s *WorksheetService) CreateOrUpdate(ctx context.Context, body *CreateWorksheet) (*Worksheet, error) {
	return put[Worksheet](ctx, s.backend, worksheetBasePath, body)
}

func (s *WorksheetService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Worksheet, error) {
	return patch[Worksheet](ctx, s.backend, fmt.Sprintf("%s/%s", worksheetBasePath, id), ops)
}

func (s *WorksheetService) Delete(ctx context.Context, id string, params *DeleteWorksheetParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", worksheetBasePath, id), params)
}

func (s *WorksheetService) DeleteByName(ctx context.Context, fqn string, params *DeleteWorksheetByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", worksheetBasePath, fqn), params)
}

func (s *WorksheetService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", worksheetBasePath, id), nil)
}

func (s *WorksheetService) GetVersion(ctx context.Context, id string, version string) (*Worksheet, error) {
	return get[Worksheet](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", worksheetBasePath, id, version), nil)
}

func (s *WorksheetService) Restore(ctx context.Context, body *RestoreEntity) (*Worksheet, error) {
	return put[Worksheet](ctx, s.backend, fmt.Sprintf("%s/restore", worksheetBasePath), body)
}
