package roundly_test

import (
	"strconv"
	"testing"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/logic"
	. "github.com/KrischanCS/roundly/text"
)

func TestExplainFunc(t *testing.T) {
	w := roundly.NewWriter()

	comp := Article(nil,
		H1Text("The title"),
		PText("Some content"),
		Range([]int{0, 1, 2}, func(_ int, i int) roundly.Element {
			return Div(nil, Text(strconv.Itoa(i)))
		}),
	)

	err := comp.RenderElementWithOptions(w,
		&roundly.RenderOptions{ExplainYourself: true},
	)
	if err != nil {
		t.Fatal(err)
	}
}
