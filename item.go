package main

type item struct {
	nome string
	preco float64
}

func newItem(nome string, preco float64) item{
	itm:= item{
		nome: nome,
		preco: preco,
	}
	return itm
}