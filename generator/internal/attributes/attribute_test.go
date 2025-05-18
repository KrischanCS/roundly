package attributes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/htmfunc/generator/internal/standard"
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
			name:        "Should Adapt the name of the second link when there are duplicates with different URLs",
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
