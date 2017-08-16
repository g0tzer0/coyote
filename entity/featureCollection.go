package entity

import (
	"encoding/json"
	"fmt"
)

// Coordinates : Reprensent the x,y coordinates of the geometry
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// UnmarshalJSON : Deserialize JSON Coordinates
func (c *Coordinates) UnmarshalJSON(b []byte) error {
	coordinates := [2]json.Number{}

	err := json.Unmarshal(b, &coordinates)
	if err != nil {
		return err
	}

	c.Longitude, err = coordinates[0].Float64()
	if err != nil {
		return err
	}
	c.Latitude, err = coordinates[1].Float64()
	if err != nil {
		return err
	}

	return err
}

// MarshalJSON : Serialize JSON Coordinates
func (c Coordinates) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"[%f, %f]"`, c.Latitude, c.Longitude)), nil
}

// Geometry : Represent the Geometry details of each feature
type Geometry struct {
	Type        string
	Coordinates Coordinates `json:"coordinates"`
}

// Properties : Represent the Properties of each feature
type Properties struct {
	Name       string `json:"name"`
	Population int    `json:"population"`
	CartoDBId  int    `json:"cartodb_id"`
}

// FeatureEntry : Represent one feature from the collection
type FeatureEntry struct {
	ID         int
	Type       string
	Geometry   Geometry
	Properties Properties
}

// FeatureCollection : Represent the collection of features
type FeatureCollection struct {
	Type     string
	Features []FeatureEntry
}
