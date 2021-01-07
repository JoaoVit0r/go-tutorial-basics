// package name
package main

// imported packages
import "fmt"

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
			// startSearch()
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
