package flow

import (
	"iter"

	"github.com/valyala/bytebufferpool"

	"github.com/ch-schulz/htmfunc"
)

func Range[T any](items []T, component func(int, T) htmfunc.Element) htmfunc.Element {
	buf := bytebufferpool.Get()
	for i, e := range items {
		err := component(i, e)(buf)
		if err != nil {
			return func(_ htmfunc.Writer) error {
				return err
			}
		}
	}

	return func(w htmfunc.Writer) error {
		_, err := buf.WriteTo(w)
		return err
	}
}

func RangeInt(limit int, component func(int) htmfunc.Element) htmfunc.Element {
	buf := bytebufferpool.Get()
	for i := range limit {
		err := component(i)(buf)
		if err != nil {
			return func(_ htmfunc.Writer) error {
				return err
			}
		}
	}
	return func(w htmfunc.Writer) error {
		_, err := buf.WriteTo(w)
		return err
	}
}

func RangeIter(seq iter.Seq2[int, int], component func(int, int) htmfunc.Element) htmfunc.Element {
	if seq == nil {
		return func(w htmfunc.Writer) error {
			return nil
		}
	}

	buf := bytebufferpool.Get()
	for t1, t2 := range seq {
		err := component(t1, t2)(buf)
		if err != nil {
			return func(_ htmfunc.Writer) error {
				return err
			}
		}
	}

	return func(w htmfunc.Writer) error {
		_, err := buf.WriteTo(w)
		return err
	}
}
