package dal

// struct for storing data
type User struct {
	ID   string `bson:"_id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}
