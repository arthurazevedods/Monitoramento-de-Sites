package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const delay = 5
const ciclos_monitoramento = 4

func main() {

	introducao()
	registrarLog("algo.com", false)
	for {
		comando := menu()

		switchComandos(comando)
	}

}

func introducao() {
	nome := "Arthur"

	versao := 0.7

	fmt.Println("\nOlá, sr.", nome)
	fmt.Println("Esse programa está na versão:", versao)
}

func menu() int {
	fmt.Println("MENU")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando digitado foi:", comando)
	return comando
}

func switchComandos(comando int) {
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
		iniciarMonitoramento()
	case 2:
		fmt.Println("Imprimindo os Logs...")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	sites, _ := leSitesDoArquivo()
	for i := 0; i < ciclos_monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func exibirLogs() {
	fmt.Println("Exbindo Logs...")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registrarLog(site, true)
	}
}

func leSitesDoArquivo() ([]string, error) {
	// Abrir o arquivo JSON
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório:", err)
	} else {
		fmt.Println("Diretório atual:", dir)
	}
	file, err := os.Open("sites.json")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return nil, err
	}
	defer file.Close()

	// Decodificar o JSON diretamente do arquivo
	var data struct {
		Sites []string `json:"sites"`
	}

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nOs sites lidos são:")
	for i, site := range data.Sites {
		fmt.Println(i, " - ", site)
	}
	// Retornar o array de sites
	return data.Sites, nil
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
