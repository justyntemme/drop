package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"gitlab.com/nextwavedevs/drop/database"
	models "gitlab.com/nextwavedevs/drop/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "listings") // get collection "listings" from db() which returns *mongo.Client

func GetListingById(ctx context.Context, traceID string, uid string) (models.Listing, error) {
	var result []bson.M
	var user models.Listing

	pipeline := make([]bson.M, 0)
	log.Println("GetUserByID: ID: " + uid)

	matchStage := bson.M{
		"$match": bson.M{
			"id": uid,
		},
	}

	pipeline = append(pipeline, matchStage)

	userProfileCursor, err := userCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to execute aggregation %s", err.Error())
	}
	log.Println(pipeline)

	err = userProfileCursor.All(ctx, &result)
	if result == nil {
		return user, err

	}
	rawJson, err := json.Marshal(result[0])
	if err != nil {
		log.Println(err)
	}
	log.Println(string(rawJson))
	json.Unmarshal(rawJson, &user)

	return user, nil // returns a raw JSON String

}
