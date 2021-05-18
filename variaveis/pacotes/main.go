package main

import (
	"fmt"
	"hello/variaveis/pacotes/operadora"
	"hello/variaveis/pacotes/prefixo"
)

var NomedoUsuario = "Glênio"

func main() {
	fmt.Printf("nome do user é: %s\r\n", NomedoUsuario)
	fmt.Printf("Prefixo de SP: %d\r\n", prefixo.Capital)
	fmt.Printf("Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Teste Prefixo: %s\r\n", prefixo.TesteComPrefixo)
}
