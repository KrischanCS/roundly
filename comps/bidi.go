package comps

import (
	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/html"
)

// BdoRtl creates a [htmlement.Bdo] (bidirectional override) with right-to-left direction and the
// given/text.
func BdoRtl(text string) roundly.Element {
	return Bdo(DirRtl(), Text(text))
}

// BdoLtr creates a [htmlement.Bdo] (bidirectional override) with left-to-right direction and the
// given text.
func BdoLtr(text string) roundly.Element {
	return Bdo(DirLtr(), Text(text))
}

// BdiAuto creates a [htmlement.Bdi] (bidirectional isolation) with automatic direction and the
// given text.
func BdiAuto(text string) roundly.Element {
	return Bdi(DirAuto(), Text(text))
}

// BdiRtl creates a [htmlement.Bdi] (bidirectional isolation) with right-to-left direction and the
// given text.
func BdiRtl(text string) roundly.Element {
	return Bdi(DirRtl(), Text(text))
}

// BdiLtr creates a [htmlement.Bdi] (bidirectional isolation) with left-to-right direction and the
// given text.
func BdiLtr(text string) roundly.Element {
	return Bdi(DirLtr(), Text(text))
}
