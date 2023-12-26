package el

import (
	"github.com/ch-schulz/htmfunc"
)

type Iterator[T any] func() (element T, hasNext bool)

func IteratorOf[T any](elements ...T) Iterator[T] {
	if len(elements) == 0 {
		return func() (element T, hasNext bool) {
			return
		}
	}

	return func() (element T, hasNext bool) {
		element, elements = elements[0], elements[1:]
		return element, len(elements) != 0
	}
}

func Range[T any](data Iterator[T], component func(T) htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		for d, hasNext := data(); hasNext; d, hasNext = data() {
			err := component(d)(w)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func If(ok bool, component htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		if !ok {
			return nil
		}

		return component(w)
	}
}
