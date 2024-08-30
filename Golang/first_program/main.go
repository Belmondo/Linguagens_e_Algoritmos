package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {

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
			IniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs....")
		case 0:
			fmt.Println("Saindo do Programa")
			saiPrograma()
		default:
			fmt.Println("Não reconheço este comando")
		}
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

func IniciarMonitoramento() {
	fmt.Println("Monitorando....")
	site := "https://www.alura.com.br/"
	response, _ := http.Get(site)
	//fmt.Print(response)

	if response.StatusCode == 200 {
		fmt.Print("Site: ", site, "Carregado com sucesso \n")
	} else {
		fmt.Print("Site: ", site, "Está com problemas \n")
	}
}
