package robots

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/riqueemn/video-maker-go/entities"
)

var (
	robotState State
)

func init() {
	file, err := ioutil.ReadFile("")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &secrets)

}

//Input -> struct do Input dos dados
type Input struct{}

//RobotProcess -> Sequência de processos do Robô
func (i *Input) RobotProcess() {
	var content entities.Content

	content.SearchTerm = askAndReturnSearchTerm()
	content.PrefixName = askAndReturnPrefix()

	robotState.Save(content)

}

//askAndReturnSearchTerm -> Retorna o termo de pesquisa
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

//askAndReturnPrefix -> Retorna o prefixo de pesquisa
func askAndReturnPrefix() string {
	prefixes := []string{"Who is", "What is", "The history of", "Cancel"}
	//const prefixes = ["Quem é", "O que é", "A história de", "Cancelar"]

	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')
	//log.Println(text)
	//n, _ := strconv.Atoi(text)
	//log.Println(n)
	//n, _ := strconv.Atoi(text)

	for i, prefixe := range prefixes {
		if i != len(prefixes)-1 {
			fmt.Print("\n[", i+1, "] ", prefixe)

			continue
		}
		fmt.Print("\n[", 0, "] ", prefixe)

	}
	fmt.Print("\n\n")

	var text string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text = scanner.Text()
		if text != "1" && text != "2" && text != "3" && text != "0" {
			fmt.Println("Valor indevido!!!")
			os.Exit(1)
		} else {
			if text == "0" {
				os.Exit(1)
			}

			n, _ := strconv.Atoi(text)

			return prefixes[n-1]
		}
	}

	return ""
}
