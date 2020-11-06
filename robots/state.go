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
	err := ioutil.WriteFile(secrets.Dir, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//SaveScript -> Salva o Content com o formato Script utilizado pelo After Effects
func (s *State) SaveScript(content entities.Content) {
	b, jsonErr := json.Marshal(content)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	script := append([]byte("var content = "), b...)

	err := ioutil.WriteFile(secrets.ScriptFilePath, script, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//Load -> Carrega o Content
func (s *State) Load() entities.Content {
	var contentJSON entities.Content
	file, err := ioutil.ReadFile(secrets.Dir)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &contentJSON)

	return contentJSON
}
