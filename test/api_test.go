package test

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var (
	client             http.Client
	expectedStatusCode int
	expectedBody       string
	url                string
)

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = http.Client{Transport: tr}
}

func TestGetRoot(t *testing.T) {
	url = "https://localhost:8443/"
	expectedStatusCode = 400
	expectedBody = `{"error": "Unknown request"}`

	resp := getRequest(t)

	validate(t, resp)
}

func TestGetIdRoot(t *testing.T) {
	url = "https://localhost:8443/id/"
	expectedStatusCode = 400
	expectedBody = `{"error": "Unknown request"}`

	resp := getRequest(t)

	validate(t, resp)
}

func TestGetFeaturesByInvalidID(t *testing.T) {
	url = "https://localhost:8443/id/abc"
	expectedStatusCode = 400
	expectedBody = `{"error": "Unknown request"}`

	resp := getRequest(t)

	validate(t, resp)
}

func TestGetFeaturesByValidIDWithoutContent(t *testing.T) {
	url = "https://localhost:8443/id/321654"
	expectedStatusCode = 204
	expectedBody = `{}`

	resp := getRequest(t)

	validate(t, resp)
}

func TestGetFeaturesByValidIDWithContent(t *testing.T) {
	url = "https://localhost:8443/id/744"
	expectedStatusCode = 200
	expectedBody = `{"cartodb_id":744,"name":"Oriel","population":2500,"coordinates":[43.069946,-80.643498]}`

	resp := getRequest(t)

	validate(t, resp)
}

func TestGetFeaturesByIDAndInvalidDist(t *testing.T) {
	url = "https://localhost:8443/id/744?dist=a"
	expectedStatusCode = 400
	expectedBody = `{"error": "Unknown request"}`

	resp := getRequest(t)

	validate(t, resp)
}

func TestGetFeaturesByIDAndValidDistWithoutContent(t *testing.T) {
	url = "https://localhost:8443/id/321456?dist=5"
	expectedStatusCode = 204
	expectedBody = ``

	resp := getRequest(t)

	validate(t, resp)
}

func TestGetFeaturesByIDAndValidDistWithContent(t *testing.T) {
	url = "https://localhost:8443/id/744?dist=5"
	expectedStatusCode = 200
	expectedBody = `[{"cartodb_id":737,"name":"Beaconsfield","population":2500,"coordinates":[43.064133,-80.598719]},{"cartodb_id":776,"name":"Oxford Centre","population":109,"coordinates":[43.098307,-80.682714]},{"cartodb_id":778,"name":"Vandecar","population":2500,"coordinates":[43.099452,-80.62027]}]`

	resp := getRequest(t)

	validate(t, resp)
}

func getRequest(t *testing.T) (resp *http.Response) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	return resp
}

func validate(t *testing.T, resp *http.Response) {
	fmt.Print("Header:\n")

	for k, v := range resp.Header {
		fmt.Printf("  %s: %s\n", k, v)
	}

	fmt.Printf("Status Code: \n  %d\n", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Body: \n  %s\n", body)

	if resp.StatusCode != expectedStatusCode {
		t.Fail()
	}

	if strings.EqualFold(expectedBody, ``) {
		if string(body) != expectedBody {
			t.Fail()
		}
	}
}
