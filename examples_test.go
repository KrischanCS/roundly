package htmfunc_test

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ch-schulz/htmfunc/attribute"
	"github.com/ch-schulz/htmfunc/element"
)

//nolint:lll
func ExampleHtml() {
	page := element.Document("html",
		element.Html(attribute.Lang("en"),
			element.Head(
				element.Title("The Title of the Page"),
			),
			element.Body(nil,
				element.Nav(nil,
					element.A(attribute.HRef_AArea("/main"), element.TextTrusted("Main")),
					element.A(attribute.HRef_AArea("/details"), element.TextTrusted("Details")),
				),
				element.Main(nil,
					element.H1(nil, element.TextTrusted("The Heading")),
					element.Div(nil,
						element.Text("Here could be your content"),
					),
					element.Div(attribute.Class("escaped", "something"),
						element.Text("HTML characters <div> in </div> here will be escaped &"),
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
