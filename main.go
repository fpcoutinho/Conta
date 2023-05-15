package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	conta:= criaConta()
	prompOptions(conta)
	
}

func criaConta() conta{
	reader := bufio.NewReader(os.Stdin)
	
	nome, _ := getInput("Digite o nome do cliente: ", reader)
	conta := newConta(nome)
	fmt.Printf("Conta de %v criada com sucesso!\n" , nome)

	return conta
}

func getInput(prompt string, reader *bufio.Reader) (string, error){
	fmt.Print(prompt)
	text, erro := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text, erro
}

func prompOptions (c conta){
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Digite a opção desejada: \n1 - adicionar item\n2 - adicionar gorjeta\n3 - salvar conta\n4 - remover item\n", reader)
	switch option {
		case "1":
			nome , _ := getInput("Digite o nome do item: ", reader)
			preco , _ := getInput("Digite o preço do item: ", reader)
			p, e := strconv.ParseFloat(preco, 64)

			if e != nil{
				fmt.Println("O preço deve ser um número!")
				prompOptions(c)
			}

			c.adicionarItem(nome, p)
			fmt.Printf("Item adicionado: %v - R$%v\n", nome, preco)
			prompOptions(c)

		case "2":
			gorjeta , _ := getInput("Digite o valor da gorjeta: ", reader)
			g, e := strconv.ParseFloat(gorjeta, 64)

			if e != nil{
				fmt.Println("A gorjeta deve ser um número!")
				prompOptions(c)
			}

			c.adicionarGorjeta(g)
			fmt.Printf("Gorjeta adicionada: R$%v\n", gorjeta)
			prompOptions(c)

		case "3":
			c.salvar()
			fmt.Printf("Conta de %v salva com sucesso!\n", c.nome)
			prompOptions(c)
		
		case "4":
			nome , _ := getInput("Digite o nome do item: ", reader)
			if(c.removerItem(nome)){
				fmt.Printf("Item %v removido com sucesso!\n", nome)
			}else{
				fmt.Printf("Item %v não encontrado!\n", nome)
			}
			prompOptions(c)

		default:
			fmt.Println("Opção inválida")
			prompOptions(c)
	}
}
