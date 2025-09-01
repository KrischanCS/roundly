package comps

import (
	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
	el "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/text"
)

// BdoRtl creates a [element.Bdo] (bidirectional override) with right-to-left direction and the
// given/text.
func BdoRtl(text string) roundly.Element {
	return el.Bdo(DirRtl(), Text(text))
}

// BdoLtr creates a [element.Bdo] (bidirectional override) with left-to-right direction and the
// given text.
func BdoLtr(text string) roundly.Element {
	return el.Bdo(DirLtr(), Text(text))
}

// BdiAuto creates a [element.Bdi] (bidirectional isolation) with automatic direction and the
// given text.
func BdiAuto(text string) roundly.Element {
	return el.Bdi(DirAuto(), Text(text))
}

// BdiRtl creates a [element.Bdi] (bidirectional isolation) with right-to-left direction and the
// given text.
func BdiRtl(text string) roundly.Element {
	return el.Bdi(DirRtl(), Text(text))
}

// BdiLtr creates a [element.Bdi] (bidirectional isolation) with left-to-right direction and the
// given text.
func BdiLtr(text string) roundly.Element {
	return el.Bdi(DirLtr(), Text(text))
}
