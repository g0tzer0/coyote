package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/g0tzer0/coyote/entity"
	"github.com/g0tzer0/coyote/web/model"
)

// FeatureController : controller handling features
type FeatureController struct{}

func mapFeatureToEntry(feature entity.FeatureEntry) (result entity.ResultEntry) {
	var coordinates = [2]float64{
		feature.Geometry.Coordinates.Latitude,
		feature.Geometry.Coordinates.Longitude}

	return entity.ResultEntry{
		Name:        feature.Properties.Name,
		CartoDBId:   feature.Properties.CartoDBId,
		Population:  feature.Properties.Population,
		Coordinates: coordinates}
}

// GetFeaturesByID : return the feature cartodbId, name, population and coordinates in JSON format
func (f *FeatureController) GetFeaturesByID(w http.ResponseWriter, r *http.Request) {
	matches := featuresByIDPath.FindStringSubmatch(r.URL.Path)
	cartodbID, _ := strconv.Atoi(matches[1])

	w.Header().Add("Location", "/id/"+strconv.Itoa(cartodbID))

	features, err := model.GetFeaturesByID(cartodbID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Unknown request"}`))
		return
	}

	if len(features) == 0 {
		w.Header().Add("Location", "/id/"+strconv.Itoa(cartodbID))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response, err := json.Marshal(mapFeatureToEntry(features[0]))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Unknown request"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

// GetFeaturesByIDAndDist : return all features based on the cartodbID and within a scare of the distance parameter
func (f *FeatureController) GetFeaturesByIDAndDist(w http.ResponseWriter, r *http.Request) {
	matches := featuresByIDPath.FindStringSubmatch(r.URL.Path)
	cartodbID, _ := strconv.Atoi(matches[1])

	w.Header().Add("Location", "/id/"+strconv.Itoa(cartodbID)+"?"+r.URL.RawQuery)

	dist, err := strconv.Atoi(r.URL.Query().Get("dist"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Unknown request"}`))
		return
	}

	features, err := model.GetFeaturesByIDAndDist(cartodbID, dist)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Unknown request"}`))
		return
	}

	if len(features) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var results []entity.ResultEntry

	for _, feature := range features {
		results = append(results, mapFeatureToEntry(feature))
	}

	response, err := json.Marshal(results)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Unknown request"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
