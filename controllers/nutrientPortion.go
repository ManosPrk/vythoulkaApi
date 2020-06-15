package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/manosprk/vythoulka_api/models"
)

const (
	foodId = "foodId"
)

type nutrientPortionController struct {
	nutrientPortionRouter *mux.Router
}

type nutrientPortionRequest struct {
	FoodId string `schema:"foodId"`
}

func (ntp *nutrientPortionController) SetupRouter() {
	ntp.nutrientPortionRouter.Path("/").
		Queries(foodId, "{"+foodId+"}").
		HandlerFunc(ntp.getByFoodId).
		Name(foodId)

	ntp.nutrientPortionRouter.HandleFunc("/{id:[0-9]+}", ntp.getById)
}

func (ntp *nutrientPortionController) getAll(w http.ResponseWriter, r *http.Request) {
	nutrientPortions, err := models.GetNutrientPortions()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	encodeResponseAsJSON(nutrientPortions, w)
}

func (ntp *nutrientPortionController) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nutrientPortion, err := models.GetNutrientPortionById(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encodeResponseAsJSON(nutrientPortion, w)
}

func (ntp *nutrientPortionController) getByFoodId(w http.ResponseWriter, r *http.Request) {
	fr, err := ntp.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	foodId, err := strconv.Atoi(fr.FoodId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	nutrientPortions, err := models.GetNutrientPortionsByFoodId(foodId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	encodeResponseAsJSON(nutrientPortions, w)
}

func (ntp *nutrientPortionController) parseRequest(r *http.Request) (*nutrientPortionRequest, error) {
	var ntpr nutrientPortionRequest

	err := schema.NewDecoder().Decode(&ntpr, r.URL.Query())
	if err != nil {
		return &nutrientPortionRequest{}, err
	}

	return &ntpr, nil
}

func newNutrientPortionController(subrouter *mux.Router) {
	ntp := nutrientPortionController{nutrientPortionRouter: subrouter}
	ntp.SetupRouter()
}
