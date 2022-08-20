package main

import (
	"exemplo/golang/bank/contas"
	"fmt"
)

// type ContaCorrente struct {
// 	Titular       string
// 	NumeroAgencia int
// 	NumeroConta   int
// 	Saldo         float64
// }

// func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
// 	podeSacar := valorDoSaque > -1 && valorDoSaque <= c.Saldo
// 	if podeSacar {
// 		c.Saldo -= valorDoSaque
// 		return "Saque realizado com sucesso"
// 	} else {
// 		return "Saldo insuficiente"
// 	}
// }

// func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
// 	if valorDoDeposito > -1 {
// 		c.Saldo += valorDoDeposito
// 		return "Deposito realizado com sucesso", c.Saldo
// 	} else {
// 		return "Valor do deposito menor que zero", c.Saldo
// 	}
// }

// func (c *ContaCorrente) Tranferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
// 	if valorDaTransferencia < c.Saldo && valorDaTransferencia > -1 {
// 		c.Saldo -= valorDaTransferencia
// 		contaDestino.Depositar(valorDaTransferencia)
// 		return true
// 	} else {
// 		return false
// 	}
// }

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Tranferir(-200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
