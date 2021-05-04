package main

import (
	"github.com/Firmwave11/otten-coffee/controllers"
	"github.com/Firmwave11/otten-coffee/routes"
	"github.com/Firmwave11/otten-coffee/usecase"
)

func main() {

	uc := usecase.NewUC()
	ctrl := controllers.NewCtrl(uc)

	router := routes.NewRouter(ctrl)
	router.Router("8000")
}
