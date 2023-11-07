package main

import (
	"fmt"
	"os"
)

func main() {

	exibirIntroducao()
	exibirMenu()

	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Iniciar Monitoramento")
	case 2:
		fmt.Println("Exibir Logs")
	case 0:
		fmt.Println("Sair do Programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando!")
		os.Exit(-1)
	}
}

func exibirIntroducao() {
	nome := "Sérgio"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa esta na versão", versao)
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
