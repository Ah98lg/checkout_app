package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bill struct {
	table     int
	orders    map[string]float64
	total     float64
	tip       float64
	consumers int
}

func openNewBill(table int) Bill {
	newBill := Bill{
		table:     table,
		orders:    map[string]float64{},
		total:     0.0,
		tip:       0.0,
		consumers: 1,
	}

	return newBill
}

func (b *Bill) addNewItem(item string, price float64) {
	b.orders[item] = price
}

func (b *Bill) removeItem(item string) {
	delete(b.orders, item)
}

func (b *Bill) addTip(value float64) {
	b.tip = value
}

func (b *Bill) updateItemValue(item string, value float64) {
	b.orders[item] = value
}

func (b *Bill) updateAmoutOfConsumers(consumers int) {
	b.consumers = consumers
}

func (b *Bill) formatBill() string {

	formatedBill := "Resumo da conta: \n"

	separator := fmt.Sprintf("\n%-25v\n", "")
	strings.ReplaceAll(separator, " ", "-")

	formatedBill += separator

	for k, v := range b.orders {
		formatedBill += fmt.Sprintf("%-25v ...%v\n", k+":", v)
		b.total += v
	}

	formatedBill += separator

	formatedBill += fmt.Sprintf("%-25v ...%v\n", "Gorjeta:", b.tip)

	formatedBill += separator

	formatedBill += fmt.Sprintf("%-25v ...%v\n", "Valor total:", b.total+b.tip)

	return formatedBill
}

func (b *Bill) saveBill() {
	data := []byte(b.formatBill())
	err := os.WriteFile("Bills/"+strconv.Itoa(b.table)+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conta encerrada e salva com sucesso.")
}
