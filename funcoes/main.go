package main

import (
	"fmt"
	"hello/funcoes/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 10, 20)
	fmt.Printf("Multiplicação é %d\r\n", resultado)
	resultado = matematica.Soma(2, 2)
	fmt.Printf("Soma é %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 10, 2)
	fmt.Printf("O resultado da divisão é %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(15, 6)
	fmt.Printf("O resultado da divisão com resto é %d e o resto é %d\r\n", resultado, resto)
}
