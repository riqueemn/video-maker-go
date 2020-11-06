package robots

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/riqueemn/video-maker-go/entities"
	"gopkg.in/gographics/imagick.v3/imagick"
)

//Video -> struct robô de image
type Video struct{}

type templateSetting struct {
	width     uint
	height    uint
	gravity   imagick.GravityType
	lineWidth int
}

//RobotProcess -> Sequência de processos do Robô
func (i *Video) RobotProcess() {
	content := robotState.Load()

	convertAllImages(content)
	createAllSetenceImages(content)
	createYoutubeThumbnail()

	createAfterEffectsScript(content)
	//robotState.Save(content)

	renderVideoWithAfterEffects()
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
	inputFile := "images/input/" + fmt.Sprintf("[%v][1]", sentenceIndex) + "-original.png[0]"
	outputFile := "images/output/" + fmt.Sprintf("%v", sentenceIndex-1) + "-converted.png"
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
		inputFile = "images/input/" + fmt.Sprintf("[%v][2]", sentenceIndex) + "-original.png[0]"
		err = img.ReadImage(inputFile)
		if err != nil {
			return err
		}
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
	outputFile := "images/output/" + fmt.Sprintf("%v", sentenceIndex) + "-sentence.png"

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

func createAfterEffectsScript(content entities.Content) {
	robotState.SaveScript(content)
}

func renderVideoWithAfterEffects() {
	aeRenderFilePath := ""
	templateFilePath := ""
	outputFilePath := ""

	fmt.Println("> Starting After Effects")

	cmd := exec.Command(aeRenderFilePath, "-comp", "main", "-project", templateFilePath, "-output", outputFilePath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	n, err := cmd.Stdout.Write([]byte("data"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("in all caps: %q\n", out.String())
	fmt.Println("> After Effects closed", n)

}
