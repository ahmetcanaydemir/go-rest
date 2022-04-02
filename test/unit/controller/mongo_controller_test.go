package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmetcanaydemir/go-rest/pkg/api/controller"
	"github.com/ahmetcanaydemir/go-rest/pkg/model"
	"github.com/ahmetcanaydemir/go-rest/test/unit/service/mock"
)

func Test_PostMongoController(t *testing.T) {
	var tests = []struct {
		name           string
		bodyjson       string
		expectedStatus int
	}{
		{"successful request", `{"startDate":"2016-01-26","endDate":"2018-02-02","minCount":10,"maxCount":100}`, http.StatusOK},
		{"stardDate missing request", `{"endDate":"2018-02-02","minCount":10,"maxCount":100}`, http.StatusBadRequest},
		{"endDate missing request", `{"startDate":"2016-01-26","minCount":10,maxCount":100}`, http.StatusBadRequest},
		{"minCount missing request", `{"startDate":"2016-01-26","endDate":"2018-02-02","maxCount":100}`, http.StatusBadRequest},
		{"maxCount missing request", `{"startDate":"2016-01-26","endDate":"2018-02-02","minCount":10}`, http.StatusBadRequest},
		{"wrong body request", "{wrongbody}", http.StatusBadRequest},
	}

	var requestBody model.MongoRequest
	err := json.Unmarshal([]byte(`{"startDate":"2016-01-26","endDate":"2018-02-02","minCount":10,"maxCount":100}`), &requestBody)
	if err != nil {
		t.Errorf("PostMongoController error occured during json unmarshall %s", err.Error())
	}

	response := &model.MongoResponse{
		Code:    1,
		Msg:     "Success",
		Records: []model.Record{},
	}
	recordsServiceMock := new(mock.RecordsServiceMock)
	recordsServiceMock.On("GetRecords", requestBody).Return(response)
	mongoController := &controller.MongoController{
		Service: recordsServiceMock,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/in-memory", bytes.NewBuffer([]byte(tt.bodyjson)))
			rr := httptest.NewRecorder()
			mongoController.MongoHandler(rr, req)
			gotStatus := rr.Code
			if gotStatus != tt.expectedStatus {
				t.Errorf("GetMongoController(%s) got %v, want %v", tt.name, gotStatus, tt.expectedStatus)
			}
		})
	}
}

func Test_NotAllowedMongoController(t *testing.T) {
	mongoController := controller.MongoController{}
	req, _ := http.NewRequest("PUT", "/mongo", nil)
	rr := httptest.NewRecorder()
	mongoController.MongoHandler(rr, req)

	gotStatus := rr.Code
	if gotStatus != http.StatusMethodNotAllowed {
		t.Errorf("MongoController(%s) got %v, want %v", "method not allowed request", gotStatus, http.StatusMethodNotAllowed)
	}

}
