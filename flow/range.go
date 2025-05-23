package flow

import (
	"iter"

	"github.com/KrischanCS/htmfunc"
)

// Range creates an htmfunc.Element for each item if items.
func Range[E any](items []E, component func(int, E) htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		for i, e := range items {
			err := component(i, e).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// RangeInt creates an htmfunc.Element for each int in [0, limit).
func RangeInt(limit int, component func(int) htmfunc.Element) htmfunc.Element {
	return func(w htmfunc.Writer) error {
		for i := range limit {
			err := component(i).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// RangeIter creates an htmfunc.Element yielded by seq.
func RangeIter[E any](seq iter.Seq[E], component func(E) htmfunc.Element) htmfunc.Element {
	if seq == nil {
		return func(_ htmfunc.Writer) error {
			return nil
		}
	}

	return func(w htmfunc.Writer) error {
		for t := range seq {
			err := component(t).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
