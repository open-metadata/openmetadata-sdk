package ometa

import (
	"context"
	"fmt"
	"iter"
)

const documentBasePath = "docStore"

type DocumentService struct {
	backend Backend
}

func (s *DocumentService) List(
	ctx context.Context,
	params *ListDocumentsParams,
) iter.Seq2[Document, error] {
	return newEntityIterator[Document](ctx, s.backend, documentBasePath, params)
}

func (s *DocumentService) GetByID(ctx context.Context, id string, params *Get2Params) (*Document, error) {
	return get[Document](ctx, s.backend, fmt.Sprintf("%s/%s", documentBasePath, id), params)
}

func (s *DocumentService) GetByName(ctx context.Context, fqn string, params *GetDocumentByFQNParams) (*Document, error) {
	return get[Document](ctx, s.backend, fmt.Sprintf("%s/name/%s", documentBasePath, fqn), params)
}

func (s *DocumentService) Create(ctx context.Context, body *CreateDocument) (*Document, error) {
	return create[Document](ctx, s.backend, documentBasePath, body)
}

func (s *DocumentService) CreateOrUpdate(ctx context.Context, body *CreateDocument) (*Document, error) {
	return put[Document](ctx, s.backend, documentBasePath, body)
}

func (s *DocumentService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*Document, error) {
	return patch[Document](ctx, s.backend, fmt.Sprintf("%s/%s", documentBasePath, id), ops)
}

func (s *DocumentService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", documentBasePath, id), nil)
}

func (s *DocumentService) GetVersion(ctx context.Context, id string, version string) (*Document, error) {
	return get[Document](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", documentBasePath, id, version), nil)
}

func (s *DocumentService) Restore(ctx context.Context, body *RestoreEntity) (*Document, error) {
	return put[Document](ctx, s.backend, fmt.Sprintf("%s/restore", documentBasePath), body)
}
