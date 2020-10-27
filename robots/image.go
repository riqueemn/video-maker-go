package robots

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	downloadAllImages(&content)
	fmt.Println(content.DownloadedImages)
	fmt.Println(len(content.DownloadedImages))
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

func downloadAllImages(content *entities.Content) {

	//content.Sentences[1].Images[0] = "https://i.guim.co.uk/img/media/a0cc4209a19ad699b2a7324e6644180f9f378b06/0_302_3660_2195/master/3660.jpg?width=700&quality=85&auto=format&fit=max&s=dd16a05d307547f102206a9771a65045"
	for i, sentence := range content.Sentences {
		for k := 0; k < 1; k++ {
			imageURL := sentence.Images[k]
			if existsImage(content.DownloadedImages, imageURL) {
				fmt.Println("Já existe")
			} else {
				content.DownloadedImages = append(content.DownloadedImages, imageURL)
				err := downloadAndSave(imageURL, fmt.Sprintf("%v", i+1)+"-original.png")
				t := fmt.Sprintf("[%v][%v]", i, k)
				fmt.Println(t, imageURL)
				if err != nil {
					fmt.Println("Erro ao baixar a imagem")
				} else {
					fmt.Println("Baixado com sucesso!")
				}
			}

		}
	}
}

func existsImage(imagesURLs []string, url string) bool {
	for _, image := range imagesURLs {
		if url == image {
			return true
		}
	}
	return false
}

func downloadAndSave(URL string, fileName string) error {

	response, err := http.Get(URL)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	file, err := os.Create("images/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return err
}
