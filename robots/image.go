package robots

import (
	"log"
	"net/http"

	"github.com/riqueemn/video-maker-go/entities"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

var (
	apiKeyGoogleCloud = secrets.APIKeyGoogleCloud
	cx                = secrets.APIKeyGoogleSearch
)

const (
	//query = "Michael Jackson"
	num = 2
)

//Image -> struct robô de image
type Image struct{}

//RobotProcess -> Sequência de processos do Robô
func (i *Image) RobotProcess() {
	content := robotState.Load()

	apiKeyGoogleCloud = secrets.APIKeyGoogleCloud
	cx = secrets.APIKeyGoogleSearch

	fetchImagesOfAllSentences(&content)

	robotState.Save(content)
}

func fetchGoogleAndReturnImagesLinks(query string) []string {
	client := &http.Client{Transport: &transport.APIKey{Key: apiKeyGoogleCloud}}

	svc, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := svc.Cse.List().Cx(cx).Q(query).SearchType("image").Num(2).Do()
	if err != nil {
		log.Fatal(err)
	}

	var imageURLs []string
	for _, result := range resp.Items {

		imageURLs = append(imageURLs, result.Link)
	}

	return imageURLs
}

func fetchImagesOfAllSentences(content *entities.Content) {
	for i, sentence := range content.Sentences {
		query := content.SearchTerm + " " + sentence.Keywords[0]
		content.Sentences[i].GoogleSearchQuery = query

		content.Sentences[i].Images = fetchGoogleAndReturnImagesLinks(query)

	}
}
