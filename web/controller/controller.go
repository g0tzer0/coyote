package controller

import (
	"net/http"
	"regexp"
)

var (
	featureController = new(FeatureController)

	featuresByIDPath = regexp.MustCompile(`^/id/(\d+)`)
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	dist := r.URL.Query().Get("dist")

	switch r.Method {
	case http.MethodGet:
		if featuresByIDPath.MatchString(r.URL.Path) && dist != "" {
			featureController.GetFeaturesByIDAndDist(w, r)
		} else if featuresByIDPath.MatchString(r.URL.Path) {
			featureController.GetFeaturesByID(w, r)
		} else {
			w.Header().Add("Location", "/")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Unknown request"}`))
		}
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Unknown request"}`))
	}
}

// Setup : setup controllers
func Setup() {
	http.HandleFunc("/", handleRequest)
}
