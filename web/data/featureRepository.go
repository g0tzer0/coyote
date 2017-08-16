package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/g0tzer0/coyote/entity"
	"github.com/g0tzer0/coyote/util"
)

// FeatureRepository : repository of features
type FeatureRepository interface {
	GetByID(cartodbID int) ([]entity.FeatureEntry, error)
	GetByIDAndDist(cartodbID int, dist int) ([]entity.FeatureEntry, error)
}

// NewFeatureRepository : get the repository instance
func NewFeatureRepository() FeatureRepository {
	return &featureRepository{}
}

type featureRepository struct{}

var featureCollection *entity.FeatureCollection

// Read geojson file only once and return the featureCollection
func readGeoFile() (*entity.FeatureCollection, error) {
	if featureCollection == nil {
		file, err := os.Open("../data/cities.geojson")
		if err != nil {
			fmt.Println("error:", err)
			return nil, err
		}

		dec := json.NewDecoder(file)

		err = dec.Decode(&featureCollection)
		if err != nil {
			fmt.Println("error:", err)
			return nil, err
		}

		defer file.Close()
	}

	return featureCollection, nil
}

// GetByID : Get feature by cartodbID
func (f *featureRepository) GetByID(cartodbID int) ([]entity.FeatureEntry, error) {
	var featureEntries []entity.FeatureEntry

	featureCollection, err := readGeoFile()
	if err != nil {
		return nil, err
	}

	for _, feature := range featureCollection.Features {
		if feature.Properties.CartoDBId == cartodbID {
			featureEntries = append(featureEntries, feature)
		}
	}

	return featureEntries, nil
}

// GetByIDAndDist : Get feature by cartodbId and distance
func (f *featureRepository) GetByIDAndDist(cartodbID int, dist int) ([]entity.FeatureEntry, error) {

	var featureEntries []entity.FeatureEntry
	var featureCentral entity.FeatureEntry

	featureCollection, err := readGeoFile()
	if err != nil {
		return nil, err
	}

	for _, feature := range featureCollection.Features {
		if feature.Properties.CartoDBId == cartodbID {
			featureCentral = feature
			break
		}
	}

	var boundings = util.BoundingCoordinates(
		featureCentral.Geometry.Coordinates.Latitude,
		featureCentral.Geometry.Coordinates.Longitude,
		float64(dist))

	var minLat = boundings[0]
	var minLong = boundings[1]
	var maxLat = boundings[2]
	var maxLong = boundings[3]

	for _, feature := range featureCollection.Features {
		if feature.Geometry.Coordinates.Longitude <= maxLong &&
			feature.Geometry.Coordinates.Longitude >= minLong &&
			feature.Geometry.Coordinates.Latitude <= maxLat &&
			feature.Geometry.Coordinates.Latitude >= minLat &&
			feature.Properties.CartoDBId != featureCentral.Properties.CartoDBId {
			featureEntries = append(featureEntries, feature)
		}
	}

	return featureEntries, nil
}
