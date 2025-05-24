package logic

import "github.com/KrischanCS/htmfunc"

func Group(elements ...htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		for _, e := range elements {
			err := e.RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
