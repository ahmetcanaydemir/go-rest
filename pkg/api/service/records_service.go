package service

import (
	"log"
	"time"

	"github.com/ahmetcanaydemir/go-rest/pkg/api/repository"
	"github.com/ahmetcanaydemir/go-rest/pkg/model"
)

type RecordsService interface {
	GetRecords(request model.MongoRequest) *model.MongoResponse
}

type recordsService struct {
	repository repository.RecordsRepository
}

const dateLayout = "2006-01-02"

// NewRecordsService creates new instance of RecordsService
func NewRecordsService(repositories ...repository.RecordsRepository) RecordsService {
	var repo repository.RecordsRepository
	if len(repositories) == 0 {
		repo = repository.NewRecordsRepository()
	} else {
		repo = repositories[0]
	}

	return &recordsService{
		repository: repo,
	}
}

// GetRecords verifies request body fields and returns repository's response
func (s recordsService) GetRecords(request model.MongoRequest) *model.MongoResponse {
	startDate, err := time.Parse(dateLayout, *request.StartDate)
	if err != nil {
		log.Println("error occured while parsing the startDate", err)
		return s.errorResponse(MessageCodes.DATE_PARSE_ERROR)
	}

	endDate, err := time.Parse(dateLayout, *request.EndDate)
	if err != nil {
		log.Println("error occured while parsing the endDate", err)
		return s.errorResponse(MessageCodes.DATE_PARSE_ERROR)
	}

	if endDate.Before(startDate) {
		return s.errorResponse(MessageCodes.WRONG_DATE_ORDER)
	}
	if *request.MaxCount < *request.MinCount {
		return s.errorResponse(MessageCodes.WRONG_COUNT_ORDER)
	}
	if *request.MaxCount < 0 || *request.MinCount < 0 {
		return s.errorResponse(MessageCodes.COUNT_IS_NEGATIVE)
	}

	records, err := s.repository.GetRecords(request)
	if err != nil {
		return s.errorResponse(MessageCodes.SERVER_ERROR)
	}

	return &model.MongoResponse{
		Code:    MessageCodes.SUCCESS,
		Msg:     messages[MessageCodes.SUCCESS],
		Records: records,
	}
}

func (s recordsService) errorResponse(messageCode int) *model.MongoResponse {
	return &model.MongoResponse{
		Code: messageCode,
		Msg:  messages[messageCode],
	}
}

//#
var messages map[int]string = map[int]string{
	0:  "Success",
	1:  "End date must after than start date",
	2:  "MaxCount must greater than MinCount",
	3:  "MaxCount and MinCount must be positive numbers",
	88: "Cannot parse date, please verify the date you entered",
	99: "Unhandled server error occured",
}

var MessageCodes = struct {
	SUCCESS           int
	WRONG_DATE_ORDER  int
	WRONG_COUNT_ORDER int
	COUNT_IS_NEGATIVE int
	DATE_PARSE_ERROR  int
	SERVER_ERROR      int
}{
	SUCCESS:           0,
	WRONG_DATE_ORDER:  1,
	WRONG_COUNT_ORDER: 2,
	COUNT_IS_NEGATIVE: 3,
	DATE_PARSE_ERROR:  88,
	SERVER_ERROR:      99,
}
