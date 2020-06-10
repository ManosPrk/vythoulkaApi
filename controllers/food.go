package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manosprk/vythoulka_api/models"
)

type foodController struct {
	foodRouter *mux.Router
}

func (fc *foodController) SetupRouter() {
	fc.foodRouter.HandleFunc("/", fc.GetIndexHandler)
	fc.foodRouter.HandleFunc("/{id:[0-9]+}", fc.getByApiId)
}

func (fc *foodController) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	food, err := fc.parseRequest(r)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	switch {
	case food.Name != "":
		fc.getByName(food, w)
	case food.NameGr != "":
		fc.getByNameGr(food, w)
	default:
		fc.getAll(w, r)
	}
}

func (fc *foodController) getAll(w http.ResponseWriter, r *http.Request) {
	foods, err := models.GetFoods()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	encodeResponseAsJSON(foods, w)
}

func (fc *foodController) getByApiId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	food, err := models.GetFoodById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encodeResponseAsJSON(food, w)
}

func (fc *foodController) getByName(food *models.Food, w http.ResponseWriter) {
	foods, err := models.GetFoodsByName(food.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(foods, w)
}

// func (fc *foodController) getByFilter(food *models.Food, w http.ResponseWriter) {
// 	foods, err := models.GetFoodsByNameGr(food.NameGR)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	encodeResponseAsJSON(foods, w)
// }

func (fc *foodController) getByNameGr(food *models.Food, w http.ResponseWriter) {
	foods, err := models.GetFoodsByNameGr(food.NameGr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(foods, w)
}

func (fc *foodController) parseRequest(r *http.Request) (*models.Food, error) {
	dec := json.NewDecoder(r.Body)
	var food models.Food
	err := dec.Decode(&food)
	if err != nil {
		return &models.Food{}, err
	}
	return &food, nil
}

func newFoodController(subrouter *mux.Router) {
	fc := foodController{foodRouter: subrouter}
	fc.SetupRouter()
}
