package logic

import "github.com/KrischanCS/roundly"

func Group(elements ...roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts ...*roundly.RenderOptions) error {
		if len(opts) != 0 {
			return renderGroupWithOptions(w, elements, opts[0])
		}

		for _, e := range elements {
			err := e.RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func renderGroupWithOptions(
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
