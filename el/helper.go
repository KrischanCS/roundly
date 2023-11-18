package el

import (
	"github.com/ch-schulz/htmfunc"
)

func Range[T any](component func(T) htmfunc.Element, data ...T) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		for _, d := range data {
			err := component(d)(w)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func If(component htmfunc.Element, ok bool) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		if !ok {
			return nil
		}

		return component(w)
	}
}
