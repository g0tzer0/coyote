package model

import (
	"github.com/g0tzer0/coyote/entity"
	"github.com/g0tzer0/coyote/web/data"
)

var (
	featureRepository = data.NewFeatureRepository()
)

// GetFeaturesByID : Get feature definition by cartodbId
func GetFeaturesByID(cartodbID int) ([]entity.FeatureEntry, error) {
	return featureRepository.GetByID(cartodbID)
}

// GetFeaturesByIDAndDist : Get feature definition by cartodbId and distance
func GetFeaturesByIDAndDist(cartodbID int, dist int) ([]entity.FeatureEntry, error) {
	return featureRepository.GetByIDAndDist(cartodbID, dist)
}
