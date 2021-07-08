package models

type Address struct {
	Type        string    `json:"type" fake:"Point"`
	Coordinates []float64 `json:"coordinates" faker"skip"`
	Street      string    `json:"street" xml:"street" fake:"{address.street}`
	City        string    `json:"city" xml:"city" fake:"{address.city}`
	State       string    `json:"state" xml:"state" fake:"{address.state}`
	Zip         string    `json:"zip" xml:"zip" fake:"{address.zip}`
	Country     string    `json:"country" xml:"country" fake:"{address.country}`
}
