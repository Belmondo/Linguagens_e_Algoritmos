package main

import "fmt"

func main() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	//first method
	//var reader int
	//fmt.Scanf("%d", &reader)

	//second method to get
	var reader int
	fmt.Scan(&reader)

	fmt.Println("O comando escolhido foi: ", reader)

	//if reader == 1 {
	//	fmt.Println("Monitorando....")
	//} else if reader == 2 {
	//		fmt.Println("Exibindo Logs....")
	//} else if reader == 0 {
	//		fmt.Println("Saindo do Programa")
	//} else {
	//		fmt.Println("Não reconheço este comando")
	//	}

	switch reader {
	case 1:
		fmt.Println("Monitorando....")
	case 2:
		fmt.Println("Exibindo Logs....")
	case 0:
		fmt.Println("Saindo do Programa")
	default:
		fmt.Println("Não reconheço este comando")
	}

}
