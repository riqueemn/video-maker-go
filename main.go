package main

import (
	"fmt"

	"github.com/riqueemn/video-maker-go/robots"
)

const ()

var (
	robotCredential robots.Credential
	robotState      robots.State
	robotInput      robots.Input
	robotText       robots.Text
	robotImage      robots.Image
	robotVideo      robots.Video
)

func main() {
	robotCredential.RobotProcess()
	robotInput.RobotProcess()
	robotText.RobotProcess()
	robotImage.RobotProcess()
	robotVideo.RobotProcess()

	//print()

}

func print() {
	content := robotState.Load()

	fmt.Print("\n")
	//fmt.Println(content)
	for _, line := range content.Sentences {

		fmt.Println(line)
	}
}
