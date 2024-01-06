package database

import (
	"context"
	"os"
	"time"

	"nyx/logger"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Id    	string
	Name 	string
	Balance int
}

var client *mongo.Client
var GoDB *mongo.Database
var UsersCollection *mongo.Collection

func Connect() {
	logger.Logger.InfoF("Connecting to database.")

	godotenv.Load()

	DB_TOKEN := os.Getenv("DB_TOKEN")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(DB_TOKEN).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Logger.FatalF("[ERROR]: %v", err.Error())
	}

	GoDB = client.Database("go")
	UsersCollection = GoDB.Collection("users")

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Logger.FatalF("[ERROR]: %v", err.Error())
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		logger.Logger.FatalF("[ERROR]: %v", err.Error())
	}
	logger.Logger.InfoF("Connected to database.")
	
	logger.Logger.InfoF("Databases connected: %s", databases)
}

func Disconnect() {
	logger.Logger.InfoF("Disconnecting from database.")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	logger.Logger.InfoF("Disconnected from database.")
	client.Disconnect(ctx)
}
