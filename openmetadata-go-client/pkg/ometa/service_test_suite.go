package ometa

import (
	"context"
	"fmt"
	"iter"
)

const testSuiteBasePath = "dataQuality/testSuites"

type TestSuiteService struct {
	backend Backend
}

func (s *TestSuiteService) List(
	ctx context.Context,
	params *ListTestSuitesParams,
) iter.Seq2[TestSuite, error] {
	return newEntityIterator[TestSuite](ctx, s.backend, testSuiteBasePath, params)
}

func (s *TestSuiteService) GetByID(ctx context.Context, id string, params *Get5Params) (*TestSuite, error) {
	return get[TestSuite](ctx, s.backend, fmt.Sprintf("%s/%s", testSuiteBasePath, id), params)
}

func (s *TestSuiteService) GetByName(ctx context.Context, fqn string, params *GetTestSuiteByNameParams) (*TestSuite, error) {
	return get[TestSuite](ctx, s.backend, fmt.Sprintf("%s/name/%s", testSuiteBasePath, fqn), params)
}

func (s *TestSuiteService) Create(ctx context.Context, body *CreateTestSuite) (*TestSuite, error) {
	return create[TestSuite](ctx, s.backend, testSuiteBasePath, body)
}

func (s *TestSuiteService) CreateOrUpdate(ctx context.Context, body *CreateTestSuite) (*TestSuite, error) {
	return put[TestSuite](ctx, s.backend, testSuiteBasePath, body)
}

func (s *TestSuiteService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*TestSuite, error) {
	return patch[TestSuite](ctx, s.backend, fmt.Sprintf("%s/%s", testSuiteBasePath, id), ops)
}

func (s *TestSuiteService) Delete(ctx context.Context, id string, params *DeleteTestSuiteParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", testSuiteBasePath, id), params)
}

func (s *TestSuiteService) DeleteByName(ctx context.Context, fqn string, params *DeleteTestSuiteByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", testSuiteBasePath, fqn), params)
}

func (s *TestSuiteService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", testSuiteBasePath, id), nil)
}

func (s *TestSuiteService) GetVersion(ctx context.Context, id string, version string) (*TestSuite, error) {
	return get[TestSuite](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", testSuiteBasePath, id, version), nil)
}

func (s *TestSuiteService) Restore(ctx context.Context, body *RestoreEntity) (*TestSuite, error) {
	return put[TestSuite](ctx, s.backend, fmt.Sprintf("%s/restore", testSuiteBasePath), body)
}
