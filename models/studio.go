package models

type Studio struct {
	UID      string   `json:"uid" fake:"skip"`
	Name     string   `json:"name" fake:"{company}"`
	Address  Address  `json:"address" fake:"skip"`
	Listings []string `json:"listingsIds" fake:"skip"`
	About    string   `json:"about" fake:"{sentence:10}"`
}
