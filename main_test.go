package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/branohricardo/workshop-category-microservice/models"
)

func insertTestCategory() {
	// Insert a category if it doesn't not exists
	categoryId := 533

	cnt, err := DB.Where("cat_id = ?", categoryId).Count(&Category{})
	if err != nil {
		fmt.Println("Error querying categories")
		return
	}
	if cnt > 0 {
		// Category with ID 533 already exists.
		return
	}

	// Create new category
	newCat := &Category{
		CatId: categoryId,
		Name:  "iPhone",
	}
	err = DB.Create(newCat)
	if err != nil {
		fmt.Println("Error creating test category!")
	}
}

func TestGetCategory(t *testing.T) {
	insertTestCategory()

	// Test returning category
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
	var cat Category
	err = dec.Decode(&cat)

	// Check if the decoding failed
	if err != nil {
		t.Error("Category decoding from JSON failed!")
	}

	// Check the decoded category ID
	expectedId := 533
	if cat.CatId != expectedId {
		t.Errorf("Category ID expected value: %d, received: %d", expectedId, cat.ID)
	}

	// Check the decoded category name
	expectedName := "iPhone"
	if cat.Name != expectedName {
		t.Errorf("Category name expected value: %v, received: %v", expectedName, cat.Name)
	}
}
