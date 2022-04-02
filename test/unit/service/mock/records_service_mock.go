package mock

import (
	"github.com/ahmetcanaydemir/go-rest/pkg/api/service"
	"github.com/ahmetcanaydemir/go-rest/pkg/model"
	"github.com/stretchr/testify/mock"
)

type RecordsServiceMock struct {
	mock.Mock
}

var _ service.RecordsService = (*RecordsServiceMock)(nil)

func (m *RecordsServiceMock) GetRecords(request model.MongoRequest) *model.MongoResponse {
	args := m.Called(request)
	return args.Get(0).(*model.MongoResponse)
}
