package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaMarco := contas.ContaCorrente{}
	contaMarco.Depositar(500)
	PagarBoleto(&contaMarco, 200)

	contaMaria := contas.ContaPoupanca{}
	contaMaria.Depositar(400)
	PagarBoleto(&contaMaria, 100)

	fmt.Println(contaMarco.ObterSaldo())
	fmt.Println(contaMaria.ObterSaldo())
}
