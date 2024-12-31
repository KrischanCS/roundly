package flow

import "github.com/ch-schulz/htmfunc"

// If returns a the given element if condition is true, else a NOP renderer
func If(condition bool, element htmfunc.ElementRenderer) htmfunc.ElementRenderer {
	if condition {
		return element
	}

	return htmfunc.ElementWriteFunc(func(_ htmfunc.Writer) error {
		return nil
	})
}

// IfAttr returns a the given attribute if condition is true, else a NOP renderer
func IfAttr(condition bool, attribute htmfunc.AttributeRenderer) htmfunc.AttributeRenderer {
	if condition {
		return attribute
	}

	return htmfunc.AttributeWriteFunc(func(_ htmfunc.Writer) error {
		return nil
	})
}

// IfStr returns a the given string if condition is true, else the empty string
func IfStr(condition bool, str string) string {
	if condition {
		return str
	}

	return ""
}
