package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valorDoSaque float64) string {
	saque := valorDoSaque <= conta.saldo
	if saque {
		conta.saldo -= valorDoSaque
		return "Saque realizado com sucesso !"
	} else {
		return "Saldo insuficiente !"
	}
}

func (conta *ContaCorrente) Depositar(valorDoDeposito float64) string {
	deposito := valorDoDeposito >= conta.saldo
	if deposito {
		conta.saldo += valorDoDeposito
		return "Deposito realizado com sucesso !"
	} else {
		return "Não foi possivel realizar o Deposito !"
	}
}

func (conta *ContaCorrente) Transferir(valorDaTransferencia float64) string {
	transferencia := valorDaTransferencia <= conta.saldo
	if transferencia {
		conta.saldo -= valorDaTransferencia
		return "Tranferencia realizada com sucesso !"
	} else {
		return "Saldo insuficiente !"
	}
}

func main() {
	var contaVinicius = ContaCorrente{
		titular:       "Vinicius",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50}

	fmt.Println(contaVinicius.Depositar(300))
	fmt.Println(contaVinicius.saldo)

}
