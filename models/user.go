package models

type User struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}
