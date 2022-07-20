package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	//"reflect"
)

const monitoramento = 2
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa")
		// } else {
		// 	fmt.Println("Não conheço esse comando")
		// }

		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {

	nome := "Alexandre"
	versao := 1.1

	fmt.Println("Olá sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://random-status-code.herokuapp.com/",
	// 	"https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")

	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreou um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreou um erro:", err)
	}
	//arquivo, err := ioutil.ReadFile("sites.txt")
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(string(linha))
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreou um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:06") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLog() {

	fmt.Println("Exibindo Logs...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreou um erro:", err)
	}

	fmt.Println(string(arquivo))

}
