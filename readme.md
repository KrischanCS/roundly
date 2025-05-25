# Htmfunc

Htmfunc is a go library for dynamic HTML generation. It works without template parsing at runtime or
code generation during development. Instead it provides functions for elements and attributes.
These functions can be composed to reusable components.

```go
Table(                                                  // <table
    Attributes(
        Id("mascot-table"),                             //   id="mascot-table"
        Class("fancy-table", "mascots"),                //   class="fancy-table mascots"
    ),                                                  // >
    Thead(nil,                                          //   <thead>
        Tr(nil,                                         //     <tr>
            ThText("Language"),                         //       <th>Language</th>  
            ThText("Kind"),                             //       <th>Kind</th>
            ThText("Name"),                             //       <th>Name</th>
        ),                                              //     </tr>
    ),                                                  //   </thead>
    Tbody(nil,                                          //   <tbody>
        Range([]Mascot{
            {"Go", "Blue Gopher", "The Go Gopher"},     //     <tr><th scope="row">Go</th><td>Blue Gopher</td><td>The Go Gopher</td></tr>
            {"Rust", "Orange Crab", "Ferris"},          //     <tr><th scope="row">Rust</th><td>Orange Crab</td><td>Ferris</td></tr>
            {"Gleam", "Pink Starfish", "Lucy"},         //     <tr><th scope="row">Gleam</th><td>Pink Starfish</td><td>Lucy</td></tr>
            {"Zig", "Ziguana", "Ziggy & Zero"},         //     <tr><th scope="row">Zig</th><td>Ziguana</td><td>Ziggy & Zero</td></tr>
            {"Java", "?", "Duke"},                      //     <tr><th scope="row">Java</th><td>?</td><td>Duke</td></tr>
        }, func(i int, mascot Mascot) htmfunc.Element {
            return Tr(nil,
                Th(ScopeRow(), Text(mascot.Language)),
                TdText(mascot.Kind),
                TdText(mascot.Name),
            )
        }),                                               
    ),                                                  //   </tbody>
)                                                       // </table>
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

// <div></div>
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

Attributes have different parameter types, depending on what the standard allows:

```go
// Text attributes take a single string as value:
Id("an-id") // id="an-id"

// Text list attributes takes an arbitrary amount of strings
// and renders them, according to the standard, either comma or space separated:
Class("class1")                        // class="class1"
Class("class1", "class2", )            // class="class1 class2"
Accept("text/xml", "application/json") // accept="text/xml, application/json"

// Equivalently to the text attributes, there are int, uint, float and float-list attributes
TableIndex(1)              // tableindex="1"
ColSpan(2)                 // colspan="2"
Value(6.28)                // value="6.28"
Coords(52.44897, 13.18063) // coords="52.44897, 13.18063"

// Bool attributes take no value, they just render their name if present:
Disabled() // disabled

// Attributes wich have a fixed set of allowed values, exist once for each Value, with the value as
// suffix:
DirAuto() // dir="auto"
DirLtr()  // dir="ltr"
DirRtl()  // dir="rtl"
```



### Child Elements

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
Article(nil,
    H1Text("Hello World!", Id("heading")),
    PText("My first htmfunc paragraph!", Id("first-paragraph"), Class("opener")),
    PText("And Already the next one…"),
)

// <article>
//     <h1 id="heading">Hello World!</h1>
//     <p id="first-paragraph" class="opener">This is a paragraph.</p>
//     <p>And another one…</p>
// </article>
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

The `logic` package provides meta-elements, which are not rendered itself, but they influence how
their children are rendered. Specifically there are elements for grouping, conditional rendering and
repeated rendering.

#### Groups

A group simple renders all its children without a parent. Use this if you wnat to create a component
which renders multiple elements without a common parent.

```go
Group(
    H1(nil, Text("Title")),                     // <h1>Title</h1>
    P(nil, Text("This is a paragraph.")),       // <p>This is a paragraph.</p>
    P(nil, Text("This is another paragraph.")), // <p>This is another paragraph.</p>
)
```

#### Conditionals

Conditional renderings represent classical if/if-else statements.

The `If` function takes a boolean and an element. If the boolean is true, the element is rendered,
otherwise not.

```go
Group(
    If(2>1, H1(nil, Text("This is rendered"))),     // <h1>This is rendered</h1>
    If(1>2, H1(nil, Text("This is not rendered"))), //
)
```

IfElse takes a boolean and two elements. If the boolean is true, the first element is rendered,
otherwise the second element is rendered.

```go
Group(
    IfElse(2>1,                                //
        H1(nil, Text("This is rendered")),     // <h1>This is rendered</h1>
        H1(nil, Text("This is not rendered")), //
    ),                                         //
    IfElse(1>2,                                //
        H1(nil, Text("This is not rendered")), //
        H1(nil, Text("This is rendered")),     // <h1>This is rendered</h1>
    ),
)

