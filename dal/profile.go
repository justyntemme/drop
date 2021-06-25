package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gitlab.com/nextwavedevs/drop/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		db: db,
	}
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "profile") // get collection "profile" from db() which returns *mongo.Client

// Create Profile or Signup

func (p Profile) CreateProfile(ctx context.Context, traceID string, u User) (User, error) {

	//validating the user input 
	if err := Check(u); err != nil {
		return User{}, errors.Wrap(err, "validating data")
	}

	//parsing the user input into the User model
	person := User{
		Name: u.Name,
		City: u.City,
		Age: u.Age,
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

// Get Profile of a particular User by Name

func GetUserById(ctx context.Context, w http.ResponseWriter, uid string) string {

	w.Header().Set("Content-Type", "application/json")
	var user []bson.M

	//err := userCollection.FindOne(context.TODO(), bson.D{{"uid", uid}}).Decode(&result)

	pipeline := make([]bson.M, 0)

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
	if err != nil {
		fmt.Println(err)
	}

	userProfileCursor.All(ctx, &user)
	rawJson, err := json.Marshal(user)
	fmt.Print(rawJson)

	return string(rawJson) // returns a JSON String

}

//Update Profile of User

func updateProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Name string `json:"name"` //value that has to be matched
		City string `json:"city"` // value that has to be modified
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"name", body.Name}} // converting value to BSON type
	after := options.After                // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"city", body.City}}}}
	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete Profile of User

func deleteProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                   //slice for multiple documents
	cur, err := userCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}
