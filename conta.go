package main

import (
	"fmt"
	"os"
)

type conta struct {
	nome	string
	itens []item
	gorjeta float64
}

func newConta(nome string) conta{
	c:= conta{
		nome: nome,
		itens: []item{},
		gorjeta: 0,
	}
	return c
}

func (c *conta) adicionarItem(nome string, preco float64){
	itm:= newItem(nome, preco)
	c.itens = append(c.itens, itm)
}

func (c *conta) calcularTotal() float64{
	total:= 0.0
	for _, item:= range c.itens{
		total+= item.preco
	}
	return total
}

func (c *conta) adicionarGorjeta(gorjeta float64){
	c.gorjeta = gorjeta
}

func (c *conta) calcularTotalComGorjeta() float64{
	return c.calcularTotal() + c.gorjeta
}

func (c *conta) removerItem(nome string) bool{
	for index, item:= range c.itens{
		if item.nome == nome{
			c.itens = append(c.itens[:index], c.itens[index+1:]...)
			return true
		}
	}
	return false
}

func (c *conta) getDados() string{
	dados:= fmt.Sprintf("Conta de %v\n----------------------------\n", c.nome)
	for index, item:= range c.itens{
		dados+= fmt.Sprintf("Item %v: %s - R$%.2f\n", index, item.nome, item.preco)
	}
	dados+= fmt.Sprintf("----------------------------\nTotal: R$%.2f\n", c.calcularTotal())
	dados+= fmt.Sprintf("Gorjeta: R$%.2f\n", c.gorjeta)
	dados+= fmt.Sprintf("Total com gorjeta: R$%.2f\n----------------------------\n", c.calcularTotalComGorjeta())
	return dados
}

func (c *conta) salvar(){
	data := []byte(c.getDados())

	err := os.WriteFile("contas/"+c.nome+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
}