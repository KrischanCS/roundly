package logic

import "github.com/KrischanCS/roundly"

func Group(elements ...roundly.Element) roundly.Element {
	return func(w roundly.Writer) error {
		for _, e := range elements {
			err := e.RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
