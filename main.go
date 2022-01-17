package main

import (
	"go-fundamentos-aplicacao-web/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":3000", nil)
}
