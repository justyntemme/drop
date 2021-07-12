package models

type Studio struct {
	UID        string    `json:"uid" fake:"skip"`
	Name       string    `json:"name" fake:"{company}"`
	Address    Address   `json:"address" fake:"skip"`
	ListingIds []string  `json:"listingsIds" fake:"skip"`
	Listings   []Listing `json:"listings" fake:"skip"`
	About      string    `json:"about" fake:"{sentence:10}"`
}
