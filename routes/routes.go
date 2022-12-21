package routes

import (
	"AMM/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.IndexHanddler)
}
