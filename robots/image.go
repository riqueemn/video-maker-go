package robots

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/riqueemn/video-maker-go/entities"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"

	"gopkg.in/gographics/imagick.v3/imagick"
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

type templateSetting struct {
	width     uint
	height    uint
	gravity   imagick.GravityType
	lineWidth int
}

//RobotProcess -> Sequência de processos do Robô
func (i *Image) RobotProcess() {
	content := robotState.Load()

	apiKeyGoogleCloud = secrets.APIKeyGoogleCloud
	cx = secrets.APIKeyGoogleSearch

	/*fetchImagesOfAllSentences(&content)
	downloadAllImages(&content)
	fmt.Println(content.DownloadedImages)
	fmt.Println(len(content.DownloadedImages))
	*/
	//convertAllImages(content)
	//createAllSetenceImages(content)
	createYoutubeThumbnail()
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

	file, err := os.Create("images/input/" + fileName)
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

func convertAllImages(content entities.Content) {
	for i := 1; i <= len(content.Sentences); i++ {
		err := convertImages(i)
		var text string
		if err != nil {
			text = fmt.Sprintf("[%v] Image converted failed", i)
			fmt.Println(text)

			log.Fatal(err)
		} else {
			text = fmt.Sprintf("[%v] Image converted sucess", i)
			fmt.Println(text)
		}
	}
}

func convertImages(sentenceIndex int) error {
	inputFile := "images/input/" + fmt.Sprintf("%v", sentenceIndex) + "-original.png[0]"
	outputFile := "images/output/" + fmt.Sprintf("%v", sentenceIndex) + "-converted.png"
	width := uint(1920)
	height := uint(1080)

	imagick.Initialize()
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()
	cw := imagick.NewPixelWand()

	cw.SetColor("white")

	err = mw.NewImage(width, height, cw)
	if err != nil {
		return err
	}

	err = mw.ResizeImage(width, height, imagick.FILTER_LANCZOS)
	if err != nil {
		return err
	}

	img := imagick.NewMagickWand()

	err = img.ReadImage(inputFile)
	if err != nil {
		return err
	}

	imgClone := img.Clone()

	err = img.BlurImage(50, 10)
	if err != nil {
		return err
	}

	w := img.GetImageWidth()
	h := img.GetImageHeight()

	newH := uint((h * width) / w)

	err = img.ResizeImage(width, newH, imagick.FILTER_LANCZOS)
	if err != nil {
		return err
	}

	err = mw.CompositeImageGravity(img, imagick.COMPOSITE_OP_OVER, imagick.GRAVITY_CENTER)
	if err != nil {
		return err
	}

	w = imgClone.GetImageWidth()
	h = imgClone.GetImageHeight()

	newW := uint((w * height) / h)

	err = imgClone.ResizeImage(newW, height, imagick.FILTER_LANCZOS)
	if err != nil {
		return err
	}

	err = mw.CompositeImageGravity(imgClone, imagick.COMPOSITE_OP_OVER, imagick.GRAVITY_CENTER)
	if err != nil {
		return err
	}

	mw.WriteImage(outputFile)

	return err
}

func createAllSetenceImages(content entities.Content) {
	for i := 0; i < len(content.Sentences); i++ {
		createSentenceImage(i, content.Sentences[i].Text)
		text := fmt.Sprintf("[%v] Create Sentence Image sucess", i+1)
		fmt.Println(text)
	}
}

func createSentenceImage(sentenceIndex int, sententeText string) {
	outputFile := "images/output/" + fmt.Sprintf("%v", sentenceIndex+1) + "-sentence.png"

	templateSettings := []templateSetting{
		{
			width:     1920,
			height:    400,
			gravity:   imagick.GRAVITY_CENTER,
			lineWidth: 50,
		},
		{
			width:     1920,
			height:    1080,
			gravity:   imagick.GRAVITY_CENTER,
			lineWidth: 50,
		},
		{
			width:     800,
			height:    1080,
			gravity:   imagick.GRAVITY_WEST,
			lineWidth: 28,
		},
		{
			width:     1920,
			height:    400,
			gravity:   imagick.GRAVITY_CENTER,
			lineWidth: 50,
		},
		{
			width:     1920,
			height:    1080,
			gravity:   imagick.GRAVITY_CENTER,
			lineWidth: 50,
		},
		{
			width:     800,
			height:    1080,
			gravity:   imagick.GRAVITY_WEST,
			lineWidth: 28,
		},
		{
			width:     1920,
			height:    400,
			gravity:   imagick.GRAVITY_CENTER,
			lineWidth: 50,
		},
	}

	text := breakLine(sententeText, templateSettings[sentenceIndex].lineWidth)

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()
	pw := imagick.NewPixelWand()

	pw.SetColor("transparent")

	mw.NewImage(templateSettings[sentenceIndex].width, templateSettings[sentenceIndex].height, pw)
	//mw.SetSize(templateSettings[sentenceIndex].width, templateSettings[sentenceIndex].height)

	pw.SetColor("white")
	dw.SetGravity(templateSettings[sentenceIndex].gravity)
	dw.SetFontResolution(float64(templateSettings[sentenceIndex].width), float64(templateSettings[sentenceIndex].height))
	dw.SetTextKerning(1)

	dw.SetFillColor(pw)
	dw.SetFont("Arial")
	dw.SetFontSize(50)

	dw.Annotation(0, 0, text)

	mw.DrawImage(dw)

	mw.WriteImage(outputFile)
}

func breakLine(text string, lineWidth int) string {
	t := strings.Split(text, " ")

	var k int
	for i, x := range t {
		k += len(x)

		if k > lineWidth {
			x = "\n " + x
			t[i] = x
			k = len(x)
		}
	}

	return strings.Join(t, " ")
}

func createYoutubeThumbnail() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()

	mw.ReadImage("images/output/1-converted.png")

	mw.WriteImage("images/output/thumbnail.jpg")

	fmt.Println("Create Thumbnail")
}
