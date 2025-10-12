package main

import (
	"fmt"
)

// Saldo agora é int para representar centavos. R$ 1.00 = 100
type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	saldo         int
}

// Sacar agora valida o valor e retorna um erro se falhar
func (c *ContaCorrente) Sacar(valorDoSaque int) error {
	if valorDoSaque <= 0 {
		return fmt.Errorf("valor do saque deve ser positivo")
	}

	if c.saldo < valorDoSaque {
		return fmt.Errorf("saldo insuficiente para realizar o saque")
	}

	c.saldo -= valorDoSaque
	return nil // nil significa "sem erro", ou seja, sucesso!
}

// Depositar agora valida o valor e a lógica está correta
func (c *ContaCorrente) Depositar(valorDoDeposito int) error {
	if valorDoDeposito <= 0 {
		return fmt.Errorf("valor do depósito deve ser positivo")
	}

	c.saldo += valorDoDeposito
	return nil
}

// Transferir agora recebe a conta de destino e é mais segura
func (cOrigem *ContaCorrente) Transferir(valorDaTransferencia int, cDestino *ContaCorrente) error {
	if valorDaTransferencia <= 0 {
		return fmt.Errorf("valor da transferência deve ser positivo")
	}

	if cOrigem.saldo < valorDaTransferencia {
		return fmt.Errorf("saldo insuficiente para realizar a transferência")
	}

	// Realiza as operações
	cOrigem.saldo -= valorDaTransferencia
	cDestino.saldo += valorDaTransferencia

	return nil
}

// Método para exibir o saldo formatado em Reais
func (c *ContaCorrente) ObterSaldo() float64 {
	return float64(c.saldo) / 100
}

func main() {
	// Usando centavos para os saldos iniciais
	contaVinicius := ContaCorrente{Titular: "Vinicius", saldo: 12550} // R$ 125,50
	contaGabriela := ContaCorrente{Titular: "Gabriela", saldo: 20000} // R$ 200,00

	fmt.Printf("Saldo inicial do Vinicius: R$ %.2f\n", contaVinicius.ObterSaldo())
	fmt.Printf("Saldo inicial da Gabriela: R$ %.2f\n\n", contaGabriela.ObterSaldo())

	// Tentando fazer um depósito
	valorDeposito := 5000 // R$ 50,00
	if err := contaVinicius.Depositar(valorDeposito); err != nil {
		fmt.Printf("Falha ao depositar R$ %.2f: %v\n", float64(valorDeposito)/100, err)
	} else {
		fmt.Printf("Depósito de R$ %.2f realizado com sucesso!\n", float64(valorDeposito)/100)
	}
	fmt.Printf("Novo saldo do Vinicius: R$ %.2f\n\n", contaVinicius.ObterSaldo())

	// Tentando fazer uma transferência
	valorTransferencia := 10000 // R$ 100,00
	fmt.Printf("Vinicius tentando transferir R$ %.2f para Gabriela...\n", float64(valorTransferencia)/100)
	if err := contaVinicius.Transferir(valorTransferencia, &contaGabriela); err != nil {
		fmt.Printf("Falha na transferência: %v\n", err)
	} else {
		fmt.Println("Transferência realizada com sucesso!")
	}

	fmt.Printf("Saldo final do Vinicius: R$ %.2f\n", contaVinicius.ObterSaldo())
	fmt.Printf("Saldo final da Gabriela: R$ %.2f\n", contaGabriela.ObterSaldo())
}
