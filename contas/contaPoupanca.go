package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	Saldo                                float64
}

func (c *ContaPoupanca) Saque(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito > 0 {
		c.Saldo += valordoDeposito
		return "Deposito realizado", c.Saldo
	} else {
		return "Valor deve ser maior que zero", c.Saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.Saldo
}
