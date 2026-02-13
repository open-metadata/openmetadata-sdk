package ometa

import (
	"context"
	"fmt"
	"iter"
)

const dataContractBasePath = "dataContracts"

type DataContractService struct {
	backend Backend
}

func (s *DataContractService) List(
	ctx context.Context,
	params *ListDataContractsParams,
) iter.Seq2[DataContract, error] {
	return newEntityIterator[DataContract](ctx, s.backend, dataContractBasePath, params)
}

func (s *DataContractService) GetByID(ctx context.Context, id string, params *GetDataContractByIDParams) (*DataContract, error) {
	return get[DataContract](ctx, s.backend, fmt.Sprintf("%s/%s", dataContractBasePath, id), params)
}

func (s *DataContractService) GetByName(ctx context.Context, fqn string, params *GetDataContractByFQNParams) (*DataContract, error) {
	return get[DataContract](ctx, s.backend, fmt.Sprintf("%s/name/%s", dataContractBasePath, fqn), params)
}

func (s *DataContractService) Create(ctx context.Context, body *CreateDataContract) (*DataContract, error) {
	return create[DataContract](ctx, s.backend, dataContractBasePath, body)
}

func (s *DataContractService) CreateOrUpdate(ctx context.Context, body *CreateDataContract) (*DataContract, error) {
	return put[DataContract](ctx, s.backend, dataContractBasePath, body)
}

func (s *DataContractService) Patch(ctx context.Context, id string, ops []JSONPatchOp) (*DataContract, error) {
	return patch[DataContract](ctx, s.backend, fmt.Sprintf("%s/%s", dataContractBasePath, id), ops)
}

func (s *DataContractService) Delete(ctx context.Context, id string, params *DeleteDataContractParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/%s", dataContractBasePath, id), params)
}

func (s *DataContractService) DeleteByName(ctx context.Context, fqn string, params *DeleteDataContractByNameParams) error {
	return del(ctx, s.backend, fmt.Sprintf("%s/name/%s", dataContractBasePath, fqn), params)
}

func (s *DataContractService) ListVersions(ctx context.Context, id string) (*EntityHistory, error) {
	return get[EntityHistory](ctx, s.backend, fmt.Sprintf("%s/%s/versions", dataContractBasePath, id), nil)
}

func (s *DataContractService) GetVersion(ctx context.Context, id string, version string) (*DataContract, error) {
	return get[DataContract](ctx, s.backend, fmt.Sprintf("%s/%s/versions/%s", dataContractBasePath, id, version), nil)
}

func (s *DataContractService) Restore(ctx context.Context, body *RestoreEntity) (*DataContract, error) {
	return put[DataContract](ctx, s.backend, fmt.Sprintf("%s/restore", dataContractBasePath), body)
}
