package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var contaVinicius = ContaCorrente{
		titular:       "Vinicius",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50}

	var contaPedro = ContaCorrente{
		titular:       "Pedro",
		numeroAgencia: 222,
		numeroConta:   200,
		saldo:         128.96}

	fmt.Println(contaVinicius)
	fmt.Println(contaPedro)
}
