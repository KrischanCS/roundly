package htmfunc_test

import (
	"bufio"
	"fmt"
	"github.com/ch-schulz/htmfunc/attr"
	"github.com/ch-schulz/htmfunc/attr/cl"
	"github.com/ch-schulz/htmfunc/el"
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
					el.A(attr.Ls{attr.HRef("/main")}, "Main"),
					el.A(attr.Ls{attr.HRef("/details")}, "Details"),
				),
				el.Main(nil,
					el.H(1, nil, el.TextNoEscape("The Heading")),
					el.Div(nil,
						el.Text("Here could be your content"),
					),
					el.Div(attr.Ls{attr.Class(cl.C("escaped", "something"))},
						el.Text("HTML characters <div> in </div> here will be escaped &"),
					),
				),
			),
		),
	)

	w := bufio.NewWriter(os.Stdout)

	err := page(w)
	if err != nil {
		fmt.Print("Unexpected error: ", err.Error())
	}

	err = w.Flush()
	if err != nil {
		fmt.Print("Unexpected error: ", err.Error())
	}

	//output: <!doctype html><html lang="en"><head><title>The Title of the Page</title></head><body><nav><a href="/main">Main</a><a href="/details">Details</a></nav><main><h1>The Heading</h1><div>Here could be your content</div><div class="escaped something">HTML characters &lt;div&gt; in &lt;/div&gt; here will be escaped &amp;</div></main></body></html>
}
