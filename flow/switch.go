package flow

import "github.com/KrischanCS/htmfunc"

// Switch might be useful? keeping it as a note more than anything else.
func Switch[V any, T ~func(w htmfunc.Writer) error](
	value V,
	fn func(V) T,
) T {
	// TODO does this have any value?
	return func(w htmfunc.Writer) error {
		return fn(value)(w)
	}
}
