package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetCategory(t *testing.T) {

	req, err := http.NewRequest("GET", "/category/533", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	if rr.Body.String() == "" {
		t.Error("Handler returned empty body")
	}

	// Decode category from the returned JSON response
	dec := json.NewDecoder(strings.NewReader(rr.Body.String()))
	var cat category
	err = dec.Decode(&cat)

	// Check if the decoding failed
	if err != nil {
		t.Error("Category decoding from JSON failed!")
	}

	// Check the decoded category ID
	expectedId := 533
	if cat.ID != expectedId {
		t.Errorf("Category ID expected value: %d, received: %d", expectedId, cat.ID)
	}

	// Check the decoded category name
	expectedName := "iPhone"
	if cat.Name != expectedName {
		t.Errorf("Category name expected value: %v, received: %v", expectedName, cat.Name)
	}
}
