package flow

import "github.com/ch-schulz/htmfunc"

func If(condition bool, component htmfunc.ElementRenderer) htmfunc.ElementRenderer {
	if condition {
		return component
	}

	return htmfunc.WriteElementFunc(func(_ htmfunc.Writer) error {
		return nil
	})
}
