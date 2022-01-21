package controllers

import (
	"go-fundamentos-aplicacao-web/models"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")

	produto := models.EditarProduto(idDoProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.AtualizaProduto(idInt, nome, descricao, precoFloat, quantidadeInt)

		http.Redirect(w, r, "/", 301)
	}
}
