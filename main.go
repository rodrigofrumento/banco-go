package main

import (
	c "banco/contas"
	"fmt"
)

func PagarBoleto(c verificarConta, valorBoleto float64) {
	c.Saque(valorBoleto)
}

type verificarConta interface {
	Saque(valor float64) string
}

func main() {
	contaDoRodrigo := c.ContaPoupanca{}
	contaDoRodrigo.Depositar(100)
	PagarBoleto(&contaDoRodrigo, 60)

	fmt.Println(contaDoRodrigo.ObterSaldo())
}
