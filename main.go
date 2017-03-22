package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Category struct {
	ID   int
	Name string
}

func main() {
	http.HandleFunc("/category", getCategory)
	http.ListenAndServe(":8080", nil)
}

func getCategory(w http.ResponseWriter, r *http.Request) {

	catID, _ := strconv.Atoi(r.URL.Query().Get("ID"))
	cat := Category{
		ID:   catID,
		Name: "iPhone",
	}
	c, err := json.Marshal(cat)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Fprintf(w, "%s", c)
}
