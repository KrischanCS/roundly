package attributes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc/internal/standard"
)

func Test_distinguishLinkDuplicates(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		inputString string
		inputLinks  []standard.Link
		wantString  string
		wantLinks   []standard.Link
	}

	testCases := []testCase{
		{
			name:        "Should not change anything when there are no duplicates",
			inputString: "This is a test [linkA] and [linkB]",
			inputLinks: []standard.Link{
				{Name: "linkA", Url: "https://example.com/link1"},
				{Name: "linkB", Url: "https://example.com/link2"}},
			wantString: "This is a test [linkA] and [linkB]",
			wantLinks: []standard.Link{
				{Name: "linkA", Url: "https://example.com/link1"},
				{Name: "linkB", Url: "https://example.com/link2"},
			},
		},
		{
			name: "Should Adapt the name of the second link when there are duplicates " +
				"with different URLs",
			inputString: "This is a test [linkA] and [linkA]",
			inputLinks: []standard.Link{
				{Name: "linkA", Url: "https://example.com/link1"},
				{Name: "linkA", Url: "https://example.com/link2"},
			},
			wantString: "This is a test [linkA] and [linkA (1)]",
			wantLinks: []standard.Link{
				{Name: "linkA", Url: "https://example.com/link1"},
				{Name: "linkA (1)", Url: "https://example.com/link2"},
			},
		},
		{
			name:        "Should remove the second link when the urls are the same",
			inputString: "This is a test [linkA] and [linkA]",
			inputLinks: []standard.Link{
				{Name: "linkA", Url: "https://example.com/link1"},
				{Name: "linkA", Url: "https://example.com/link1"},
			},
			wantString: "This is a test [linkA] and [linkA]",
			wantLinks: []standard.Link{
				{Name: "linkA", Url: "https://example.com/link1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			gotString, gotLinks := standard.DistinguishLinkDuplicates(tc.inputString, tc.inputLinks)

			// Assert
			assert.Equal(t, tc.wantString, gotString)
			assert.Equal(t, tc.wantLinks, gotLinks)
		})
	}
}

//nolint:funlen
func Test_handleOrderedListTypeAttributes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name         string
		attribute    attribute
		value        string
		funcNameArg  string
		funcNameWant string
		ok           bool
	}

	testCases := []testCase{
		{
			name: "Should not modify the function name when attribute.Name is not 'type' " +
				"and ok must be false",
			attribute:    attribute{Name: "notType"},
			value:        "value",
			funcNameArg:  "FuncName",
			funcNameWant: "FuncName",
			ok:           false,
		},
		{
			name: "Should not modify the function name when attribute.Name is 'type' and " +
				"value is none of [1 i I a A]",
			attribute:    attribute{Name: "type"},
			value:        "value",
			funcNameArg:  "FuncName",
			funcNameWant: "FuncName",
			ok:           false,
		},
		{
			name: "Should not modify the function name when attribute.Name is not 'type' " +
				"and the value is i",
			attribute:    attribute{Name: "notType"},
			value:        "i",
			funcNameArg:  "FuncName",
			funcNameWant: "FuncName",
			ok:           false,
		},
		{
			name: "Should append 'Numeric' to the function name when attribute.Name is " +
				"'type' and value is 1",
			attribute:    attribute{Name: "type"},
			value:        "1",
			funcNameArg:  "FuncName",
			funcNameWant: "FuncNameNumeric",
			ok:           true,
		},
		{
			name: "Should append 'RomanLower' to the function name when attribute.Name " +
				"is 'type' and value is i",
			attribute:    attribute{Name: "type"},
			value:        "i",
			funcNameArg:  "FuncName",
			funcNameWant: "FuncNameRomanLower",
			ok:           true,
		},
		{
			name: "Should append 'Roman' to the function name when attribute.Name " +
				"is 'type' and value is I",
			attribute:    attribute{Name: "type"},
			value:        "I",
			funcNameArg:  "FuncName",
			funcNameWant: "FuncNameRoman",
			ok:           true,
		},
		{
			name: "Should append 'AlphaLower' to the function name when attribute.Name " +
				"is 'type' and value is a",
			attribute:    attribute{Name: "type"},
			value:        "a",
			funcNameArg:  "FuncName",
			funcNameWant: "FuncNameAlphaLower",
			ok:           true,
		},
		{
			name: "Should append 'Alpha' to the function name when attribute.Name is " +
				"'type' and value is A",
			attribute:    attribute{Name: "type"},
			value:        "A",
			funcNameArg:  "FuncName",
			funcNameWant: "FuncNameAlpha",
			ok:           true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Act
			funcNameGot, okGot := handleOrderedListTypeAttributes(
				tc.attribute,
				tc.value,
				tc.funcNameArg,
			)

			// Assert
			assert.Equal(t, tc.funcNameWant, funcNameGot)
			assert.Equal(t, tc.ok, okGot)
		})
	}
}

