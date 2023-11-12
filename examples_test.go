package htmfunc_test

import (
	"bufio"
	"fmt"
	"github.com/ch-schulz/htmfunc/attr"
	"github.com/ch-schulz/htmfunc/el"
	"os"
)

//nolint:lll
func ExampleHTML() {
	page := el.HTML(
		attr.Lang("en"),
		el.Head(
			el.Title("The Title of the Page"),
		),
		el.Body(nil,
			el.Nav(nil,
				el.A(attr.Ls{attr.HRef("/main")}, "Main"),
				el.A(attr.Ls{attr.HRef("/details")}, "Details"),
			),
			el.Main(nil,
				el.H(1, nil,
					el.Div(nil, el.Text("Here could be your content")),
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

	//output: <html lang="en"><head><title>The Title of the Page</title></head><body><nav><a href="/main">Main</a><a href="/details">Details</a></nav><main><h1><div>Here could be your content</div></h1></main></body></html>
}
