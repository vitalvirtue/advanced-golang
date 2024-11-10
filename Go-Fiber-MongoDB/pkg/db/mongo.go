package db

import (
	"context"
	"log"
	"time"
	"github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var HRMSDatabase *mongo.Database

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetMongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	HRMSDatabase = client.Database("go-fiber-hrms")
}