package logic

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
	if items == nil {
		return func(_ roundly.Writer, _ *roundly.RenderOptions) error {
			return nil
		}
	}

	switch items := any(items).(type) {
	case []E:
		return rangeSlice(items, component)
	case iter.Seq2[int, E]:
		return rangeSeq2(items, component)
	case iter.Seq[E]:
		return rangeSeq(items, component)
	default:
		panic(fmt.Sprintf("ranger called with unsupported type: %T", items))
	}
}

func rangeSlice[E any](items []E, component func(int, E) roundly.Element) roundly.Element {
	elements := make([]roundly.Element, len(items))
	for i, item := range items {
		elements[i] = component(i, item)
	}

	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		return Group(elements...)(w, opts)
	}
}

func rangeSeq2[E any](items iter.Seq2[int, E], component func(int, E) roundly.Element) roundly.Element {
	elements := make([]roundly.Element, 0, 16)
	for i, item := range items {
		elements = append(elements, component(i, item))
	}

	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		return Group(elements...)(w, opts)
	}
}

func rangeSeq[E any](items iter.Seq[E], component func(int, E) roundly.Element) roundly.Element {
	elements := make([]roundly.Element, 0, 16)
	i := 0

	for item := range items {
		elements = append(elements, component(i, item))
		i++
	}

	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		return Group(elements...)(w, opts)
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
