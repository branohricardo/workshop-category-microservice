package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type category struct {
	ID   int
	Name string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/category/{ID}", getCategory)
	http.ListenAndServe(":8080", router)
}

func getCategory(w http.ResponseWriter, r *http.Request) {

	catID, err := strconv.Atoi(mux.Vars(r)["ID"])

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	cat := category{
		ID:   catID,
		Name: "iPhone",
	}
	c, err := json.Marshal(cat)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Fprintf(w, "%s", c)
}
