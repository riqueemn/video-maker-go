package entities

//Sentence -> Frases com todo conteúdo da pesquisa
type Sentence struct {
	Text              string   `json:"text"`
	Keywords          []string `json:"keywords"`
	Images            []string `json:"images"`
	GoogleSearchQuery string   `json:"googleSearchQuery"`
}
