package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AhmadHussainRandhawa/mongoapi/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://officialahmadrandhawa_db_user:roughpassword@cluster0.z6slbvm.mongodb.net/?appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection // represent and provide access to one MongoDB collection (table).

// connect with mongoDb

func init() {
	// Client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Create client
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

// MongoDB helpers

// insert one record
func InsertOne(movie model.Netflix) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertOneResult, err := collection.InsertOne(ctx, movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id:", insertOneResult.InsertedID)

}
