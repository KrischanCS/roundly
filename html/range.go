// Copied by generator. DO NOT EDIT.
// Original file: src/logic/range.go

package html

import (
	"fmt"
	"iter"

	"github.com/KrischanCS/roundly"
)

type Iterable[E any] interface {
	~[]E | iter.Seq[E] | iter.Seq2[int, E]
}

// Range creates a roundly.Element for each item if items.
func Range[E any, I Iterable[E]](items I, component func(int, E) roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		if items == nil {
			return nil
		}

		switch items := any(items).(type) {
		case []E:
			return rangeSlice(items, component)(w, opts)
		case iter.Seq2[int, E]:
			return rangeSeq2(items, component)(w, opts)
		case iter.Seq[E]:
			return rangeSeq(items, component)(w, opts)
		default:
			panic(fmt.Sprintf("ranger called with unsupported type: %T", items))
		}
	}
}

func rangeSlice[E any](items []E, component func(int, E) roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		for i, item := range items {
			err := component(i, item)(w, opts)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func rangeSeq2[E any](items iter.Seq2[int, E], component func(int, E) roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		for i, item := range items {
			err := component(i, item)(w, opts)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func rangeSeq[E any](items iter.Seq[E], component func(int, E) roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		i := 0
		for item := range items {
			err := component(i, item)(w, opts)
			if err != nil {
				return err
			}
			i++
		}

		return nil
	}
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
