package html

import "github.com/KrischanCS/roundly"

type ElementOrAttribute interface {
	roundly.Element | roundly.Attribute
}

// If returns the given renderer if condition is true, else a NOP renderer.
func If[EA ElementOrAttribute](condition bool, then EA) EA {
	if condition {
		return then
	}

	return func(_ roundly.Writer, _ *roundly.RenderOptions) error {
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

// IfValue returns the given string if condition is true, else an empty string. Intended to be used
// on  multi-value attributes, where some values are optional.
func IfValue(condition bool, then string) string {
	if condition {
		return then
	}

	return ""
}

// IfElseValue returns the then string, else the otherwise string.
func IfElseValue(condition bool, then, otherwise string) string {
	if condition {
		return then
	}

	return otherwise
}