//nolint:lll
func TestDecomposeEnums(t *testing.T) {
	t.Parallel()

	type test struct {
		name  string
		input []attribute
		want  []attribute
	}

	testCases := []test{
		{
			name: "Should add the value to the funcName and camelcase it when a attribute has a single value",
			input: []attribute{
				{Name: "charset", FuncName: "Charset", ParamName: "charset", Value: "", Values: []string{"utf-8"}},
			},
			want: []attribute{
				{Name: "charset", FuncName: "CharsetUtf8", ParamName: "", Value: "utf-8", Values: []string{"utf-8"}},
			},
		},
		{
			name: "Should create multiple attributes when a attribute has multiple values",
			input: []attribute{
				{Name: "dir", FuncName: "Dir", ParamName: "bdo", Elements: []string{"bdo"}, Value: "", Values: []string{"ltr", "rtl"}},
			},
			want: []attribute{
				{Name: "dir", FuncName: "DirLtr", ParamName: "", Elements: []string{"bdo"}, Value: "ltr", Values: []string{"ltr", "rtl"}},
				{Name: "dir", FuncName: "DirRtl", ParamName: "", Elements: []string{"bdo"}, Value: "rtl", Values: []string{"ltr", "rtl"}},
			},
		},
		{
			name: "Should combine attributes for different elements when they have the same name and values",
			input: []attribute{
				{Name: "dir", FuncName: "Dir", ParamName: "bdo", Elements: []string{"bdo"}, Value: "", Values: []string{"ltr", "rtl"}, Links: []standard.Link{{Name: "linkA", Url: "https://example.com"}}},
				{Name: "dir", FuncName: "Dir", ParamName: "span", Elements: []string{"HTML Elements"}, Value: "", Values: []string{"ltr", "rtl", "auto"}, Links: []standard.Link{{Name: "linkB", Url: "https://example.com"}}},
			},
			want: []attribute{
				{Name: "dir", FuncName: "DirLtr", ParamName: "", Elements: []string{"HTML Elements", "bdo"}, Value: "ltr", Values: []string{"auto", "ltr", "rtl"}, Links: []standard.Link{{Name: "linkA", Url: "https://example.com"}, {Name: "linkB", Url: "https://example.com"}}},
				{Name: "dir", FuncName: "DirRtl", ParamName: "", Elements: []string{"HTML Elements", "bdo"}, Value: "rtl", Values: []string{"auto", "ltr", "rtl"}, Links: []standard.Link{{Name: "linkA", Url: "https://example.com"}, {Name: "linkB", Url: "https://example.com"}}},
				{Name: "dir", FuncName: "DirAuto", ParamName: "", Elements: []string{"HTML Elements"}, Value: "auto", Values: []string{"ltr", "rtl", "auto"}, Links: []standard.Link{{Name: "linkB", Url: "https://example.com"}}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Act
			got := DecomposeEnums(tc.input)

			// Assert
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
