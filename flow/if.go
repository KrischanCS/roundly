// Package flow provides typical "flow control" to be used in generating htmfunc
// elements. This includes mainly variants of if and for-range.
package flow

import "github.com/KrischanCS/htmfunc"

// If returns the given renderer if condition is true, else a NOP renderer.
func If[T ~func(w htmfunc.Writer) error](condition bool, renderer T) T {
	if condition {
		return renderer
	}

	return func(_ htmfunc.Writer) error {
		return nil
	}
}

// IfStr returns the given string if condition is true, else the empty string.
func IfStr(condition bool, str string) string {
	if condition {
		return str
	}

	return ""
}
