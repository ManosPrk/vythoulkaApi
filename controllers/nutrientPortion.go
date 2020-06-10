package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manosprk/vythoulka_api/models"
)

type nutrientPortionController struct {
	nutrientPortionRouter *mux.Router
}

func (ntp *nutrientPortionController) SetupRouter() {
	ntp.nutrientPortionRouter.HandleFunc("/", ntp.getAll)
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

func (ntp *nutrientPortionController) parseRequest(r *http.Request) (*models.NutrientPortion, error) {
	dec := json.NewDecoder(r.Body)
	var nutrient models.NutrientPortion
	err := dec.Decode(&nutrient)
	if err != nil {
		return &models.NutrientPortion{}, err
	}
	return &nutrient, nil
}

func newNutrientPortionController(subrouter *mux.Router) {
	ntp := nutrientPortionController{nutrientPortionRouter: subrouter}
	ntp.SetupRouter()
}
