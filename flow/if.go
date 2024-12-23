package flow

import "github.com/ch-schulz/htmfunc"

func If(condition bool, component htmfunc.Element) htmfunc.Element {
	if condition {
		return component
	}

	return func(w htmfunc.Writer) error {
		return nil
	}
}
