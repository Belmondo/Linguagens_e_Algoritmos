package main

import (
	"fmt"
	"os"
)

func main() {

	exibeMenu()
	//first method
	//var reader int
	//fmt.Scanf("%d", &reader)

	//second method to get

	//if reader == 1 {
	//	fmt.Println("Monitorando....")
	//} else if reader == 2 {
	//		fmt.Println("Exibindo Logs....")
	//} else if reader == 0 {
	//		fmt.Println("Saindo do Programa")
	//} else {
	//		fmt.Println("Não reconheço este comando")
	//	}
	reader := leComando()
	switch reader {
	case 1:
		fmt.Println("Monitorando....")
	case 2:
		fmt.Println("Exibindo Logs....")
	case 0:
		fmt.Println("Saindo do Programa")
		saiPrograma()
	default:
		fmt.Println("Não reconheço este comando")
	}

}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

}

func leComando() int {
	var reader int
	fmt.Scan(&reader)

	fmt.Println("O comando escolhido foi: ", reader)
	return reader
}

func saiPrograma() {
	os.Exit(0)
}
