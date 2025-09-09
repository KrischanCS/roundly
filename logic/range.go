package logic

import (
	"fmt"
	"iter"

	"github.com/KrischanCS/roundly"
)

type Iterable[E any] interface {
	~[]E |
		iter.Seq[E] |
		iter.Seq2[int, E]
}

// Range creates a roundly.Element for each item if items.
func Range[E any, I Iterable[E]](items I, component func(int, E) roundly.Element) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		if items == nil {
			return nil
		}

		if opts == nil {
			return ranger[E, I](w, items, component)
		}

		return rangeOpts[E, I](w, items, component, opts)
	}
}

func ranger[E any, I Iterable[E]](
	w roundly.Writer,
	items I,
	component func(int, E) roundly.Element,
) error {

	switch items := any(items).(type) {
	case []E:
		for i, e := range items {
			err := component(i, e).RenderElement(w)
			if err != nil {
				return err
			}
		}
	case iter.Seq2[int, E]:
		for i, e := range items {
			err := component(i, e).RenderElement(w)
			if err != nil {
				return err
			}
		}
	case iter.Seq[E]:
		i := 0
		for e := range items {
			err := component(i, e).RenderElement(w)
			if err != nil {
				return err
			}
			i++
		}
	default:
		panic(fmt.Sprintf("ranger called with unsupported type: %T", items))
	}

	return nil
}

func rangeOpts[E any, I Iterable[E]](
	w roundly.Writer,
	items I,
	component func(int, E) roundly.Element,
	opt *roundly.RenderOptions,
) error {

	switch items := any(items).(type) {
	case []E:
		for i, e := range items {
			err := component(i, e).RenderElementWithOptions(w, opt)
			if err != nil {
				return err
			}
		}
	case iter.Seq2[int, E]:
		for i, e := range items {
			err := component(i, e).RenderElementWithOptions(w, opt)
			if err != nil {
				return err
			}
		}
	case iter.Seq[E]:
		i := 0
		for e := range items {
			err := component(i, e).RenderElementWithOptions(w, opt)
			if err != nil {
				return err
			}
			i++
		}
	default:
		panic(fmt.Sprintf("ranger called with unsupported type: %T", items))
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
