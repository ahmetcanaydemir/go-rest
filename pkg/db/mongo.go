package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ahmetcanaydemir/go-rest/pkg/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const (
	mongoConnTimeout = 10 * time.Second
)

type Database struct {
	DB     *mongo.Database
	Client *mongo.Client
}

var database *Database
var once sync.Once

func GetDatabase() *Database {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), mongoConnTimeout)
		defer cancel()

		cs, err := connstring.ParseAndValidate(configs.Server.Config.DbConnectionString)
		if err != nil {
			panic(fmt.Sprintf("could not extract table name from mongo dsn : %s", err))
		}

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.Server.Config.DbConnectionString))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to mongo: %s", err))
		}

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			panic(fmt.Sprintf("could not ping mongodb after connect: %s", err))
		}

		mongoDB := client.Database(cs.Database)

		database = &Database{
			DB:     mongoDB,
			Client: client,
		}
	})

	return database
}
