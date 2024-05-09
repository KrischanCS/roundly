package el

import (
	"github.com/ch-schulz/htmfunc"
)

type Iterator[T any] func() (element T, ok bool)

func IteratorOf[T any](elements ...T) Iterator[T] {
	return func() (element T, ok bool) {
		if len(elements) == 0 {
			return element, false
		}

		element, elements = elements[0], elements[1:]
		return element, true
	}
}

func IteratorFromTo(from, to int) Iterator[int] {
	i := from
	return func() (element int, ok bool) {
		if i == to {
			return i, false
		}
		i++
		return i - 1, true
	}
}

func Range[T any](data Iterator[T], component func(T) htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		for d, ok := data(); ok; d, ok = data() {
			err := component(d)(w)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func If(condition bool, component htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		if !condition {
			return nil
		}

		return component(w)
	}
}
