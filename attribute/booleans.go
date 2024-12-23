package attribute

import "github.com/ch-schulz/htmfunc"

func Selected() htmfunc.Attribute {
	return BooleanAttribute("selected")
}

func Checked() htmfunc.Attribute {
	return BooleanAttribute("checked")
}
