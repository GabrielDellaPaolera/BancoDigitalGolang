package main

import (
	"Teste/Alura/BancoDigital/contas"
	"fmt"
)

func PagarBoleto(conta VerificarConta, ValorDoBoleto float64) {
	conta.Sacar(ValorDoBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {
	// ContaDoGabriel := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome:      "Gabriel",
	// 	CPF:       "420",
	// 	Profissao: "arquiteto"},
	// 	NumeroAgencia: 101,
	// 	NumeroConta:   2020,
	// 	saldo:         10}
	// fmt.Println(ContaDoGabriel)

	contaDaRafa := contas.ContaCorrente{}
	contaDaRafa.Depositar(100)
	fmt.Println(contaDaRafa.Sacar(50))
	PagarBoleto(&contaDaRafa, 10)
	fmt.Println(contaDaRafa.Obtersaldo())

	contaDoGabriel := contas.ContaPoupanca{}
	contaDoGabriel.Depositar(500)
	PagarBoleto(&contaDoGabriel, 30)
	fmt.Println(contaDoGabriel.Obtersaldo())

}
