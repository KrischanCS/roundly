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
