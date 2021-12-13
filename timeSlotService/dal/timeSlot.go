package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"gitlab.com/nextwavedevs/drop/protodrop"
	"gitlab.com/nextwavedevs/drop/timeslotservice/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "timeSlots") // get collection "timeSlots" from db() which returns *mongo.Client

func GetTimeSlotById(ctx context.Context, traceID string, uid string) (*protodrop.TimeSlot, error) {
	var result []bson.M
	var timeSlot protodrop.TimeSlot

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
		return &timeSlot, err

	}
	rawJson, err := json.Marshal(result[0])
	if err != nil {
		log.Println(err)
	}
	log.Println(string(rawJson))
	json.Unmarshal(rawJson, &timeSlot)

	return &timeSlot, nil // returns a raw JSON String

}
