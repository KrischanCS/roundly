package html

import (
	"fmt"
)

func ExampleStep() {
	el := Input(Attributes(TypeNumber(), Step(0.5)))
	fmt.Println(el.String())
	// Output: <input type="number" step="0.5">
}

func ExampleValueFloat() {
	el := Progress(Attributes(ValueFloat(0.3), MaxFloat(1)))
	fmt.Println(el.String())
	// Output: <progress value="0.3" max="1"></progress>
}

func ExampleHigh() {
	el := Meter(Attributes(Low(0.2), High(0.8), Optimum(0.6)))
	fmt.Println(el.String())
	// Output: <meter low="0.2" high="0.8" optimum="0.6"></meter>
}

func ExampleLow() {
	el := Meter(Attributes(Low(0.2), High(0.8), Optimum(0.6)))
	fmt.Println(el.String())
	// Output: <meter low="0.2" high="0.8" optimum="0.6"></meter>
}

func ExampleMinFloat() {
	el := Meter(MinFloat(0.1))
	fmt.Println(el.String())
	// Output: <meter min="0.1"></meter>
}

func ExampleOptimum() {
	el := Meter(Attributes(Low(0.2), High(0.8), Optimum(0.6)))
	fmt.Println(el.String())
	// Output: <meter low="0.2" high="0.8" optimum="0.6"></meter>
}
