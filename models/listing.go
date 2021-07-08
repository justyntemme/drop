package models

type Listing struct {
	UID         string  `json:"uid"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Address     Address `json:"address"`
}
