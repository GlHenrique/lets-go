package main

import (
	"encoding/json"
	"fmt"
	"hello/structs2/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Verde"
	casa.X = 18
	casa.Y = 50
	casa.SetValor(60000)

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("Casa valor da casa é: %d\r\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}
