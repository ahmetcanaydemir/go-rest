package service_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/ahmetcanaydemir/go-rest/pkg/api/service"
	"github.com/ahmetcanaydemir/go-rest/pkg/model"
	"github.com/ahmetcanaydemir/go-rest/test/unit/repository/mock"
)

func Test_RecordsService(t *testing.T) {
	var tests = []struct {
		name                string
		mongoRequestJson    string
		expectedMessageCode int
	}{
		{"successful request", `{"startDate":"2016-01-26","endDate":"2018-02-02","minCount":10,"maxCount":100}`, service.MessageCodes.SUCCESS},
		{"stardDate cannot parse", `{"startDate":"3301-26","endDate":"2018-02-02","minCount":10,"maxCount":100}`, service.MessageCodes.DATE_PARSE_ERROR},
		{"endDate cannot parse", `{"startDate":"2016-01-26","endDate":"3301-26","minCount":10,"maxCount":100}`, service.MessageCodes.DATE_PARSE_ERROR},
		{"endDate less than startDate", `{"startDate":"2016-01-26","endDate":"2014-02-02","minCount":10,"maxCount":100}`, service.MessageCodes.WRONG_DATE_ORDER},
		{"maxCount less than minCount", `{"startDate":"2016-01-26","endDate":"2018-02-02","minCount":150,"maxCount":100}`, service.MessageCodes.WRONG_COUNT_ORDER},
		{"count is negative", `{"startDate":"2016-01-26","endDate":"2018-02-02","minCount":-120,"maxCount":-100}`, service.MessageCodes.COUNT_IS_NEGATIVE},
		{"mongoDB error occured", `{"startDate":"2016-01-26","endDate":"2018-02-02","minCount":10,"maxCount":100}`, service.MessageCodes.SERVER_ERROR},
	}

	var request model.MongoRequest
	err := json.Unmarshal([]byte(tests[0].mongoRequestJson), &request)
	if err != nil {
		t.Errorf("RecordsService error occured during json unmarshall %s", err.Error())
	}

	response := []model.Record{
		{Key: "test-1", CreatedAt: time.Now(), TotalCount: 10},
		{Key: "test-2", CreatedAt: time.Now(), TotalCount: 20},
		{Key: "test-3", CreatedAt: time.Now(), TotalCount: 30},
	}

	recordsRepositoryMock := new(mock.RecordsRepositoryMock)
	recordsRepositoryMock.On("GetRecords", request).Return(response, nil)
	recordsService := service.NewRecordsService(recordsRepositoryMock)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.expectedMessageCode == service.MessageCodes.SERVER_ERROR {
				recordsRepositoryMock = new(mock.RecordsRepositoryMock)
				recordsRepositoryMock.On("GetRecords", request).Return(response, errors.New("test mongo error"))
				recordsService = service.NewRecordsService(recordsRepositoryMock)
			}

			var mongoRequest model.MongoRequest
			err := json.Unmarshal([]byte(tt.mongoRequestJson), &mongoRequest)
			if err != nil {
				t.Errorf("GetMongoController(%s) error occured during json unmarshall %s", tt.name, err.Error())
			}

			gotResponse := recordsService.GetRecords(mongoRequest)
			if gotResponse.Code != tt.expectedMessageCode {
				t.Errorf("GetMongoController(%s) got %v, want %v", tt.name, gotResponse.Code, tt.expectedMessageCode)
			}
		})
	}
}
