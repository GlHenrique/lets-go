package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Print("o número ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("Você é da família do UNIX?")

	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim!")
	default:
		fmt.Println("Deixa essa pergunta para lá")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Opa, dorme ai campeão")
	default:
		fmt.Println("Acorda ai maluco.. Vai trabalhar")
	}

	numero = 10

	fmt.Print("Este número cabe num dígito? ")
	switch {
	case numero < 10:
		fmt.Println("Sim!")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois dígitos?...")
	case numero >= 100:
		fmt.Println("Ai é demais...")
	}

}
