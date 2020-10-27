package entities

//Content -> Conteúdo da pesquisa
type Content struct {
	SearchTerm             string     `json:"SearchTerm"`
	PrefixName             string     `json:"PrefixName"`
	SourceContentOriginal  string     `json:"SourceContentOriginal"`
	SourceContentSanitized string     `json:"SourceContentSanitized"`
	Sentences              []Sentence `json:"Sentences"`
	DownloadedImages       []string   `json:"DownloadedImages"`
}
