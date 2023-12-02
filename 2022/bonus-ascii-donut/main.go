package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println()
	fmt.Println("Begin again")
	fmt.Println()

	// for true {
	for z := 0; z < 1; z++ {
		frameChars := []string{}
		for y := float64(0); y < 25; y++ {
			for x := float64(0); x < 100; x++ {
				// frameChars = append(frameChars, sampleCheckerboard(x, y))
				// map x to -1...1 range
				remappedX := float64(x/100*2 - 1)
				// correct for aspect ratio range of y
				// valuable lesson about floats, calling float64 after the math is done does no good.
				// remappedY := float64((y/float64(25))*float64(2)-1.0) * float64(2*25/100)
				remappedY := float64((y/25*2 - 1.0) * .5)
				spew.Dump(float64(2*25) / 100)
				// spew.Dump("remy", sample((remappedX), (remappedY)))
				frameChars = append(frameChars, sample((remappedX), (remappedY)))
				// spew.Dump(remappedX, remappedY)

			}
			frameChars = append(frameChars, "\n")
		}

		// spew.Dump(frameChars)
		// fmt.Print("\033[2J", strings.Join(frameChars, ""))
		fmt.Print(strings.Join(frameChars, ""))

		// cap at 30 fps
		time.Sleep(time.Second / 30)
	}

}

func circle(x, y float64) float64 {
	// the range of x is -1...1 the circle's radius will be
	// 40% meaning the diameter is 40% of the screen
	radius := float64(0.4)
	// the distance from the center of the screen and subtract
	// the radius so d is < 0 inside the circle 0 on the edge, and
	// > 0 on the outside of the circle
	area := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	return area - radius
}

func sample(x, y float64) string {
	if circle(x, y) <= 0 {
		return "#"
	}
	return " "
}

// func sampleCheckerboard(x, y int) string {
// 	secondsSince1970 := time.Now().Unix()
// 	// draw an alternating checkerboard pattern
// 	if (x+y+int(secondsSince1970))%2 == 0 {
// 		return "#"
// 	}
// 	return " "
// }
