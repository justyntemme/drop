package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"gitlab.com/nextwavedevs/drop/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p Profile) GetAllListingsByCompanyId(ctx context.Context, traceID string, uid string) ([]models.Listing, error) {
	var result []bson.M
	var listings []models.Listing

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
			"from":         "studios",
			"localField":   "listingsIds",
			"foreignField": "uid",
			"as":           "$$ROOT",
		},
	}

	replaceRootStage := bson.M{
		"newRoot": "$listings",
	}

	pipeline = append(pipeline, matchStage, lookupStage, replaceRootStage)

	ListingsCursor, err := userCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to execute aggregation %s", err.Error())
	}
	log.Println(pipeline)

	err = ListingsCursor.All(ctx, &result)
	if result == nil {
		return listings, ErrNotFound

	}
	rawJson, err := json.Marshal(result[0])
	if err != nil {
		log.Println(err)
	}
	log.Println(string(rawJson))
	json.Unmarshal(rawJson, &listings)

	return listings, nil // returns a raw JSON String

}
