package mousecontrol

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func simulateMouseMovement(x, y int) {
	robotgo.MoveMouse(x, y)
	fmt.Printf("Simulating mouse movement: x=%d, y=%d\n", x, y)
}

