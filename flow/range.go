package flow

import (
	"iter"

	"github.com/ch-schulz/htmfunc"
)

func Range[T any](items []T, component func(int, T) htmfunc.ElementRenderer) htmfunc.ElementRenderer {
	return htmfunc.WriteElementFunc(func(w htmfunc.Writer) error {
		for i, e := range items {
			err := component(i, e).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func RangeInt(limit int, component func(int) htmfunc.ElementRenderer) htmfunc.ElementRenderer {
	return htmfunc.WriteElementFunc(func(w htmfunc.Writer) error {
		for i := range limit {
			err := component(i).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func RangeIter(seq iter.Seq2[int, int], component func(int, int) htmfunc.ElementRenderer) htmfunc.ElementRenderer {
	if seq == nil {
		return htmfunc.WriteElementFunc(func(_ htmfunc.Writer) error {
			return nil
		})
	}

	return htmfunc.WriteElementFunc(func(w htmfunc.Writer) error {
		for t1, t2 := range seq {
			err := component(t1, t2).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
