package htmfunc_test

import (
	"fmt"
	"github.com/ch-schulz/htmfunc/attr"
	"github.com/ch-schulz/htmfunc/el"
	"github.com/valyala/bytebufferpool"
	"os"
)

//nolint:lll
func ExampleHTML() {
	page := el.Document("html",
		el.HTML(attr.Lang("en"),
			el.Head(
				el.Title("The Title of the Page"),
			),
			el.Body(nil,
				el.Nav(nil,
					el.A(attr.HRef("/main"), "Main"),
					el.A(attr.HRef("/details"), "Details"),
				),
				el.Main(nil,
					el.H1(nil, el.TextTrusted("The Heading")),
					el.Div(nil,
						el.Text("Here could be your content"),
					),
					el.Div(attr.Class(attr.JoinValues("escaped", "something")),
						el.Text("HTML characters <div> in </div> here will be escaped &"),
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
