package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/branohricardo/workshop-category-microservice/models"
	"github.com/gorilla/mux"
)

func main() {
	router := Router()
	http.ListenAndServe(":8080", router)
}

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/category/{ID}", getCategory)
	return router
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	catID, err := strconv.Atoi(mux.Vars(r)["ID"])
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	cat := models.Category{}
	err = models.DB.Where("cat_id = ?", catID).First(&cat)

	if err != nil {
		fmt.Printf("Error finding category with ID: %v, %v", catID, err)
		return
	}

	c, err := json.Marshal(cat)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Fprintf(w, "%s", c)
}
