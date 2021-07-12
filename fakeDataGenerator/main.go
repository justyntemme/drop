package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	fake "github.com/brianvoe/gofakeit/v6"
	"gitlab.com/nextwavedevs/drop/database"
	"gitlab.com/nextwavedevs/drop/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	database.DB()
	createData()

}
func createData() {
	ctx, _ := context.WithTimeout(context.Background(), 50000)

	var listingsCollection *mongo.Collection = database.OpenCollection(database.Client, "profile") // get collection "profile" from db() which returns *mongo.Client
	var studioCollection *mongo.Collection = database.OpenCollection(database.Client, "profile")   // get collection "profile" from db() which returns *mongo.Client

	var s models.Studio

	fake.Struct(&s)

	//Create address for studio
	fakeAddress := fake.Address()
	s.Address.Street = fakeAddress.Street
	s.Address.State = fakeAddress.State
	s.Address.Zip = fakeAddress.Zip
	s.Address.Country = fakeAddress.Country
	s.Address.Coordinates = append(s.Address.Coordinates, fakeAddress.Longitude)
	s.Address.Coordinates = append(s.Address.Coordinates, fakeAddress.Latitude)

	var listings []models.Listing
	i := 0
	for i < 4 {
		l := new(models.Listing)
		l.Title = fake.Sentence(3)
		l.Description = fake.Sentence(5)
		l.UID = fake.UUID()

		s.ListingIds = append(s.ListingIds, l.UID)

		listings = append(listings, *l)
		listingsCollection.InsertOne(nil, l)

		i++

	}

	listingsFile, _ := json.MarshalIndent(listings, "", " ")
	studioFile, _ := json.MarshalIndent(s, "", " ")
	_ = ioutil.WriteFile("listings.json", listingsFile, 0644)
	_ = ioutil.WriteFile("studio.json", studioFile, 0644)

	studioCollection.InsertOne(ctx, s)

	fmt.Println(s)
}
