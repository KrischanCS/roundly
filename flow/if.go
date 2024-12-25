package flow

import "github.com/ch-schulz/htmfunc"

func If(condition bool, component htmfunc.ElementRenderer) htmfunc.ElementRenderer {
	if condition {
		return component
	}

	return htmfunc.ElementWriteFunc(func(_ htmfunc.Writer) error {
		return nil
	})
}
