package main

import (
	"encoding/json"
	"fmt"
	"hello/erros/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(700000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição de valor: ", err.Error())
		if err == model.ErrValorImovelAlto {
			fmt.Println("Corretor, por favor verifique a sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error())
		return
	}
	fmt.Println("A casa em JSON: ", string(objJSON))
}
