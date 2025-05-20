package example_test

import (
	"fmt"

	hf "github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/flow"
	. "github.com/KrischanCS/htmfunc/text"
)

//nolint:lll
func ExampleDocument() {
	doc := hf.Document(
		"html",
		Html(nil,
			Head(nil,
				Title(nil, Text("Hello World")),
				Meta(CharSetUtf8()),
			),
			Body(nil,
				H1(nil, Text("Hello World")),
				P(Class("a-paragraph"),
					Text("This is a simple HTML document."),
				),
				IfElse(true,
					P(nil, Text("This will be rendered")),
					P(nil, Text("This will not be rendered")),
				),
			),
		),
	)

	w := hf.NewWriter(256)

	err := doc.RenderElement(w)
	if err != nil {
		panic("unexpected error: " + err.Error())
	}

	fmt.Println(w.String())

	// TODO: Should be like this later: Output:
	// <!doctype html>
	// <html>
	// 	<head>
	// 		<title>Hello World</title>
	// 		<meta charset="utf-8">
	// 	</head>
	// 	<body>
	// 		<h1>Hello World</h1>
	// 		<p class="a-paragraph">
	// 			This is a simple HTML document.
	// 		</p>
	// 		<p>
	// 			This will be rendered
	// 		</p>
	// 	</body>
	// </html>

	// Output: <!doctype html><html><head><title>Hello World</title><meta charset="utf-8"></head><body><h1>Hello World</h1><p class="a-paragraph">This is a simple HTML document.</p><p>This will be rendered</p></body></html>
}
