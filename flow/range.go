package flow

import (
	"iter"

	"github.com/ch-schulz/htmfunc"
)

func Range[T any](items []T, component func(int, T) htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteFunc(func(w htmfunc.Writer) error {
		for i, e := range items {
			err := component(i, e).RenderHTML(w)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func RangeInt(limit int, component func(int) htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteFunc(func(w htmfunc.Writer) error {
		for range limit {
			err := component(limit).RenderHTML(w)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func RangeIter(seq iter.Seq2[int, int], component func(int, int) htmfunc.Element) htmfunc.Element {
	if seq == nil {
		return htmfunc.WriteFunc(func(_ htmfunc.Writer) error {
			return nil
		})
	}

	return htmfunc.WriteFunc(func(w htmfunc.Writer) error {
		for t1, t2 := range seq {
			err := component(t1, t2).RenderHTML(w)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
