package models

type Studio struct {
	UID      string    `json:"uid"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	City     string    `json:"city"`
	Listings []Listing `json:"listings"`
}

type Listing struct {
	UID        string `json:"uid"`
	Name       string `json:"name"`
	Descrition int    `json:"description"`
	City       string `json:"city"`
}
