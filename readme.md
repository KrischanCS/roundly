# Htmfunc

Htmfunc is a go library for dynamic HTML generation. It works without template parsing at runtime or
code generation during development. Instead it provides functions for elements and attributes.
These functions can be composed to reusable components.

```go
package main

import (
	"bufio"
	"os"

	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/text"
)

func main() {
	doc := htmfunc.Document(
		"html",
		Html(nil,
			Head(
				CharSetUtf8(),
				Title(nil, Text("Htmfunc Page")),
			),
			Body(nil,
				H1(Class("the-title"), Text("Hello World!")),
			),
		),
	)

	err := doc.RenderElement(bufio.NewWriter(os.Stdout))
	if err != nil {
		panic(err)
	}
}

```

If you like this, also check out [gomponents](https://maragu.dev/gomponents) which is similar and
out there for a few years longer and probably more stable. If I had found it earlier, I would
probably just have used that.

But when I finally found that gomponents exists, I had already implemented a significant part of
htmfunc and also liked some of the things I did differently, so I decided to continue on it anyway.

## Packages

The rendering functions are distributed over the following packages:

- [element](https://github.com/KrischanCS/htmfunc/element): All HTML elements from the standard.
- [attribute](https://github.com/KrischanCS/htmfunc/attribute): All attributes from the standard.
- [text](https://github.com/KrischanCS/htmfunc/text): Functions for text nodes.
- [logic](https://github.com/KrischanCS/htmfunc/logic): Logic for conditional rendering, loops and
  groups.

I recommend to use those in files without other domain logic and dot-import them as in the
example. Even if this is usually not recommended in go. In my opinion it improves the readability
drastically in this case.

## Usage

To create an element, simple use one of the function from the element package which all returns
an `htmfunc.Element`:

```go
div := Div(nil)

fmt.Println(div.String())

// Output: <div></div>
```

### Rendering

The String function should be used mainly for debugging purposes and for brevity in testing and
examples like the one above.

Usually it will be more efficient to render the element to an `htmfunc.Writer` (an interface
composed of `io.Writer` `io.StringWriter` and `io.ByteWriter`) e.g. `bufio.Writer`. In an
http.HandleFunc for example, the `http.ResponseWriter` can be wrapped in a `bufio.Writer` and passed
to RenderElement. So with our emppty div from before, it would look like this:

```go
func handler(w http.ResponseWriter, r *http.Request) {
buf := bufio.NewWriter(w)
defer buf.Flush()

err := Div(nil).RenderElement(buf)
if err != nil {
// Handle error
}
}
```

> #### Minification
>
> The rendered HTML will not include any line breaks or indentation, so it minified already (At
> least to some extend, minifiers can produce even more compact results).
>
> So the rendered output will look like this:
>
> ```html
> <ul><li>gopher</li><li>crab</li></ul>
> ```
>
> In most of the following examples, the output will be shown with added line breaks and indentation
> for clarity, so instead of the actually produced output above, it will look like this:
>
> ```html
> <ul>
>   <li>gopher</li>
>   <li>crab</li>
> </ul>

### Attributes

The first parameter of each element function, which was `nil` receives the attributes of the
element. So adding an id to our empty div looks like this:

```go
Div(Id("the-empty-div"))

// <div id="the-empty-div"></div>
```

To add multiple attributes, you can use the `attribute.Attributes` function which takes an arbitrary
amount of attributes and concatenates them when rendering:

```go
Div(Attributes(Id("the-empty-div"), Class("fancy-div")))

// <div id="the-empty-div" class="fancy-div">
// </div>
```

Attribute functions also accept multiple string values:

<!--
```go
Div(Class("fancy-div", "glows-in-the-dark"))

// <div class="fancy-div glows-in-the-dark"></div>)
```
-->

## Child Elements

After the attributes, the element functions take an arbitrary amount of child elements:

```go
Body(                // <body>
    Aside(nil,       //   <aside>
        Nav(nil,)    //     <nav></nav>
    ),               //   </aside>
    Main(nil,        //   <main>
        Article(nil, //     <article>
            H1(nil), //       <h1></h1>
            P(nil),  //       <p></p>
            P(nil),  //       <p></p>
        ),           //     </article>
    ),               //   </main>
)                    // </body>
```

## Text

To add text nodes, the functions from the `text` package can be used. They each take a string, which
they will write as as text as content to their parent element.

```go
Article(nil,                                     // <article>
    H1(nil, Text("Hello World!")),               //     <h1>Hello World!</h1>
    P(nil, Text("My first htmfunc paragraph!")), //     <p>This is a paragraph.</p>
    P(nil, Text("And Already the next one…")),   //     <p>And another one…</p>
)                                                // </article>
```

Strings passed to the `Text` function will be escaped, so they can be used safely in HTML:

<!-- The example is double escaped (& -> &amp;) to display it correctly in the rendered markdown -->
```go
Div(nil,                                       // <div>
    Text("<script>alert('Hello!');</script>"), //   &amp;lt;script&amp;gt;alert('Hello!');&amp;lt;/script&amp;gt;
)                                              // </div>
```

If you want to include raw/prerendered HTML, you can use the `RawTrusted` function. It will render
the string as is, without escaping it. This should only be used with trusted content, as it can
lead to XSS vulnerabilities if used with user input or untrusted data.

```go
Div(nil,                                       // <div>
    RawTrusted("<script>alert('Hello!');</script>"), //   <script>alert('Hello!');</script>
)                                              // </div>
```

Another way to include Text, are the convenience variants of the element functions, with the `Text`
suffix. They take a string as the first argument and create an element with that text as its only 
content. Our example from above could be written like this:

```go
Article(nil,                              // <article>
    H1Text("Hello World!"),               //     <h1>Hello World!</h1>
    PText("My first htmfunc paragraph!"), //     <p>This is a paragraph.</p>
    PText("And Already the next one…"),   //     <p>And another one…</p>
)                                         // </article>
```

If you want to add attributes to the convience-text functions, you can add them after the text (
The `Attributes` function is not required here)


```go
Article(nil,                                                                      // <article>
    H1Text("Hello World!", Id("heading")),                                        //     <h1 id="heading">Hello World!</h1>
    PText("My first htmfunc paragraph!", Id("first-paragraph"), Class("opener")), //     <p id="first-paragraph" class="opener">This is a paragraph.</p>
    PText("And Already the next one…"),                                           //     <p>And another one…</p>
)                                                                                 // </article>
```

```go
Ul(Class("mascot-list"),     // <ul class="mascot-list">
    Li(nil, Text("Gopher")), //   <li>Gopher</li>
    Li(nil, Text("Ferris")), //   <li>Ferris</li>
    Li(nil, Text("Lucy")),   //   <li>Lucy</li>
    Li(nil, Text("Duke")),   //   <li>Duke</li>
)                            // </ul>
```

### Logic

The `logic` package provides functions for conditional rendering, loops and groups. They can be
used as element functions, but don't render themself to the writer.

```go
group...
```

```go
if...
```

```go
ifElse...
```

```go
Range...
```

```go
RangeInt...
```

```go
RangeIter...
```
