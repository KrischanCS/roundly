package htmfunc_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/KrischanCS/htmfunc"
	attr "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/text"
)

//nolint:nolintlint,lll
func ExampleHtml() {
	page := Document("html",
		Html(attr.Lang("en"),
			Head(
				nil,
				Title(nil, Text("The Title of the Page")),
			),
			Body(nil,
				Nav(nil,
					A(attr.HRef_AArea("/main"), TextTrusted("Main")),
					A(attr.HRef_AArea("/details"), TextTrusted("Details")),
				),
				Main(nil,
					H1(nil, TextTrusted("The Heading")),
					Div(nil,
						Text("Here could be your content"),
					),
					Div(attr.Class("escaped", "something"),
						Text("HTML characters <div> in </div> here will be escaped &"),
					),
				),
			),
		),
	)

	w := bufio.NewWriter(os.Stdout)

	err := page.RenderElement(w)
	if err != nil {
		fmt.Print("Unexpected error: ", err.Error())
	}

	err = w.Flush()
	if err != nil {
		fmt.Print("Unexpected error: ", err.Error())
	}

	//output: <!doctype html><html lang="en"><head><title>The Title of the Page</title></head><body><nav><a href="/main">Main</a><a href="/details">Details</a></nav><main><h1>The Heading</h1><div>Here could be your content</div><div class="escaped something">HTML characters &lt;div&gt; in &lt;/div&gt; here will be escaped &amp;</div></main></body></html>
}
