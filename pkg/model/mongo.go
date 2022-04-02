package model

import (
	"time"
)

type MongoRequest struct {
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
	MinCount  *int    `json:"minCount"`
	MaxCount  *int    `json:"maxCount"`
}

type MongoResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []Record `json:"records"`
}

type Record struct {
	Key        string    `bson:"key" json:"key"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	TotalCount int       `bson:"totalCount" json:"totalCount"`
}
