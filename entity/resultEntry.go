package entity

// ResultEntry : Represent one feature from the collection
type ResultEntry struct {
	CartoDBId   int        `json:"cartodb_id"`
	Name        string     `json:"name"`
	Population  int        `json:"population"`
	Coordinates [2]float64 `json:"coordinates"`
}
