package main

import (
	"net/http"

	"github.com/manosprk/vythoulka_api/controllers"
	"github.com/manosprk/vythoulka_api/models"
)

func main() {
	models.NewDB("root:19121991@tcp(127.0.0.1:3308)/vythoulka")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
