package hfcomp

import (
	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/logic"
	. "github.com/KrischanCS/roundly/text"
)

type RubySegment struct {
	Text        string
	Explanation string
}

func RubyText(segments []RubySegment) roundly.Element {
	return func(w roundly.Writer, opts ...*roundly.RenderOptions) error {
		el := Ruby(nil,
			Range(segments, func(_ int, segment RubySegment) roundly.Element {
				return Group(
					Text(segment.Text),
					Rp(nil, RawTrusted("(")),
					RtText(segment.Text),
					Rp(nil, RawTrusted(")")),
				)
			}),
		)

		if len(opts) != 0 {
			return el.RenderElementWithOptions(w, opts[0])
		}

		return el.RenderElement(w)
	}
}
