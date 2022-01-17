package routes

import (
	"go-fundamentos-aplicacao-web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
