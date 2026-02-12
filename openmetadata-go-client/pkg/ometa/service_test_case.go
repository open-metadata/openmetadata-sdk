package ometa

import (
	"context"
	"fmt"
	"iter"
)

const testCaseBasePath = "dataQuality/testCases"

type TestCaseService struct {
	backend Backend
}

func (s *TestCaseService) List(
	ctx context.Context,
	params *ListTestCasesParams,
) iter.Seq2[TestCase, error] {
	return newEntityIterator[TestCase](ctx, s.backend, testCaseBasePath, params)
}

func (s *TestCaseService) GetByID(ctx context.Context, id string, params *Get3Params) (*TestCase, error) {
	return get[TestCase](ctx, s.backend, fmt.Sprintf("%s/%s", testCaseBasePath, id), params)
}

func (s *TestCaseService) GetByName(ctx context.Context, fqn string, params *GetTestCaseByNameParams) (*TestCase, error) {
	return get[TestCase](ctx, s.backend, fmt.Sprintf("%s/name/%s", testCaseBasePath, fqn), params)
}

func (s *TestCaseService) Create(ctx context.Context, body *CreateTestCase) (*TestCase, error) {
	return create[TestCase](ctx, s.backend, testCaseBasePath, body)
}

func (s *TestCaseService) CreateOrUpdate(ctx context.Context, body *CreateTestCase) (*TestCase, error) {
	return put[TestCase](ctx, s.backend, testCaseBasePath, body)
}

func (s *TestCaseService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*TestCase, error) {
	return patch[TestCase](ctx, s.backend, fmt.Sprintf("%s/%s", testCaseBasePath, id), ops)
}

func (s *TestCaseService) Delete(ctx context.Context, id string, params *DeleteTestCaseParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", testCaseBasePath, id), params)
}

func (s *TestCaseService) DeleteByName(ctx context.Context, fqn string, params *DeleteTestCaseByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", testCaseBasePath, fqn), params)
}

func (s *TestCaseService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", testCaseBasePath, id), nil)
}

func (s *TestCaseService) GetVersion(ctx context.Context, id string, version string) (*TestCase, error) {
	return get[TestCase](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", testCaseBasePath, id, version), nil)
}

func (s *TestCaseService) Restore(ctx context.Context, body *RestoreEntity) (*TestCase, error) {
	return put[TestCase](ctx, s.backend, fmt.Sprintf("%s/restore", testCaseBasePath), body)
}
