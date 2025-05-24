// Package flow provides typical "flow control" to be used in generating htmfunc
// elements. This includes mainly variants of if and for-range.
package flow

import "github.com/KrischanCS/htmfunc"

type ElementOrAttribute interface {
	htmfunc.Element | htmfunc.Attribute
}

// If returns the given renderer if condition is true, else a NOP renderer.
func If[EA ElementOrAttribute](condition bool, then EA) EA {
	if condition {
		return then
	}

	return func(_ htmfunc.Writer) error {
		return nil
	}
}

// IfElse returns the then renderer, else the otherwise renderer.
func IfElse[EA ElementOrAttribute](condition bool, then, otherwise EA) EA {
	if condition {
		return then
	}

	return otherwise
}

// IfText returns the given string if condition is true, else the empty string.
func IfText(condition bool, str string) string {
	if condition {
		return str
	}

	return ""
}

// IfElseText returns the then string, else the otherwise string.
func IfElseText(condition bool, then, otherwise string) string {
	if condition {
		return then
	}

	return otherwise
}
