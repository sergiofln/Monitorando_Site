package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibirIntroducao()
	for {
		exibirMenu()

		comando := lerComando()

		switch comando {
		case 1:
			IniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1)
		}
	}
}

func exibirIntroducao() {
	nome := "Sérgio"
	versao := 1.1
	fmt.Println("Olá, Srº", nome)
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

func IniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// Site com URL inexistente
	site := "https://httpbin.org/status/404" // ou 200
	// Site existente
	//site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
