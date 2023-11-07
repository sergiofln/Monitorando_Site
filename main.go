package main

import "fmt"

func main() {
	nome := "Sérgio"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa esta na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Iniciar Monitoramento")
	case 2:
		fmt.Println("Exibir Logs")
	case 0:
		fmt.Println("Sair do Programa")
	default:
		fmt.Println("Não conheço este comando!")
	}
}
