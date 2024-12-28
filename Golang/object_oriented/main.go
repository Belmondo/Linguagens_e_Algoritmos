package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Depósito não permitido", c.saldo
	}
}

func main() {
	// contaTeste := ContaCorrente{titular: "Belmondo", numeroAgencia: 123456, numeroConta: 654321, saldo: 70.00}
	// fmt.Println(contaTeste)

	// var contaTesteDois *ContaCorrente
	// contaTesteDois = new(ContaCorrente)
	// contaTesteDois.titular = "Jr"
	// contaTesteDois.saldo = 500

	// fmt.Println(*contaTesteDois)

	contaTesteTres := ContaCorrente{}
	contaTesteTres.titular = "Aragorn"
	contaTesteTres.saldo = 1000

	fmt.Println(contaTesteTres.saldo)
	fmt.Println(contaTesteTres.Sacar(300))
	fmt.Println(contaTesteTres.saldo)

	status, saldoAtual := contaTesteTres.Depositar(500)
	fmt.Println(status)
	fmt.Println("Saldo Atual", saldoAtual)
}
