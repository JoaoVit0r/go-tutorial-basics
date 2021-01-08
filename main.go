// package name
package main

// imported packages
import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// global constants
const times = 1
const delay = 5

func main() {
	// infinite loop
	for {
		fmt.Println("=================================")
		fmt.Println("(1) Fazer Buscas")
		fmt.Println("(2) Exibir Historico")
		fmt.Println("(0) Sair do Programa")

		choice, err := getChoice()

		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}

		// alternatively -> switch choice, _ := getChoice(); choice {
		switch choice {
		case 1:
			fmt.Println("Iniciando Buscas...")
			startSearch()
		case 2:
			fmt.Println("Exibindo Histórico...")
			// showHistory()
		case 0:
			fmt.Println("Saindo do programa")
			// os.Exit(0)
			return
		default:
			fmt.Println("Opção Inválida")
		}
	}
}

func getChoice() (int, error) {
	var choice int

	fmt.Print("	Escolha: ")

	_, err := fmt.Scan(&choice)

	if err != nil {
		return -1, err
	}

	fmt.Println("__________________________________")

	return choice, nil
}

func getUrls(cep string, uf string) []string {

	var urls []string

	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("Error:", err)

		return urls
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		line = strings.Replace(line, "<CEP>", cep, -1)
		line = strings.Replace(line, "<UF>", uf, -1)

		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return urls
}

func startSearch() {
	var cep string
	var uf string

	fmt.Print("	Digite o Cep: ")

	_, err := fmt.Scan(&cep)

	if err != nil {
		return
	}

	fmt.Print("	Digite o UF: ")

	_, err = fmt.Scan(&uf)

	if err != nil {
		return
	}

	// get each url from urls.txt
	urls := getUrls(cep, uf)

	// repeat the search defined times
	for i := 0; i < times && len(urls) != 0; i++ {
		for i, url := range urls {
			fmt.Println("Site", i+1, ":", url)
			// do a Get Request
			getRequest(url)
			fmt.Println("")
		}

		// wait a few seconds between each search
		if i != times-1 {
			time.Sleep(delay * time.Second)
		}
	}

	fmt.Println("")
}

func getRequest(url string) {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println(url, "foi carregado com sucesso!")
	} else {
		fmt.Println(url, "está com problemas. Status Code:", resp.StatusCode)
	}

	// stores the response in hist.txt
	// storeResponse(url, resp)
}
