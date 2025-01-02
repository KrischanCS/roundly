package flow

import "github.com/ch-schulz/htmfunc"

// If returns the given element if condition is true, else a NOP renderer.
func If(condition bool, element htmfunc.Element) htmfunc.Element {
	if condition {
		return element
	}

	return func(_ htmfunc.Writer) error {
		return nil
	}
}

// IfAttr returns the given attribute if condition is true, else a NOP renderer.
func IfAttr(condition bool, attribute htmfunc.Attribute) htmfunc.Attribute {
	if condition {
		return attribute
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
