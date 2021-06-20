package dal

// struct for storing data
type User struct {
	Name string `json:name`
	Age  int    `json:age`
	City string `json:city`
}

type userProfileResult struct {
	User  []User `bson:"user"`
	Total int    `bson:"total"`
}
