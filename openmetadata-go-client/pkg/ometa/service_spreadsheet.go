package ometa

import (
	"context"
	"fmt"
	"iter"
)

const spreadsheetBasePath = "drives/spreadsheets"

type SpreadsheetService struct {
	backend Backend
}

func (s *SpreadsheetService) List(
	ctx context.Context,
	params *ListSpreadsheetsParams,
) iter.Seq2[Spreadsheet, error] {
	return newEntityIterator[Spreadsheet](ctx, s.backend, spreadsheetBasePath, params)
}

func (s *SpreadsheetService) GetByID(ctx context.Context, id string, params *GetSpreadsheetByIDParams) (*Spreadsheet, error) {
	return get[Spreadsheet](ctx, s.backend, fmt.Sprintf("%s/%s", spreadsheetBasePath, id), params)
}

func (s *SpreadsheetService) GetByName(ctx context.Context, fqn string, params *GetSpreadsheetByFQNParams) (*Spreadsheet, error) {
	return get[Spreadsheet](ctx, s.backend, fmt.Sprintf("%s/name/%s", spreadsheetBasePath, fqn), params)
}

func (s *SpreadsheetService) Create(ctx context.Context, body *CreateSpreadsheet) (*Spreadsheet, error) {
	return create[Spreadsheet](ctx, s.backend, spreadsheetBasePath, body)
}

func (s *SpreadsheetService) CreateOrUpdate(ctx context.Context, body *CreateSpreadsheet) (*Spreadsheet, error) {
	return put[Spreadsheet](ctx, s.backend, spreadsheetBasePath, body)
}

func (s *SpreadsheetService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Spreadsheet, error) {
	return patch[Spreadsheet](ctx, s.backend, fmt.Sprintf("%s/%s", spreadsheetBasePath, id), ops)
}

func (s *SpreadsheetService) Delete(ctx context.Context, id string, params *DeleteSpreadsheetParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", spreadsheetBasePath, id), params)
}

func (s *SpreadsheetService) DeleteByName(ctx context.Context, fqn string, params *DeleteSpreadsheetByFQNParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", spreadsheetBasePath, fqn), params)
}

func (s *SpreadsheetService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", spreadsheetBasePath, id), nil)
}

func (s *SpreadsheetService) GetVersion(ctx context.Context, id string, version string) (*Spreadsheet, error) {
	return get[Spreadsheet](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", spreadsheetBasePath, id, version), nil)
}

func (s *SpreadsheetService) Restore(ctx context.Context, body *RestoreEntity) (*Spreadsheet, error) {
	return put[Spreadsheet](ctx, s.backend, fmt.Sprintf("%s/restore", spreadsheetBasePath), body)
}
