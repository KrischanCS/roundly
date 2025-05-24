package hfcomp

import (
	"fmt"

	"github.com/KrischanCS/htmfunc"
)

func ExampleBdiAuto() {
	element := BdiAuto("Text.")

	w := htmfunc.NewWriter()

	err := element.RenderElement(w)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <bdi dir="auto">Text.</bdi>
}

func ExampleBdiLtr() {
	element := BdiLtr("Text.")

	w := htmfunc.NewWriter()

	err := element.RenderElement(w)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <bdi dir="ltr">Text.</bdi>
}

func ExampleBdiRtl() {
	element := BdiRtl("مرحباً بالعالم")

	w := htmfunc.NewWriter()

	err := element.RenderElement(w)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <bdi dir="rtl">مرحباً بالعالم</bdi>
}

func ExampleBdoLtr() {
	element := BdoLtr("Text.")

	w := htmfunc.NewWriter()

	err := element.RenderElement(w)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <bdo dir="ltr">Text.</bdo>
}

func ExampleBdoRtl() {
	element := BdoRtl("مرحباً بالعالم")

	w := htmfunc.NewWriter()

	err := element.RenderElement(w)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <bdo dir="rtl">مرحباً بالعالم</bdo>
}
