package main

import (
	"fmt"
	"object_oriented/contas"
)

func main() {
	// contaTeste := ContaCorrente{Titular: "Belmondo", numeroAgencia: 123456, numeroConta: 654321, Saldo: 70.00}
	// fmt.Println(contaTeste)

	// var contaTesteDois *ContaCorrente
	// contaTesteDois = new(ContaCorrente)
	// contaTesteDois.Titular = "Jr"
	// contaTesteDois.Saldo = 500

	// fmt.Println(*contaTesteDois)

	// contaTesteTres := ContaCorrente{}
	// contaTesteTres.Titular = "Aragorn"
	// contaTesteTres.Saldo = 1000

	// fmt.Println(contaTesteTres.Saldo)
	// fmt.Println(contaTesteTres.Sacar(300))
	// fmt.Println(contaTesteTres.Saldo)

	// status, SaldoAtual := contaTesteTres.Depositar(500)
	// fmt.Println(status)
	// fmt.Println("Saldo Atual", SaldoAtual)

	contaAragorn := contas.ContaCorrente{Titular: "Aragorn", Saldo: 300}
	contaLegolas := contas.ContaCorrente{Titular: "Legolas", Saldo: 100}

	statusTransferencia := contaAragorn.Transferir(200, &contaLegolas)

	fmt.Println(statusTransferencia)
	fmt.Println(contaAragorn)
	fmt.Println(contaLegolas)
}
