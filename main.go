package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput(prompt string, reader *bufio.Reader) (string, error) {

	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	tableNumber, _ := getUserInput("Digite o numero da mesa para abertura de conta: ", reader)

	t, err := strconv.Atoi(tableNumber)

	if err != nil {
		fmt.Println("A mesa só pode ser número")
	}

	createdBill := openNewBill(t)

	fmt.Printf("Conta da mesa %v criada com sucesso\n", tableNumber)

	return createdBill
}

func programExecution(b Bill) {
	reader := bufio.NewReader(os.Stdin)

	action, _ := getUserInput("Escolha a opção desejada\n(a - Adicionar Itens, r - Remover um item, u - Corrigir preço de item, c - Alterar numero de clientes, t - Adicionar gorjeta, s - Salvar e imprimir conta)\n", reader)

	switch action {
	case "a":
		item, _ := getUserInput("Item: ", reader)
		price, _ := getUserInput("Preço: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("O preço deve ser um número.")
			programExecution(b)
		}

		b.addNewItem(item, p)

		fmt.Printf("%v adicionado com sucesso.\n", item)
		programExecution(b)

	case "r":
		item, _ := getUserInput("Item a ser removido: ", reader)

		b.removeItem(item)

		fmt.Printf("%v removido da conta.\n", item)
		programExecution(b)

	case "u":
		item, _ := getUserInput("Item para atualização: ", reader)
		price, _ := getUserInput("Novo valor do item: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("O preço deve ser um número.")
			programExecution(b)
		}

		b.updateItemValue(item, p)

		fmt.Printf("Valor do item %v atualizado para %v\n", item, p)
		programExecution(b)

	case "c":
		numberOfClients, _ := getUserInput("Digite a quantidade de clientes da mesa: ", reader)

		n, err := strconv.Atoi(numberOfClients)

		if err != nil {
			fmt.Println("Numero de clientes inválido")
			programExecution(b)
		}

		b.updateAmoutOfConsumers(n)

		fmt.Println("Numero de clientes atualizado.")
		programExecution(b)

	case "t":
		tip, _ := getUserInput("Valor da gorjeta a ser dada: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("O valor digitado é inválido")
			programExecution(b)
		}

		b.addTip(t)

		fmt.Println("Gorjeta adicionada a conta da mesa.")
		programExecution(b)

	case "s":
		b.saveBill()
		fmt.Println("Conta da mesa " + strconv.Itoa(b.table) + " encerrada e salva.")
	default:
		fmt.Println("Opção inválida...")
		programExecution(b)
	}
}

func main() {
	bill := createBill()
	programExecution(bill)
}
