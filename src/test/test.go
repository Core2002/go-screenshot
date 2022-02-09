package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	awa, _ := robotgo.GetScaleSize()
	qwq, _ := robotgo.GetScreenSize()
	owo := float32(awa) / float32(qwq)
	fmt.Println(owo)
	// robotgo.MoveClick(1323, 703, `left`, false)
}
