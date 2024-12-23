package element

import (
	"github.com/ch-schulz/htmfunc"
)

type Iterator[T any] func() func() (element T, ok bool)

func IteratorOf[T any](elements ...T) Iterator[T] {
	return func() func() (element T, ok bool) {
		tail := elements

		return func() (element T, ok bool) {
			if len(tail) == 0 {
				return element, false
			}

			element, tail = tail[0], tail[1:]

			return element, true
		}
	}
}

func IteratorFromTo(from, to int) Iterator[int] {
	return func() func() (element int, ok bool) {
		i := from

		return func() (element int, ok bool) {
			i++

			if i > to {
				return 0, false
			}

			return i - 1, true
		}
	}
}

func Range[T any](iterator Iterator[T], component func(T) htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		data := iterator()

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
