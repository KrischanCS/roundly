package html

import (
	"fmt"

	"github.com/KrischanCS/roundly"
)

func ExampleDelText() {
	del := DelText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := del(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <del id="id" class="class-1 class-2">Some Text.</del>
}

func ExampleInsText() {
	ins := InsText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := ins(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <ins id="id" class="class-1 class-2">Some Text.</ins>
}

func ExampleButtonText() {
	button := ButtonText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := button(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <button id="id" class="class-1 class-2">Some Text.</button>
}

func ExampleLegendText() {
	legend := LegendText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := legend(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <legend id="id" class="class-1 class-2">Some Text.</legend>
}

func ExampleOptionText() {
	option := OptionText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := option(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <option id="id" class="class-1 class-2">Some Text.</option>
}

func ExampleOutputText() {
	output := OutputText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := output(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <output id="id" class="class-1 class-2">Some Text.</output>
}

func ExampleProgressText() {
	progress := ProgressText("Some Text.", Id("id"),
		Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := progress(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <progress id="id" class="class-1 class-2">Some Text.</progress>
}

func ExampleTextareaText() {
	textarea := TextareaText("Some Text.", Id("id"),
		Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := textarea(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <textarea id="id" class="class-1 class-2">Some Text.</textarea>
}

func ExampleLabelText() {
	label := LabelText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := label(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <label id="id" class="class-1 class-2">Some Text.</label>
}

func ExampleBlockquoteText() {
	blockquote := BlockquoteText("Some Text.", Id("id"),
		Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := blockquote(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <blockquote id="id" class="class-1 class-2">Some Text.</blockquote>
}

func ExampleDdText() {
	dd := DdText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := dd(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <dd id="id" class="class-1 class-2">Some Text.</dd>
}

func ExampleDtText() {
	dt := DtText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := dt(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <dt id="id" class="class-1 class-2">Some Text.</dt>
}

func ExampleFigcaptionText() {
	figcaption := FigcaptionText("Some Text.", Id("id"),
		Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := figcaption(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <figcaption id="id" class="class-1 class-2">Some Text.</figcaption>
}

func ExampleLiText() {
	li := LiText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := li(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <li id="id" class="class-1 class-2">Some Text.</li>
}

func ExamplePText() {
	p := PText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := p(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output:  <p id="id" class="class-1 class-2">Some Text.</p>
}

func ExamplePreText() {
	pre := PreText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := pre(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <pre id="id" class="class-1 class-2">Some Text.</pre>
}

func ExampleUlText() {
	ul := UlText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := ul(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <ul id="id" class="class-1 class-2">Some Text.</ul>
}

func ExampleSummaryText() {
	summary := SummaryText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := summary(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <summary id="id" class="class-1 class-2">Some Text.</summary>
}

func ExampleNoscriptText() {
	noscript := NoscriptText("Some Text.", Id("id"),
		Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := noscript(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <noscript id="id" class="class-1 class-2">Some Text.</noscript>
}

func ExampleH1Text() {
	h1 := H1Text("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := h1(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <h1 id="id" class="class-1 class-2">Some Text.</h1>
}

func ExampleH2Text() {
	h2 := H2Text("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := h2(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <h2 id="id" class="class-1 class-2">Some Text.</h2>
}

func ExampleH3Text() {
	h3 := H3Text("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := h3(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <h3 id="id" class="class-1 class-2">Some Text.</h3>
}

func ExampleH4Text() {
	h4 := H4Text("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := h4(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <h4 id="id" class="class-1 class-2">Some Text.</h4>
}

func ExampleH5Text() {
	h5 := H5Text("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := h5(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <h5 id="id" class="class-1 class-2">Some Text.</h5>
}

func ExampleH6Text() {
	h6 := H6Text("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := h6(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <h6 id="id" class="class-1 class-2">Some Text.</h6>
}

func ExampleTitleText() {
	title := TitleText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := title(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <title id="id" class="class-1 class-2">Some Text.</title>
}

func ExampleCaptionText() {
	caption := CaptionText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := caption(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <caption id="id" class="class-1 class-2">Some Text.</caption>
}

func ExampleTdText() {
	td := TdText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := td(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <td id="id" class="class-1 class-2">Some Text.</td>
}

func ExampleThText() {
	th := ThText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := th(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <th id="id" class="class-1 class-2">Some Text.</th>
}

func ExampleAText() {
	a := AText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := a(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output:  <a id="id" class="class-1 class-2">Some Text.</a>
}

func ExampleAbbrText() {
	abbr := AbbrText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := abbr(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <abbr id="id" class="class-1 class-2">Some Text.</abbr>
}

func ExampleBText() {
	b := BText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := b(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <b id="id" class="class-1 class-2">Some Text.</b>
}

func ExampleCiteText() {
	cite := CiteText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := cite(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <cite id="id" class="class-1 class-2">Some Text.</cite>
}

func ExampleCodeText() {
	code := CodeText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := code(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <code id="id" class="class-1 class-2">Some Text.</code>
}

func ExampleDataText() {
	data := DataText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := data(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <data id="id" class="class-1 class-2">Some Text.</data>
}

func ExampleDfnText() {
	dfn := DfnText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := dfn(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <dfn id="id" class="class-1 class-2">Some Text.</dfn>
}

func ExampleEmText() {
	em := EmText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := em(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <em id="id" class="class-1 class-2">Some Text.</em>
}

func ExampleIText() {
	i := IText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := i(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <i id="id" class="class-1 class-2">Some Text.</i>
}

func ExampleKbdText() {
	kbd := KbdText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := kbd(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <kbd id="id" class="class-1 class-2">Some Text.</kbd>
}

func ExampleMarkText() {
	mark := MarkText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := mark(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <mark id="id" class="class-1 class-2">Some Text.</mark>
}

func ExampleQText() {
	q := QText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := q(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <q id="id" class="class-1 class-2">Some Text.</q>
}

func ExampleRpText() {
	rp := RpText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := rp(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <rp id="id" class="class-1 class-2">Some Text.</rp>
}

func ExampleRtText() {
	rt := RtText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := rt(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <rt id="id" class="class-1 class-2">Some Text.</rt>
}

func ExampleSText() {
	s := SText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := s(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <s id="id" class="class-1 class-2">Some Text.</s>
}

func ExampleSampText() {
	samp := SampText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := samp(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <samp id="id" class="class-1 class-2">Some Text.</samp>
}

func ExampleSmallText() {
	small := SmallText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := small(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <small id="id" class="class-1 class-2">Some Text.</small>
}

func ExampleSpanText() {
	span := SpanText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := span(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <span id="id" class="class-1 class-2">Some Text.</span>
}

func ExampleStrongText() {
	strong := StrongText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := strong(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <strong id="id" class="class-1 class-2">Some Text.</strong>
}

func ExampleSubText() {
	sub := SubText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := sub(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <sub id="id" class="class-1 class-2">Some Text.</sub>
}

func ExampleSupText() {
	sup := SupText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := sup(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <sup id="id" class="class-1 class-2">Some Text.</sup>
}

func ExampleUText() {
	u := UText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := u(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <u id="id" class="class-1 class-2">Some Text.</u>
}

func ExampleVarText() {
	varE := VarText("Some Text.", Id("id"), Class("class-1", "class-2"))

	w := roundly.NewWriter()

	err := varE(w, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output: <var id="id" class="class-1 class-2">Some Text.</var>
}
