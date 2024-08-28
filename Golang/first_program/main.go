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

}
