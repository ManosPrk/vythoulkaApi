package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manosprk/vythoulka_api/models"
)

type nutrientController struct {
	nutrientRouter *mux.Router
}

func (nt *nutrientController) SetupRouter() {
	nt.nutrientRouter.HandleFunc("/", nt.GetIndexHandler)
	nt.nutrientRouter.HandleFunc("/{id:[0-9]+}", nt.getByApiId)
}

func (nt *nutrientController) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	nutrient, err := nt.parseRequest(r)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	switch {
	case nutrient.Name != "":
		nt.getByName(nutrient, w)
	case nutrient.NameGr != "":
		nt.getByNameGr(nutrient, w)
	default:
		nt.getAll(w, r)
	}
}

func (nt *nutrientController) getAll(w http.ResponseWriter, r *http.Request) {
	nutrients, err := models.GetNutrients()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	encodeResponseAsJSON(nutrients, w)
}

func (nt *nutrientController) getByApiId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nutrient, err := models.GetNutrientById(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encodeResponseAsJSON(nutrient, w)
}

func (nt *nutrientController) getByName(nutrient *models.Nutrient, w http.ResponseWriter) {
	nutrients, err := models.GetNutrientsByName(nutrient.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(nutrients, w)
}

func (nt *nutrientController) getByNameGr(nutrient *models.Nutrient, w http.ResponseWriter) {
	nutrients, err := models.GetNutrientsByNameGr(nutrient.NameGr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(nutrients, w)
}

func (nt *nutrientController) parseRequest(r *http.Request) (*models.Nutrient, error) {
	dec := json.NewDecoder(r.Body)
	var nutrient models.Nutrient
	err := dec.Decode(&nutrient)
	if err != nil {
		return &models.Nutrient{}, err
	}
	return &nutrient, nil
}

func newNutrientController(subrouter *mux.Router) {
	nt := nutrientController{nutrientRouter: subrouter}
	nt.SetupRouter()
}
