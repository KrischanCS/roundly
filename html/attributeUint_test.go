package html

import (
	"fmt"
)

func ExampleCols() {
	el := Textarea(Cols(40))
	fmt.Println(el.String())
	// Output: <textarea cols="40"></textarea>
}

func ExampleColSpan() {
	el := Th(ColSpan(2))
	fmt.Println(el.String())
	// Output: <th colspan="2"></th>
}

func ExampleHeight() {
	el := Img(Height(200))
	fmt.Println(el.String())
	// Output: <img height="200">
}

func ExampleMaxLength() {
	el := Input(Attributes(TypeText(), MaxLength(10)))
	fmt.Println(el.String())
	// Output: <input type="text" maxlength="10">
}

func ExampleMinLength() {
	el := Input(Attributes(TypeText(), MinLength(2)))
	fmt.Println(el.String())
	// Output: <input type="text" minlength="2">
}

func ExampleRows() {
	el := Textarea(Rows(5))
	fmt.Println(el.String())
	// Output: <textarea rows="5"></textarea>
}

func ExampleRowSpan() {
	el := Td(RowSpan(3))
	fmt.Println(el.String())
	// Output: <td rowspan="3"></td>
}

func ExampleSize() {
	el := Input(Attributes(TypeText(), Size(30)))
	fmt.Println(el.String())
	// Output: <input type="text" size="30">
}

func ExampleSpanAttribute_col() {
	el := Col(SpanAttribute(4))
	fmt.Println(el.String())
	// Output: <col span="4">
}

func ExampleSpanAttribute_colgroup() {
	el := Colgroup(SpanAttribute(2), Col(nil))
	fmt.Println(el.String())
	// Output: <colgroup span="2"><col></colgroup>
}

func ExampleWidth() {
	el := Canvas(Width(640))
	fmt.Println(el.String())
	// Output: <canvas width="640"></canvas>
}
