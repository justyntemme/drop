package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/nextwavedevs/drop/database"
	"gitlab.com/nextwavedevs/drop/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*todos*/
/*work on the profile methods to handle logs and traces GET, GETby ID, UPDATE, DELETE*/
/*Handle tracing*/

// User manages the set of API's for user access.
type Profile struct {
	log *log.Logger
	db  *mongo.Client
}

// New constructs a User for api access.
func New(log *log.Logger, db *mongo.Client) Profile {
	return Profile{
		log: log,
		db:  db,
	}
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "profile") // get collection "profile" from db() which returns *mongo.Client

// Create Profile or Signup

func (p Profile) CreateProfile(ctx context.Context, traceID string, u models.User) (models.User, error) {

	//validating the user input
	if err := Check(u); err != nil {
		return models.User{}, errors.Wrap(err, "validating data")
	}

	//parsing the user input into the User model
	person := models.User{
		UID:  GenerateID(),
		Name: u.Name,
		City: u.City,
		Age:  u.Age,
	}

	//inserting into mongo db
	insertResult, err := userCollection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult, "InsertID:", insertResult.InsertedID)
	p.log.Printf("%s: %s", traceID, "profile.Create")
	return person, nil
}

// // Get Profile of a particular User by id
// func (p Profile) GetUserById(ctx context.Context, traceID string, uid string) (User, error) {

// 	//Validate if the uid entered is in correct mode
// 	if err := CheckID(uid); err != nil {
// 		return User{}, ErrInvalidID
// 	}

// 	var result User //  an unordered representation of a BSON document which is a Map

// 	err := userCollection.FindOne(ctx, bson.D{{Key: "_id", Value: uid}}).Decode(&result)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	log.Printf("%s: %s", traceID, "profile.GetUserById")

// 	return result, nil // returns a raw JSON String
// }
func (p Profile) GetUserById(ctx context.Context, traceID string, uid string) (models.User, error) {
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

	userProfileCursor.All(ctx, &result)
	if result != nil {
		rawJson, err := json.Marshal(result[0])
		if err != nil {
			log.Println(err)
		}
		log.Println(string(rawJson))
		json.Unmarshal(rawJson, &user)

	}

	return user, nil // returns a raw JSON String

}

//Update Profile of User

func (p Profile) UpdateProfile(ctx context.Context, traceID string, uid string, u models.User) error {

	//Validate uid
	if err := CheckID(uid); err != nil {
		return ErrInvalidID
	}

	//Validate coming user details
	if err := Check(u); err != nil {
		return errors.Wrap(err, "validating data")
	}

	//Get the current user details
	usr, err := p.GetUserById(ctx, traceID, uid)
	if err != nil {
		return errors.Wrap(err, "updating user")
	}

	log.Println("Formar user details: ", usr)

	filter := bson.D{{Key: "name", Value: u.Name}} // converting value to BSON type

	after := options.After // for returning updated document

	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}

	update := bson.M{
		"$set": bson.M{
			"name": u.Name,
			"city": u.City,
			"age":  u.Age,
		},
	}
	updateResult := userCollection.FindOneAndUpdate(ctx, filter, update, &returnOpt)

	var result models.User
	_ = updateResult.Decode(&result)
	log.Printf("%s: %s", traceID, "user.Update")

	return nil
}

//Delete Profile of User

func (p Profile) DeleteProfile(ctx context.Context, traceID string, uid string) error {

	if err := CheckID(uid); err != nil {
		return ErrInvalidID
	}

	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase

	res, err := userCollection.DeleteOne(ctx, bson.D{{Key: "_id", Value: uid}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	log.Printf("%s: %s", traceID, "user.Delete") // return number of documents deleted

	return nil
}

func (p Profile) GetAllUsers(ctx context.Context, traceID string, pageNumber int, rowsPerPage int) ([]*models.User, error) {

	data := struct {
		Offset      int `db:"offset"`
		RowsPerPage int `db:"rows_per_page"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}

	// Pass these options to the Find method
	findOptionsOffset := options.Find()
	findOptionPage := options.Find()
	findOptionsOffset.SetLimit(int64(data.Offset))
	findOptionPage.SetLimit(int64(data.RowsPerPage))

	var results []*models.User //slice for multiple documents

	cur, err := userCollection.Find(ctx, bson.D{{}}, findOptionsOffset, findOptionPage) //returns a *mongo.Cursor
	if err != nil {
		fmt.Println(err)
	}
	for cur.Next(ctx) { //Next() gets the next document for corresponding cursor

		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem) // appending document pointed by Next()
	}
	cur.Close(ctx) // close the cursor once stream of documents has exhausted
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	log.Printf("%s: %s", traceID, "user.Query")

	return results, nil
}
