package main

import (
	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
	. "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/logic"
	. "github.com/KrischanCS/roundly/text"
)

func Page(title string, mainContent roundly.Element) roundly.Element {
	return roundly.Document(
		"html",
		Html(Lang("en"),
			Head(nil,
				TitleText("roundly – "+title),
				Meta(CharSetUtf8()),
				Meta(Attributes(Name("viewport"), Content("width=device-width, initial-scale=1"))),
				Link(Attributes(Rel("stylesheet"), HRef("/style.css"))),
			),
			Body(nil,
				Header(nil,
					//Hgroup(Class("page-title"),
					//	H1Text("roundly"),
					//	PText("html. pure go"),
					//	A(Attributes(
					//		HRef("/"),
					//		roundly.WriteAttribute("aria-label", "Home")),
					//	),
					//),
					Img(
						Attributes(
							Src("/logo.svg"),
							Alt("roundly logo"),
							Height(120),
						),
					),
					Nav(nil,
						A(
							Attributes(
								HRef("/home"),
								If(title == "Home", Class("selected"))),
							Text("Home()"),
						),
						A(
							Attributes(
								HRef("/examples"),
								If(title == "Examples", Class("selected"))),
							Text("Examples()"),
						),
					),
				),
				Main(nil,
					mainContent,
				),
			),
		),
	)
}

func HomePage() roundly.Element {
	return Page("Home",
		Article(nil,
			H2Text("Summary"),
			PText(
				"Roundly is a library for creating composable HTML components in pure Go. Instead "+
					"of writing template files, which needs to be wired to go in some way, roundly "+
					"allows you to write your components as go functions.",
			),
			PText(
				"So there is no need for any runtime parsing of templates or code generation while "+
					"development. Instead you get the type safety you know from go",
			),
			P(nil,
				Text("Here is how roundly components looks like:"),
				Br(nil),
				Pre(nil, Text(sampleCode)),
			),
		),
	)
}

func ExamplesPage() roundly.Element {
	return Page("Examples",
		Article(nil,
			H2Text("Under Construction"),
		),
	)
}

func SideBySide(lrs []LeftRight) roundly.Element {
	return Div(Class("side-by-side"),
		Range(lrs, func(_ int, lr LeftRight) roundly.Element {
			return Group(lr.left, lr.right)
		}),
	)
}

type LeftRight struct {
	left, right roundly.Element
}

func TableExample() roundly.Element {
	return SideBySide(
		[]LeftRight{
			{
				PreText("Table("),
				PreText("<table>"),
			},
			{
				PreText("\tAttributes("),
				RawTrusted(""),
			},
			{
				PreText("\t\tId(\"mascot-table\")"),
				PreText("\t\tid=\"mascot-table\""),
			},
		})
}

const sampleCode = `Table(
    Attributes(
        Id("mascot-table"),
        Class("fancy-table", "mascots"),
    ),
    Thead(nil,
        Tr(nil,
            ThText("Language"),
            ThText("Kind"),
            ThText("Name"),
        ),
    ),
    Tbody(nil,
        Range([]Mascot{
            {"Go", "Blue Gopher", "The Go Gopher"},
            {"Rust", "Orange Crab", "Ferris"},
            {"Gleam", "Pink Starfish", "Lucy"},
            {"Zig", "Ziguana", "Ziggy & Zero"},
            {"Java", "?", "Duke"},
        }, func(i int, mascot Mascot) roundly.Element {
            return Tr(nil,
                Th(ScopeRow(), Text(mascot.Language)),
                TdText(mascot.Kind),
                TdText(mascot.Name),
            )
        }),
    ),
)
`
