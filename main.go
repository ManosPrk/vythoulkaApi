package main

import (
	"net/http"

	"github.com/manosprk/vythoulka_api/controllers"
	"github.com/manosprk/vythoulka_api/models"
)

func main() {
	controllers.RegisterControllers()
	models.NewGormDb("root:19121991@tcp(127.0.0.1:3308)/vythoulka?charset=utf8&parseTime=True&loc=Local")
	http.ListenAndServe(":3000", nil)
}
