package entities

//Content -> Conteúdo da pesquisa
type Content struct {
	SearchTerm             string
	PrefixName             string
	SourceContentOriginal  string
	SourceContentSanitized string
	Sentences              []Sentence
}
