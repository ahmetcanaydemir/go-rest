package repository

import (
	"context"
	"time"

	"log"

	"github.com/ahmetcanaydemir/go-rest/pkg/db"
	"github.com/ahmetcanaydemir/go-rest/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	recordsCollectionName = "records"
	dateLayout            = "2006-01-02"
)

type RecordsRepository interface {
	GetRecords(request model.MongoRequest) ([]model.Record, error)
}

type recordsRepository struct {
	collection *mongo.Collection
}

// NewRecordsRepository creates new instance of RecordsRepository
func NewRecordsRepository() RecordsRepository {
	collection := db.GetDatabase().DB.Collection(recordsCollectionName)
	return &recordsRepository{
		collection: collection,
	}
}

// GetRecords reads records from MongoDB with aggregating counts array and using filters
func (r recordsRepository) GetRecords(request model.MongoRequest) ([]model.Record, error) {
	var records []model.Record

	// Parse error already handled in service layer
	startDate, _ := time.Parse(dateLayout, *request.StartDate)
	endDate, _ := time.Parse(dateLayout, *request.EndDate)

	// Filter records between given dates.
	matchDatesStage := bson.D{
		{
			Key: "$match", Value: bson.M{
				"createdAt": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			},
		},
	}

	// Add a new virtual field named "totalCount" which is sum of the elements of the "count" array.
	addTotalCountFieldStage := bson.D{
		{
			Key: "$addFields", Value: bson.M{
				"totalCount": bson.M{
					"$sum": "$counts",
				},
			},
		},
	}

	// Re-filter records by "totalCount" with given count range.
	matchCountsStage := bson.D{
		{
			Key: "$match", Value: bson.M{
				"totalCount": bson.M{
					"$gte": *request.MinCount,
					"$lte": *request.MaxCount,
				},
			},
		},
	}

	cursor, err := r.collection.Aggregate(context.Background(), mongo.Pipeline{matchDatesStage, addTotalCountFieldStage, matchCountsStage})

	if err != nil {
		log.Println("error occured at mongodb aggregation ", err)
		return nil, err
	}

	err = cursor.All(context.Background(), &records)
	if err != nil {
		log.Println("error occured at mongodb cursor ", err)
		return nil, err
	}
	return records, nil
}
