package html

import (
	"fmt"
)

func ExampleAccessKey() {
	el := Div(AccessKey("ctrl", "b"))
	fmt.Println(el.String())
	// Output: <div accesskey="ctrl b"></div>
}

func ExampleBlocking() {
	el := Link(Blocking("render"))
	fmt.Println(el.String())
	// Output: <link blocking="render">
}

func ExampleClass() {
	el := Div(Class("a", "b"))
	fmt.Println(el.String())
	// Output: <div class="a b"></div>
}

func ExampleForStrings() {
	el := Output(ForStrings("a", "b"))
	fmt.Println(el.String())
	// Output: <output for="a b"></output>
}

func ExampleHeaders() {
	el := Td(Headers("h1", "h2", "h3"))
	fmt.Println(el.String())
	// Output: <td headers="h1 h2 h3"></td>
}

func ExampleItemProp() {
	el := Div(ItemProp("name", "jobTitle"))
	fmt.Println(el.String())
	// Output: <div itemprop="name jobTitle"></div>
}

func ExampleItemRef() {
	el := Div(ItemRef("a", "b", "c"))
	fmt.Println(el.String())
	// Output: <div itemref="a b c"></div>
}

func ExampleItemType() {
	el := Div(ItemType("http://schema.org/Person", "http://schema.org/Thing"))
	fmt.Println(el.String())
	// Output: <div itemtype="http://schema.org/Person http://schema.org/Thing"></div>
}

func ExamplePing() {
	el := A(Attributes(HRef("/page"), Ping("/track1", "/track2")))
	fmt.Println(el.String())
	// Output: <a href="/page" ping="/track1 /track2"></a>
}

func ExampleRel() {
	el := Link(Attributes(Rel("preload", "modulepreload"), HRef("/app.js")))
	fmt.Println(el.String())
	// Output: <link rel="preload modulepreload" href="/app.js">
}

func ExampleSandBox() {
	el := Iframe(SandBox("allow-scripts", "allow-same-origin"))
	fmt.Println(el.String())
	// Output: <iframe sandbox="allow-scripts allow-same-origin"></iframe>
}

func ExampleSizesStrings() {
	el := Link(SizesStrings("32x32", "any"))
	fmt.Println(el.String())
	// Output: <link sizes="32x32 any">
}
