package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const qtdMonitoramento = 5
const delay = 5

func main() {
	// o for sem parâmetro é para loops infinitos
	for {
		exibeMenu()
		reader := leComando()
		switch reader {
		case 1:
			IniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs....")
			imprimeLog()
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

	sites := leSitesDoArquivo()

	for i := 0; i < qtdMonitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site: ", i, "site: ", site)
			testaSite(site)

		}
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	//o underline é para ignorar a referente variável em uma função com múltiplos retornos
	//response, _ := http.Get(site)
	//fmt.Print(response)

	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site: ", site, " Carregado com sucesso \n")
		registraLog(site, true)
		//fmt.Println(response)
	} else {
		fmt.Println("Site: ", site, "Está com problemas \n")
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, _ := os.Open("sites.txt")
	//arquivo, _ := os.ReadFile("sites.txt")
	//fmt.Println(arquivo)
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linhaSemEspaco := strings.TrimSpace(linha)
		sites = append(sites, linhaSemEspaco)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLog() {
	arquivo, _ := os.ReadFile("log.txt")
	fmt.Println(string(arquivo))
}
