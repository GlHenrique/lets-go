package main

import "fmt"

func main() {
	meses := 0
	situacao := true
	cidade := "São Paulo"

	// < > != == <= >= && ||

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Ele está inadimplente")
	}

	if !situacao {
		fmt.Println("Ele está anadimplente")
	}

	if cidade == "São Paulo" {
		fmt.Println("Ele vive no UF SP")
	}

	if descricao, status := haQuantoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")

}

func haQuantoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo."
		return
	}

	descricao = "O cliente está com o carnê em dia."
	return

}
