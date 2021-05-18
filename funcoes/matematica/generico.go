package matematica

/*
Calculo executa qualquer tipo de cálculo, basta passar uma função como parâmetro
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

// Multiplica dois inteiros
func Multiplicador(x int, y int) int {
	return x * y
}

// Divisor efetua a divisão de dois inteiros
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

// Retorno o resultado e resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
