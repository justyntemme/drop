package dal

import (
	models "gitlab.com/nextwavedevs/drop/shared/models"
)

func (p Profile) GetListingById(ctx context.Context, traceID string, uid string) (models.User, error) {
	var result []bson.M
	var user models.User

	pipeline := make([]bson.M, 0)
	log.Println("GetUserByID: ID: " + uid)

	matchStage := bson.M{
		"$match": bson.M{
			"uid": uid,
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
		return user, ErrNotFound

	}
	rawJson, err := json.Marshal(result[0])
	if err != nil {
		log.Println(err)
	}
	log.Println(string(rawJson))
	json.Unmarshal(rawJson, &user)

	return user, nil // returns a raw JSON String

}
