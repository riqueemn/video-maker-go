package main

import (
	"fmt"

	"github.com/riqueemn/video-maker-go/entities"
	"github.com/riqueemn/video-maker-go/robots"
)

const ()

var (
	//content    entities.Content
	robotCredential robots.Credential
	robotState      robots.State
	robotInput      robots.Input
	robotText       robots.Text
	robotImage      robots.Image
)

func main() {
	robotCredential.RobotProcess()
	robotInput.RobotProcess()
	robotText.RobotProcess()
	robotImage.RobotProcess()

	content := robotState.Load()

	fmt.Print("\n")
	print(content)

}

func print(content entities.Content) {
	//fmt.Println(content)
	for _, line := range content.Sentences {

		fmt.Println(line)
	}
}
