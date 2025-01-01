package fuzz

import (
	"fmt"
	"math"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ch-schulz/htmfunc"
	. "github.com/ch-schulz/htmfunc/element"
)

type elementFunc func(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element

// // Currently added as reminders, not used in Fuzzing
// var (
//	document                     = Document
//	html                         = Html
//	base                         = Base
//	doctype                      = Doctype
//	head                         = Head
//	meta                         = Meta
//	style                        = Style
//	styleTrusted                 = StyleTrusted
//	title                        = Title
//	titleTrusted                 = TitleTrusted
//	input                        = Input
//	data                         = Data
//	timeMachineReadableAsContent = TimeMachineReadableAsContent
//	timeAttribute                = TimeAttribute
//	bdo                          = Bdo
//	br                           = Br
//)

var elements = []elementFunc{
	Canvas,
	Ins,
	Del,
	Area,
	Audio,
	Embed,
	Iframe,
	Img,
	Map,
	Math,
	Object,
	Picture,
	Source,
	Svg,
	Track,
	Video,
	Area,
	Blockquote,
	Dd,
	Div,
	Dl,
	Dt,
	Figcaption,
	Figure,
	Hr,
	Li,
	Main,
	Menu,
	Ol,
	P,
	Pre,
	Search,
	Ul,
	Form,
	Label,
	Button,
	Select,
	Datalist,
	Optgroup,
	Option,
	Textarea,
	Output,
	Progress,
	Meter,
	Fieldset,
	Legend,
	Details,
	Summary,
	Dialog,
	Noscript,
	Script,
	Slot,
	Template,
	Body,
	Article,
	Section,
	Nav,
	Aside,
	H1,
	H2,
	H3,
	H4,
	H5,
	H6,
	Hgroup,
	Header,
	Footer,
	Address,
	Table,
	Caption,
	Colgroup,
	Col,
	Tbody,
	Thead,
	Tfoot,
	Tr,
	Td,
	Th,
	A,
	Em,
	Strong,
	Small,
	S,
	Cite,
	Q,
	Dfn,
	Abbr,
	Ruby,
	Rt,
	Rp,
	Code,
	Var,
	Samp,
	Kbd,
	Sub,
	Sup,
	I,
	B,
	U,
	Mark,
	Bdi,
	Span,
	Wbr,
}

var texts = []string{
	"Hello, World!",
	"Hi",
	"Three < Four",
	"'Text with quotes'",
	"<>>><<''\"\"&&",
	"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt.",
}

func FuzzDom(f *testing.F) {
	f.Add(uint64(0), uint64(23))
	f.Add(uint64(math.MaxUint64), uint64(255))
	f.Add(uint64(127), uint64(47899872162134))
	f.Add(uint64(1337), uint64(234148))
	f.Add(uint64(42), uint64(986345))

	w := htmfunc.NewWriter(4096)

	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r := rand.New(rand.NewPCG(seed1, seed2)) //nolint:gosec

		depth := 0

		tree := createTree(r, depth)

		err := tree.RenderElement(w)

		assert.NoError(t, err)

		w.Reset()
	})
}

func createTree(random *rand.Rand, depth int) htmfunc.Element {
	if depth > 20 || random.Float64() < 0.05 {
		fmt.Println("text, d:", depth)
		return Div(nil, Text(randomElement(random, texts)))
	}

	depth++

	children := make([]htmfunc.Element, 0, 3)
	for random.Float64() < 0.55 && len(children) < 20 {
		children = append(children, createTree(random, depth))
	}

	if len(children) == 0 {
		return Div(nil, Text(randomElement(random, texts)))
	}

	element := randomElement(random, elements)

	return element(nil, children...)
}

func randomElement[T any](r *rand.Rand, list []T) T {
	return list[r.IntN(len(list))]
}
