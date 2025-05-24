package hfcomp

import (
	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/attribute"
	. "github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/text"
)

// BdoRtl creates a [element.Bdo] (bidirectional override) with right-to-left direction and the
// given/text.
func BdoRtl(text string) htmfunc.Element {
	return Bdo(DirRtl(), Text(text))
}

// BdoLtr creates a [element.Bdo] (bidirectional override) with left-to-right direction and the
// given text.
func BdoLtr(text string) htmfunc.Element {
	return Bdo(DirLtr(), Text(text))
}

// BdiAuto creates a [element.Bdi] (bidirectional isolation) with automatic direction and the
// given text.
func BdiAuto(text string) htmfunc.Element {
	return Bdi(DirAuto(), Text(text))
}

// BdiRtl creates a [element.Bdi] (bidirectional isolation) with right-to-left direction and the
// given text.
func BdiRtl(text string) htmfunc.Element {
	return Bdi(DirRtl(), Text(text))
}

// BdiLtr creates a [element.Bdi] (bidirectional isolation) with left-to-right direction and the
// given text.
func BdiLtr(text string) htmfunc.Element {
	return Bdi(DirLtr(), Text(text))
}
