package main

import "fmt"

// Variáveis públicas são escritas com a primeira letra maiúsculas

var (
	// Endereco é string "" e é pública
	Endereco            string
	telefone            string  // endereco = "" e é pública
	quantidade          int     // quantidade = 0
	comprou             bool    // comprou false
	valor               float64 //valor = 0.00
	caracteresEspeciais rune
)

func main() {
	teste := "Go!"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("Caracteres Especiais: %v\r\n", caracteresEspeciais)
	fmt.Printf("Teste: %v\r\n", teste)
}
