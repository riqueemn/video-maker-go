package robots

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/riqueemn/video-maker-go/entities"
)

//State -> struct do estado da estrutura de dados
type State struct{}

//Save -> Salva o Content
func (s *State) Save(content entities.Content) {
	b, jsonErr := json.Marshal(content)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	err := ioutil.WriteFile("C:/Users/Henrique/go/src/github.com/riqueemn/video-maker-go/state.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//Load -> Carrega o Content
func (s *State) Load() entities.Content {
	var contentJSON entities.Content
	file, err := ioutil.ReadFile("C:/Users/Henrique/go/src/github.com/riqueemn/video-maker-go/state.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &contentJSON)

	return contentJSON
}
