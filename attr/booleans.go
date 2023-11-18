package attr

import "github.com/ch-schulz/htmfunc"

func If(ok bool, attribute htmfunc.Attribute) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		if !ok {
			return nil
		}

		return attribute(w)
	}
}

func Selected() htmfunc.Attribute {
	return BooleanAttribute("selected")
}

func Checked() htmfunc.Attribute {
	return BooleanAttribute("checked")
}
