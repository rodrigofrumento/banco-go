package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular              clientes.Titular
	NumAgencia, NumConta int
	Saldo                float64
}

func (c *ContaCorrente) Saque(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito > 0 {
		c.Saldo += valordoDeposito
		return "Deposito realizado", c.Saldo
	} else {
		return "Valor deve ser maior que zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia <= c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}
