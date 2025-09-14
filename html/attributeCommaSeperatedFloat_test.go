package html

import (
	"fmt"
)

func ExampleCoords() {
	el := Area(Coords(0, 0.5, 3.25))
	fmt.Println(el.String())
	// Output: <area coords="0,0.5,3.25">
}
