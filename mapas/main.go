package main

import (
	"fmt"
	"hello/structs2/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(300000)

	cache["Casa Amarela"] = casa

	fmt.Println("há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, achei, e o que eu achei foi: %+v\r\n", imovel)
	}

	apartamento := model.Imovel{}
	apartamento.Nome = "Apartamento Azul"
	apartamento.X = 12
	apartamento.Y = 20
	apartamento.SetValor(150000)

	cache[apartamento.Nome] = apartamento

	fmt.Println("quantos há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("chave[%s] = %+v\n\r", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		fmt.Println("Apagando Casa Amarela...")
		delete(cache, imovel.Nome)
	}

	fmt.Println("quantos há no cache? ", len(cache))

}
