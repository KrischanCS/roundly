package html

import (
	"fmt"
)

func ExampleStart() {
	el := Ol(Start(3))
	fmt.Println(el.String())
	// Output: <ol start="3"></ol>
}

func ExampleTabIndex() {
	el := Div(TabIndex(5))
	fmt.Println(el.String())
	// Output: <div tabindex="5"></div>
}

func ExampleValueInt() {
	el := Li(ValueInt(3), Text("Three"))
	fmt.Println(el.String())
	// Output: <li value="3">Three</li>
}
