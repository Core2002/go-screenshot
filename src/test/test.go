package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println(robotgo.GetScaleSize())
	robotgo.MoveClick(1323, 703, `left`, false)
}
