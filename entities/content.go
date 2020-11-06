package entities

//Content -> Conteúdo da pesquisa
type Content struct {
	SearchTerm             string     `json:"searchTerm"`
	PrefixName             string     `json:"prefixName"`
	SourceContentOriginal  string     `json:"sourceContentOriginal"`
	SourceContentSanitized string     `json:"sourceContentSanitized"`
	Sentences              []Sentence `json:"sentences"`
	DownloadedImages       []string   `json:"downloadedImages"`
}
