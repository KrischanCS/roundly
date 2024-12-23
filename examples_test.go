package htmfunc_test

import (
	"fmt"
	"github.com/ch-schulz/htmfunc/attribute"
	"github.com/ch-schulz/htmfunc/element"
	"github.com/valyala/bytebufferpool"
	"os"
)

//nolint:lll
func ExampleHTML() {
	page := element.Document("html",
		element.HTML(attribute.Lang("en"),
			element.Head(
				element.Title("The Title of the Page"),
			),
			element.Body(nil,
				element.Nav(nil,
					element.A(attribute.HRef("/main"), "Main"),
					element.A(attribute.HRef("/details"), "Details"),
				),
				element.Main(nil,
					element.H1(nil, element.TextTrusted("The Heading")),
					element.Div(nil,
						element.Text("Here could be your content"),
					),
					element.Div(attribute.Class(attribute.JoinValues("escaped", "something")),
						element.Text("HTML characters <div> in </div> here will be escaped &"),
					),
				),
			),
		),
	)

	w := bytebufferpool.Get()

	err := page(w)
	if err != nil {
		fmt.Print("Unexpected error: ", err.Error())
	}

	_, err = w.WriteTo(os.Stdout)
	if err != nil {
		fmt.Print("Unexpected error: ", err.Error())
	}

	//output: <!doctype html><html lang="en"><head><title>The Title of the Page</title></head><body><nav><a href="/main">Main</a><a href="/details">Details</a></nav><main><h1>The Heading</h1><div>Here could be your content</div><div class="escaped something">HTML characters &lt;div&gt; in &lt;/div&gt; here will be escaped &amp;</div></main></body></html>
}
