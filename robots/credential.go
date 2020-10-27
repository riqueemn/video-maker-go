package robots

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	secrets Credential
)

// Credential -> struct das apis
type Credential struct {
	APIKeyAlgorithmia  string `json:"apiKeyAlgorithmia"`
	APIKeyWatson       string `json:"apiKeyWatson"`
	Dir                string `json:"dir"`
	APIKeyGoogleCloud  string `json:"apiKeyGoogleCloud"`
	APIKeyGoogleSearch string `json:"apiKeyGoogleSearch"`
}

//RobotProcess -> Sequência de processos do Robô
func (c *Credential) RobotProcess() {
	loadCredential()
}

func loadCredential() {
	file, err := ioutil.ReadFile("secrets.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &secrets)
}
