package example_test

import (
	"fmt"

	"github.com/yosssi/gohtml"

	hf "github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
	. "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/logic"
	. "github.com/KrischanCS/roundly/text"
)

//nolint:lll,funlen,errcheck
func ExampleDocument() {
	doc := hf.Document(
		"html",
		Html(nil,
			Head(nil,
				Title(nil, Text("Htmf Example")),
				Meta(CharSetUtf8()),
			),
			Body(nil,
				H1(nil, Text("Htmf Example")),
				P(Class("opener"),
					Text("Htmf is a Html generation library for Go, which is based on composable "+
						"functions. It is designed to be easy to use while keeping the safeguards "+
						"of a static type system."),
				),
				P(nil,
					Text("It is an alternative to template engines, which either require some"+
						" parsing at runtime and usually lose type safety or require an extra"+
						" compilation/code-generation step."),
					Em(nil,
						Text("(To be honest, htmf doesn't get rid of the code-generation, "+
							"all the elements and attributes are generated from the HTML"+
							" standard.)"),
					),
					P(nil,
						Text("Each element and attribute exist as a function in the respective "+
							"package. Each element-function takes an attribute-function as it's "+
							"first argument. To put multiple attributes on an element, the Attrs "+
							"function is used, which takes. multiple attributes as arguments. "+
							"After the attributes parameter, the element function takes multiple "+
							"child elements as arguments. That way, elements-functions can be "+
							"nested just like HTML elements."),
					),
					P(nil,
						Text("Text must be wrapped in one of the functions Text or RawTrusted "+
							"to be included as a child element. In the normal Text function, "+
							"the given string is escaped. "+
							"When you input some html here it will be rendered like this: '<em"+
							">Some content</em>'. "),
						Br(nil),
						RawTrusted("The same content in RawTrusted will be rendered as '<em"+
							">Some content</em>."),
					),
					P(Attributes(Class("info"), Id("if-else")),
						IfElse(true,
							Text("In addition to elements, attributes and text, htmf provides "+
								"the flow-package, which contains functions to renderer content "+
								"conditionally and/or repeatedly."),
							Text("This will not be rendered unless the condition is changed to"+
								" true."),
						),
					),
				),
			),
		),
	)

	w := hf.NewWriter()
	_ = doc.RenderElement(w)

	fmt.Println(gohtml.Format(w.String()))

	// Output:
	// <!doctype html>
	// <html>
	//   <head>
	//     <title>
	//       Htmf Example
	//     </title>
	//     <meta charset="utf-8">
	//   </head>
	//   <body>
	//     <h1>
	//       Htmf Example
	//     </h1>
	//     <p class="opener">
	//       Htmf is a Html generation library for Go, which is based on composable functions. It is designed to be easy to use while keeping the safeguards of a static type system.
	//     </p>
	//     <p>
	//       It is an alternative to template engines, which either require some parsing at runtime and usually lose type safety or require an extra compilation/code-generation step.
	//       <em>
	//         (To be honest, htmf doesn&#39;t get rid of the code-generation, all the elements and attributes are generated from the HTML standard.)
	//       </em>
	//       <p>
	//         Each element and attribute exist as a function in the respective package. Each element-function takes an attribute-function as it&#39;s first argument. To put multiple attributes on an element, the Attrs function is used, which takes. multiple attributes as arguments. After the attributes parameter, the element function takes multiple child elements as arguments. That way, elements-functions can be nested just like HTML elements.
	//       </p>
	//       <p>
	//         Text must be wrapped in one of the functions Text or RawTrusted to be included as a child element. In the normal Text function, the given string is escaped. When you input some html here it will be rendered like this: &#39;&lt;em&gt;Some content&lt;/em&gt;&#39;.
	//         <br>
	//         The same content in RawTrusted will be rendered as '
	//         <em>
	//           Some content
	//         </em>
	//         .
	//       </p>
	//       <p class="info" id="if-else">
	//         In addition to elements, attributes and text, htmf provides the flow-package, which contains functions to renderer content conditionally and/or repeatedly.
	//       </p>
	//     </p>
	//   </body>
	// </html>
}
