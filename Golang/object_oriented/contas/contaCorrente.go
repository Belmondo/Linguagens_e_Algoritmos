package contas

import "object_oriented/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	Saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "Depósito não permitido", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.Saldo >= valorDaTransferencia && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Saldo += valorDaTransferencia
		return true
	} else {
		return false
	}
}
