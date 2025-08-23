package roundly_test

import (
	"fmt"
	"testing"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
	. "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/text"
)

func ExampleRenderOptions_prettyShortContent() {
	div := Div(
		Attributes(
			Id("id"),
			Class("class-1", "class-2"),
		),
		Text("Some Text."),
	)

	w := roundly.NewWriter()

	err := div.RenderElementWithOptions(w, &roundly.RenderOptions{
		Pretty: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output:
	//
	// <div id="id" class="class-1 class-2">
	// 	Some Text.
	// </div>
}

func ExampleRenderOptions_prettyLongContent() {
	div := Div(
		Attributes(
			Id("id"),
			Class("class-1", "class-2"),
		),
		Text(
			"This text is longer. Line breaks should be added to it at the first white-space "+
				"which appears after at least 80 characters on each line. While 80 is the default "+
				"for this 'soft limit', it can be changed in the render options.",
		),
	)

	w := roundly.NewWriter()

	err := div.RenderElementWithOptions(w, &roundly.RenderOptions{
		Pretty: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output:
	//
	// <div id="id" class="class-1 class-2">
	// 	This text is longer. Line breaks should be added to it at the first white-space which
	// 	appears after at least 80 characters on each line. While 80 is the default for this
	// 	&#39;soft limit&#39;, it can be changed in the render options.
	// </div>
}

func ExampleRenderOptions_prettyLongContentWithOptions60Chars() {
	div := Div(
		Attributes(
			Id("id"),
			Class("class-1", "class-2"),
		),
		Text(
			"This text is longer. Line breaks are usually added to it at the first "+
				"white-space which appears after at least 80 characters on each line. In this "+
				"case this 'soft limit' is altered to 60.",
		),
	)

	w := roundly.NewWriter()

	err := div.RenderElementWithOptions(w, &roundly.RenderOptions{
		Pretty:       true,
		LineBreakMin: 60,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output:
	//
	// <div id="id" class="class-1 class-2">
	// 	This text is longer. Line breaks are usually added to it at the
	// 	first white-space which appears after at least 80 characters
	// 	on each line. In this case this &#39;soft limit&#39; is altered to 60.
	// </div>
}

func ExampleRenderOptions_pretty3OrMoreAttributesArePutOnSeparateLines() {
	div := Div(
		Attributes(
			Id("id"),
			Class("class-1", "class-2"),
			ItemScope(),
		),
		Text("Some Text."),
	)

	w := roundly.NewWriter()

	err := div.RenderElementWithOptions(w, &roundly.RenderOptions{
		Pretty: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output:
	//
	// <div
	// 	id="id"
	// 	class="class-1 class-2"
	// 	itemscope
	// >
	// 	Some Text.
	// </div>
}

func ExampleRenderOptions_prettyNestedContent() {
	article := Article(nil,
		Hgroup(nil,
			H1Text("This is a title"),
			PText("And a subtitle"),
		),
		PText(
			"This is a paragraph with some text in it. It should be indented properly. "+
				"Also line breaks should be added if the text is too long.",
		),
		PText(
			"By the way, 'StringPretty()' creates a string with default pretty options. "+
				"But usually you will want to render directly to a writer instead of creating "+
				"an intermediate string for efficiency, but for documentation and testing it's "+
				"pretty handy.",
		),
	)

	fmt.Println(article.StringPretty())

	// Output:
	//
	// <article>
	// 	<hgroup>
	// 		<h1>
	// 			This is a title
	// 		</h1>
	// 		<p>
	// 			And a subtitle
	// 		</p>
	// 	</hgroup>
	// 	<p>
	// 		This is a paragraph with some text in it. It should be indented properly. Also line
	// 		breaks should be added if the text is too long.
	// 	</p>
	// 	<p>
	// 		By the way, &#39;StringPretty()&#39; creates a string with default pretty options. But usually
	// 		you will want to render directly to a writer instead of creating an intermediate
	// 		string for efficiency, but for documentation and testing it&#39;s pretty handy.
	// 	</p>
	// </article>
}

func BenchmarkRenderOptions(b *testing.B) {
	b.ReportAllocs()

	article := Article(
		Attributes(
			Id("id"),
			Class("class-1", "class-2"),
			ItemScope(),
		),
		Hgroup(nil,
			H1Text("This is a title"),
			PText("And a subtitle"),
		),
		PText(
			"This is a paragraph with some text in it. It should be indented properly. "+
				"Also line breaks should be added if the text is too long.",
		),
		PText(
			"By the way, 'StringPretty()' creates a string with default pretty options. "+
				"But usually you will want to render directly to a writer instead of creating "+
				"an intermediate string for efficiency, but for documentation and testing it's "+
				"pretty handy.",
		),
	)

	w := roundly.NewWriter()

	for b.Loop() {
		err := article.RenderElementWithOptions(w, &roundly.RenderOptions{
			Pretty: true,
		})
		if err != nil {
			panic(err)
		}
		w.Reset()
	}
}

func BenchmarkRenderOptions_noOptions(b *testing.B) {
	b.ReportAllocs()

	article := Article(
		Attributes(
			Id("id"),
			Class("class-1", "class-2"),
			ItemScope(),
		),
		Hgroup(nil,
			H1Text("This is a title"),
			PText("And a subtitle"),
		),
		PText(
			"This is a paragraph with some text in it. It should be indented properly. "+
				"Also line breaks should be added if the text is too long.",
		),
		PText(
			"By the way, 'StringPretty()' creates a string with default pretty options. "+
				"But usually you will want to render directly to a writer instead of creating "+
				"an intermediate string for efficiency, but for documentation and testing it's "+
				"pretty handy.",
		),
	)

	w := roundly.NewWriter()

	for b.Loop() {
		err := article.RenderElement(w)
		if err != nil {
			panic(err)
		}
		w.Reset()
	}
}
