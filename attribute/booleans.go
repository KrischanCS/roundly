package attribute

import "github.com/ch-schulz/htmfunc"

func Selected() htmfunc.AttributeRenderer {
	return BooleanAttribute("selected")
}

func Checked() htmfunc.AttributeRenderer {
	return BooleanAttribute("checked")
}
