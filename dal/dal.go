package dal

// struct for storing data
type User struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}
