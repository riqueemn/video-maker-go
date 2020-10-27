package main

import (
	"fmt"

	"github.com/riqueemn/video-maker-go/entities"
	"github.com/riqueemn/video-maker-go/robots"
)

const ()

var (
	//content    entities.Content
	robotState robots.State
	robotInput robots.Input
	robotText  robots.Text
)

func main() {
	robotInput.RobotProcess()
	robotText.RobotProcess()

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