```

#### Repetitions

For repeating elements with different data, htmfunc provides the `Range` functions:

- `Range` takes a slice of elements and passes each elements index and value to the provided 
    function.
- `RangeInt` takes an integer and will call the provided function for each integer from 0 to the
    given integer - 1.
- `RangeSeq` takes an iterator (iter.Seq) and applies each yielded value to the provided
    function.


```go
Range([]string{"gopher", "ferris", "lucy", "duke"}, 
    func(i int, mascot string) Element {
        return Li(nil, Text(fmt.Sprintf("$d – %s", i, mascot)))
    },
)

// <li>0 – gopher</li>
// <li>1 – ferris</li>
// <li>2 – lucy</li>
// <li>3 – duke</li>
```

```go
RangeInt(4, 
    func(i int) Element {
        return Li(nil, Text(fmt.Sprintf("Element %d", i)))
    },
)

// <li>Element 0</li>
// <li>Element 1</li>
// <li>Element 2</li>
// <li>Element 3</li>
```

```go
RangeSeq(slices.Reverse([]string{"gopher", "ferris", "lucy", "duke"}), )
    func(mascot string) Element {
        return Li(nil, Text(mascot))
    },
)

// <li>duke</li>
// <li>lucy</li>
// <li>ferris</li>
// <li>gopher</li>
```

## Naming of Elements and Attributes

All HTML elements have a counterpart element function, which has the same name but with an uppercase
first letter, e. g.:

- `div` -> `Div`
- `span` -> `Span`
- `input` -> `Input`

For attributes the renaming is slightly more complex:

- Attributes are converted to PascalCase, so `colspan` becomes `ColSpan`.
- Dashes are removed from the name, so `accept-charset` becomes `AcceptCharset`.
- If there is a name collision between an attribute and an element, the attribute is suffixed with
  `Attritbute`. Currently these attributes are affected by this:
  - `abbr`  -> `AbbrAttribute`
  - `cite`  -> `CiteAttribute`
  - `data`  -> `DataAttribute`
  - `form`  -> `FormAttribute`
  - `label` -> `LabelAttribute`
  - `slot`  -> `SlotAttribute`
  - `span`  -> `SpanAttribute`
  - `style` -> `StyleAttribute`
  - `title` -> `TitleAttribute`
- There are a few name collisions between attributes of different parameter types, in those cases
    the variants are suffixes with `Strings`, `Int`, `Uint`, `Float`, `Floats` or `True` (text attributes
    are considered default, thus get no suffix). Currently these attributes are affected by this:
  - `for`    -> `ForStrings`
  - `sizes`  -> `SizesStrings`
  - `max`    -> `MaxFloat`
  - `min`    -> `MinFloat`
  - `value`  -> `ValueFloat`
  - `value`  -> `ValueInt`
- Finally, the type of an unordered list can have the values `1`, `a`, `A`, `i` or `I`. When adding 
    those to as uppercased suffixes, there are name collisions between a and A and i and I. Those 
    are completely renamed:
  - `type="1"` -> `TypeNumeric`
  - `type="a"` -> `TypeAlphaLower`
  - `type="A"` -> `TypeAlpha`
  - `type="i"` -> `TypeRomanLower`
  - `type="I"` -> `TypeRoman`