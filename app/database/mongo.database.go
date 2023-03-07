package database

import (
	"context"
	"fmt"
	"go-blog/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	DataBase *mongo.Database
	dbConn   *mongo.Client
}

func (db *MongoDB) Conn(ctx context.Context) {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	// Connect to MongoDB
	dbConfig := options.Client().ApplyURI(config.DBUri)
	dbConn, err := mongo.Connect(ctx, dbConfig)
	if err != nil {
		panic(err)
	}
	if err := dbConn.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB successfully connected...")

	db.DataBase = dbConn.Database("go-blog")
}

func (db *MongoDB) DisConn(ctx context.Context) {
	db.dbConn.Disconnect(ctx)
}
