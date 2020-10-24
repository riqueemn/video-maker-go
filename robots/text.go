package robots

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/algorithmiaio/algorithmia-go"
	"github.com/riqueemn/video-maker-go/entities"
	"gopkg.in/neurosnap/sentences.v1/english"
)

var ()

//Text -> struct do robô de texto
type Text struct {
}

//RobotProcess -> Sequência de processos do Robô
func (t *Text) RobotProcess(content *entities.Content) {

	fetchContentFromWikipedia(content)
	sanitizeContent(content)
	breakContentIntoSentences(content)
	//fmt.Println(content)

}

func myFunc(waitGroup *sync.WaitGroup) {
	time.Sleep(10 * time.Second)

	waitGroup.Done()
}

func fetchContentFromWikipedia(content *entities.Content) {

	var client = algorithmia.NewClient("", "")

	algo, _ := client.Algo("web/WikipediaParser/0.1.2?timeout=300")
	resp, _ := algo.Pipe(content.SearchTerm)
	response, _ := resp.(*algorithmia.AlgoResponse)

	wikiPediaContent := response.Result.(map[string]interface{})

	content.SourceContentOriginal = fmt.Sprintf("%v", wikiPediaContent["content"])

}

func sanitizeContent(content *entities.Content) {
	withoutBlankLines := removeBlankLines(content.SourceContentOriginal)
	withoutMarkdown := removeMarkdown(withoutBlankLines)
	//withoutDatesInParenteses := removeDatesInParenteses(withoutMarkdown)
	//fmt.Println(withoutMarkdown)
	content.SourceContentSanitized = withoutMarkdown
	//fmt.Println(len(withoutMarkdown))

}

func removeBlankLines(texto string) []string {
	allLines := strings.Split(texto, "\n")

	var withoutBlankLines []string
	for _, line := range allLines {
		if line != "" {
			withoutBlankLines = append(withoutBlankLines, line)
		}
	}

	return withoutBlankLines
}

func removeMarkdown(withoutBlankLines []string) string {
	var withoutMarkdown []string
	for _, line := range withoutBlankLines {
		if line[0] != '=' {
			withoutMarkdown = append(withoutMarkdown, line)
		}
	}

	return strings.Join(withoutMarkdown, " ")
}

/*
func removeDatesInParenteses(withoutMarkdown []string) []string {

	var withoutDatesInParenteses []string

	for _, line := range withoutMarkdown {

		newLine := strings.Replace(line, \(\d\d\d\d-\d\d\d\d)/, " ", -1)

		withoutDatesInParenteses = append(withoutDatesInParenteses, newLine)
	}

	return withoutDatesInParenteses
}
*/

func breakContentIntoSentences(content *entities.Content) {
	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	sentences := tokenizer.Tokenize(content.SourceContentSanitized)
	content.Sentences = make([]entities.Sentence, len(sentences))
	for i, s := range sentences {
		content.Sentences[i].Text = s.Text
		content.Sentences[i].Keywords = nil
		content.Sentences[i].Images = nil
	}
}

func print(text []string) {
	for _, line := range text {

		t := fmt.Sprintf(";;%v;;", line)
		fmt.Println(t)
	}
}
