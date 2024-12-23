package flow

import (
	"iter"

	"github.com/ch-schulz/htmfunc"
)

func Range[T any](items []T, component func(int, T) htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {

		for i, e := range items {
			err := component(i, e)(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func RangeInt(limit int, component func(int) htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {

		for i := range limit {
			err := component(i)(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func RangeIter(seq iter.Seq2[int, int], component func(int, int) htmfunc.Element) htmfunc.Element {
	if seq == nil {
		return func(w htmfunc.Writer) error {
			return nil
		}
	}

	return func(w htmfunc.Writer) error {

		for t1, t2 := range seq {
			err := component(t1, t2)(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
