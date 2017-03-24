package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/branohricardo/workshop-category-microservice/logger"
	"github.com/branohricardo/workshop-category-microservice/models"

	"github.com/gorilla/mux"
)

func init() {
	logger.New()
}

func main() {

	logger.Log.Info("Starting server")
	router := Router()
	http.ListenAndServe(":8080", router)
}

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/category/{ID}", getCategory)
	return router
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	paramID := mux.Vars(r)["ID"]
	catID, err := strconv.Atoi(paramID)

	if err != nil {
		logger.Log.Error(err, paramID)
		return
	}

	cat := models.Category{}
	err = models.DB.Where("cat_id = ?", catID).First(&cat)

	if err != nil {
		logger.Log.Error(err, fmt.Sprintf("Error finding category with ID: %v", catID))
		return
	}

	c, err := json.Marshal(cat)

	if err != nil {
		logger.Log.Error(err)
	}
	fmt.Fprintf(w, "%s", c)
}
