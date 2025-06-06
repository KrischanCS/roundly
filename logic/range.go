package logic

import (
	"iter"

	"github.com/KrischanCS/roundly"
)

// Range creates an roundly.Element for each item if items.
func Range[E any](items []E, component func(int, E) roundly.Element) roundly.Element {
	return func(w roundly.Writer) error {
		for i, e := range items {
			err := component(i, e).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// RangeInt creates an roundly.Element for each int in [0, limit).
func RangeInt(limit int, component func(int) roundly.Element) roundly.Element {
	return func(w roundly.Writer) error {
		for i := range limit {
			err := component(i).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// RangeSeq creates an roundly.Element for each value yielded by seq.
func RangeSeq[E any](seq iter.Seq[E], component func(E) roundly.Element) roundly.Element {
	if seq == nil {
		return func(_ roundly.Writer) error {
			return nil
		}
	}

	return func(w roundly.Writer) error {
		for t := range seq {
			err := component(t).RenderElement(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
