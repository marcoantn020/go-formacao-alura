package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso."
	}
	return "Valor insuficiente."
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor
		return "Deposito realizado com sucesso", c.saldo
	}
	return "valor para deposito deve ser maior que 0", valor
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	podeTransferir := valor < c.saldo && valor > 0
	if podeTransferir {
		c.saldo -= valor
		contaDestino.saldo += valor
		return true
	}
	return false
}

func (c ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
