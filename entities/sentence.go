package entities

//Sentence -> Frases com todo conteúdo da pesquisa
type Sentence struct {
	Text     string   `json:"Text"`
	Keywords []string `json:"Keywords"`
	Images   []string `json:"Images"`
}
