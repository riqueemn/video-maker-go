package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Content struct {
	searchTerm string
	//prefixNumber     int
	prefixName string
}

func askAndReturnSearchTerm() string {
	fmt.Print("\nType a Wikipedia shearch Term: ")
	//fmt.Println("Digite um termo para pesquisa do Wikipedia:")

	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')

	//fmt.Println(text)
	var text string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text = scanner.Text()
		return text
	}

	return ""
}

func askAndReturnPrefix() string {
	prefixes := []string{"Who is", "What is", "The history of"}
	//const prefixes = ["Quem é", "O que é", "A história de"]

	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')
	//log.Println(text)
	//n, _ := strconv.Atoi(text)
	//log.Println(n)
	//n, _ := strconv.Atoi(text)

	for i, prefixe := range prefixes {
		fmt.Print("\n[", i+1, "] ", prefixe)
	}
	fmt.Print("\n\n")

	var text string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text = scanner.Text()
		n, _ := strconv.Atoi(text)
		n--
		if n != 1 && n != 2 && n != 3 {
			fmt.Println("Valor indevido!!!")
			os.Exit(1)
		}

		return prefixes[n-1]
	}

	return ""
}

func main() {
	var content Content

	content.searchTerm = askAndReturnSearchTerm()
	content.prefixName = askAndReturnPrefix()

	fmt.Print("\n\n", content)

}
