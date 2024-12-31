package main

import (
	"fmt"
	"object_oriented/clientes"
	"object_oriented/contas"
)

func main() {

	contaAragorn := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Aragorn", CPF: "123.456.789.10", Profissao: "Future King"}, Saldo: 300}
	contaLegolas := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Legolas", CPF: "321.654.987.89", Profissao: "Archer"}, Saldo: 100}

	statusTransferencia := contaAragorn.Transferir(200, &contaLegolas)

	fmt.Println(statusTransferencia)
	fmt.Println(contaAragorn)
	fmt.Println(contaLegolas)
}
