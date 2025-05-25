// Package logic provides functions for meta-elements that are never rendered themselves,
// but affects how their child elements are rendered. There are functions for
//
//   - grouping elements
//   - conditionally rendering elements
//   - repeating elements with different data
package logic

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
