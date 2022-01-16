package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectaComBancoDeDados()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortavel", 89, 10},
		{"Fone", "Intrauricular", 120, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
