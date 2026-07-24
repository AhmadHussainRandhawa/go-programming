package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AhmadHussainRandhawa/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // always use context like in that way.
	defer cancel()                                                          // In the following i use less in this way, but this is ideal

	insertOneResult, err := collection.InsertOne(ctx, movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id:", insertOneResult.InsertedID)

}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.ModifiedCount)

}

// Delete one record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, _ := collection.DeleteOne(context.Background(), filter)

	fmt.Println(result.DeletedCount)
}

// Delete all records

func deleteAllMovie() {

	result, err := collection.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.DeletedCount)

}

// get all movies from database - also i will show the ideal way of using context

func getAllMovies() []bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	cur, _ := collection.Find(ctx, bson.M{})

	var movies []bson.M

	for cur.Next(ctx) {
		var movie bson.M
		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)

	}
	return movies

}

// Actual controller - file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	movies := getAllMovies()
	json.NewEncoder(w).Encode(movies)
}
