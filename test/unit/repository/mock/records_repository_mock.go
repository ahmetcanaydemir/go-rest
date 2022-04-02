package mock

import (
	"github.com/ahmetcanaydemir/go-rest/pkg/api/repository"
	"github.com/ahmetcanaydemir/go-rest/pkg/model"
	"github.com/stretchr/testify/mock"
)

type RecordsRepositoryMock struct {
	mock.Mock
}

var _ repository.RecordsRepository = (*RecordsRepositoryMock)(nil)

func (m *RecordsRepositoryMock) GetRecords(request model.MongoRequest) ([]model.Record, error) {
	args := m.Called(request)

	if args.Get(1) == nil {
		return args.Get(0).([]model.Record), nil
	} else {
		return args.Get(0).([]model.Record), args.Get(1).(error)
	}
}
