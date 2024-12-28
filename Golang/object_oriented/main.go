package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaTeste := ContaCorrente{titular: "Belmondo", numeroAgencia: 123456, numeroConta: 654321, saldo: 70.00}
	fmt.Println(contaTeste)

	var contaTesteDois *ContaCorrente
	contaTesteDois = new(ContaCorrente)
	contaTesteDois.titular = "Jr"
	contaTesteDois.saldo = 500

	fmt.Println(*contaTesteDois)
}
