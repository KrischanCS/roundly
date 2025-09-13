// Copied by generator. DO NOT EDIT.
// Original file: src/logic/group.go

package html

import "github.com/KrischanCS/roundly"

func Group(elements ...roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		if opts == nil {
			return group(w, elements)
		}

		return groupWithOptions(w, elements, opts)
	}
}

func group(w roundly.Writer, elements []roundly.Element) error {
	for _, e := range elements {
		err := e.RenderElement(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func groupWithOptions(
	w roundly.Writer,
	elements []roundly.Element,
	opts *roundly.RenderOptions,
) error {
	for _, e := range elements {
		err := e.RenderElementWithOptions(w, opts)
		if err != nil {
			return err
		}
	}

	return nil
}
