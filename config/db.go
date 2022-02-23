package config

import (
	"context"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseInitConnection() (*mongo.Database, context.CancelFunc, error) {
	dbConfig := InitConfig("dbConn")
	dbName := InitConfig("dbname")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		dbConfig).SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database(dbName)
	fmt.Println("Database connection success!")
	return db, cancel, nil
}
