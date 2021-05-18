package main

import "fmt"

//Imovel é uma struct que armazena dados de um imvóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Apartamento do Berin", 170000}
	fmt.Printf("O apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chácara",
		X:     22,
		valor: 300000,
	}
	fmt.Printf("A chácara é: %+v\r\n", chacara)

	casa.Nome = "Lar do Berin"
	casa.valor = 500000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v\r\n", casa)
}
