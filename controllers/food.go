package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/manosprk/vythoulka_api/models"
)

const (
	name   = "name"
	nameGr = "nameGr"
)

type foodController struct {
	foodRouter *mux.Router
}

type foodRequest struct {
	Name   string `schema:"name"`
	NameGr string `schema:"nameGR"`
	Offset int    `schema:"offset"`
	Limit  int    `validate:"required,min=1,max=10" schema:"limit"`
}

func (fc *foodController) SetupRouter() {
	fc.foodRouter.Path("/").
		Queries(name, "{"+name+"}").
		HandlerFunc(fc.getByName).
		Name(name)

	fc.foodRouter.Path("/").
		Queries(nameGr, "{"+nameGr+"}").
		HandlerFunc(fc.getByNameGr).
		Name(nameGr)

	fc.foodRouter.HandleFunc("/{id:[0-9]+}", fc.getByApiId)
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

func (fc *foodController) getByName(w http.ResponseWriter, r *http.Request) {
	fr, err := fc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	foods, err := models.GetFoodsByName(fr.Name, fr.Limit, fr.Offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(foods, w)
}

func (fc *foodController) getByNameGr(w http.ResponseWriter, r *http.Request) {
	fr, err := fc.parseRequest(r)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	foods, err := models.GetFoodsByNameGr(fr.NameGr, fr.Limit, fr.Offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(foods, w)
}

func (fc *foodController) parseRequest(r *http.Request) (*foodRequest, error) {
	var fr foodRequest
	err := schema.NewDecoder().Decode(&fr, r.URL.Query())

	if err != nil {
		return &foodRequest{}, err
	}

	v := validator.New()

	err = v.Struct(fr)
	if err != nil {
		return &foodRequest{}, err
	}

	return &fr, nil
}

func newFoodController(subrouter *mux.Router) {
	fc := foodController{foodRouter: subrouter}
	fc.SetupRouter()
}
