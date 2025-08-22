package logic

import (
	"iter"

	"github.com/KrischanCS/roundly"
)

// Range creates a roundly.Element for each item if items.
func Range[E any](items []E, component func(int, E) roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		if opts == nil {
			return ranger(w, items, component)
		}

		return rangeOpts(w, items, component, opts)
	}
}

func ranger[E any](w roundly.Writer, items []E, component func(int, E) roundly.Element) error {
	for i, e := range items {
		err := component(i, e).RenderElement(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func rangeOpts[E any](
	w roundly.Writer,
	items []E,
	component func(int, E) roundly.Element,
	opt *roundly.RenderOptions,
) error {
	for i, e := range items {
		err := component(i, e).RenderElementWithOptions(w, opt)
		if err != nil {
			return err
		}
	}

	return nil
}

// RangeInt creates a roundly.Element for each int in [0, limit).
func RangeInt(limit int, component func(int) roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		if opts == nil {
			return rangeInt(w, limit, component)
		}

		return rangeIntOpts(w, limit, component, opts)
	}
}

func rangeInt(w roundly.Writer, limit int, component func(int) roundly.Element) error {
	for i := range limit {
		err := component(i).RenderElement(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func rangeIntOpts(
	w roundly.Writer,
	limit int,
	component func(int) roundly.Element,
	opts *roundly.RenderOptions,
) error {
	for i := range limit {
		err := component(i).RenderElementWithOptions(w, opts)
		if err != nil {
			return err
		}
	}

	return nil
}

// RangeSeq creates a roundly.Element for each value yielded by seq.
func RangeSeq[E any](seq iter.Seq[E], component func(E) roundly.Element) roundly.Element {
	if seq == nil {
		return func(_ roundly.Writer, _ *roundly.RenderOptions) error {
			return nil
		}
	}

	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		if opts == nil {
			return rangeSeq(w, seq, component)
		}

		return rangeSeqOpts(w, seq, component, opts)
	}
}

func rangeSeq[E any](w roundly.Writer, seq iter.Seq[E], component func(E) roundly.Element) error {
	for t := range seq {
		err := component(t).RenderElement(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func rangeSeqOpts[E any](
	w roundly.Writer,
	seq iter.Seq[E],
	component func(E) roundly.Element,
	opts *roundly.RenderOptions,
) error {
	for t := range seq {
		err := component(t).RenderElementWithOptions(w, opts)
		if err != nil {
			return err
		}
	}

	return nil
}
