package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"gitlab.com/nextwavedevs/drop/database"
	"gitlab.com/nextwavedevs/drop/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var StudiosCollection *mongo.Collection = database.OpenCollection(database.Client, "studios") // get collection "profile" from db() which returns *mongo.Client

func (p Profile) GetAllListingsByCompanyId(ctx context.Context, traceID string, uid string) (models.Studio, error) {
	var result []bson.M
	studioWithListings := new(models.Studio)

	pipeline := make([]bson.M, 0)
	log.Println("GetAllListingsByCompanyId" + uid)

	id, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
	}

	matchStage := bson.M{
		"$match": bson.M{
			"_id": id,
		},
	}

	lookupStage := bson.M{
		"$lookup": bson.M{
			"from":         "listings",
			"localField":   "listingsIds",
			"foreignField": "uid",
			"as":           "listings",
		},
	}

	pipeline = append(pipeline, matchStage, lookupStage)

	StudioCursor, err := StudiosCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to execute aggregation %s", err.Error())
	}
	log.Println("-----------------------------------------------------")
	log.Println(pipeline)
	log.Println("-----------------------------------------------------")

	err = StudioCursor.All(ctx, &result)
	if result == nil {
		return models.Studio{}, ErrNotFound

	}
	rawJson, err := json.Marshal(result[0])
	if err != nil {
		log.Println(err)
	}
	log.Println(string(rawJson))
	json.Unmarshal(rawJson, &studioWithListings)

	return *studioWithListings, nil // returns a raw JSON String

}

func (p Profile) GetAllStudios(ctx context.Context, traceID string) ([]models.User, error) {
	var result []bson.M
	var users []models.User

	pipeline := make([]bson.M, 0)

	// matchStage := bson.M{
	// 	"$sort": bson.M{
	// 		"created": -1,
	// 	},
	// }
	groupStage := bson.M{
		"$match": bson.M{
			"_id": bson.M{"$ne": ""},
		},
	}

	// 	query = [
	//     {
	//         '$sort': {
	//             'created': -1
	//         }
	//     },
	//     {
	//         $group: {
	//             '_id':null,
	//             'max':{'$first':"$$ROOT"}
	//         }
	//     }
	// ]

	pipeline = append(pipeline, groupStage)

	userProfileCursor, err := userCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to execute aggregation %s", err.Error())
	}
	log.Println(pipeline)

	err = userProfileCursor.All(ctx, &result)
	if result == nil {
		return users, ErrNotFound

	}
	rawJson, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(rawJson))
	json.Unmarshal(rawJson, &users)

	return users, nil // returns a raw JSON String

}
