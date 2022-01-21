package models

import "go-fundamentos-aplicacao-web/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectTodosOsProdutos, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = selectTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(idProduto string) {
	db := db.ConectaComBancoDeDados()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(idProduto)

	defer db.Close()
}

func EditarProduto(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", idProduto)
	if err != nil {
		panic(err.Error())
	}

	var produto = Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()

	return produto
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
