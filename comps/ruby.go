package comps

import (
	"github.com/KrischanCS/roundly"
	el "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/logic"
	. "github.com/KrischanCS/roundly/text"
)

type RubySegment struct {
	Text        string
	Explanation string
}

func RubyText(segments []RubySegment) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		rbt := el.Ruby(nil,
			Range(segments, func(_ int, segment RubySegment) roundly.Element {
				return Group(
					Text(segment.Text),
					el.Rp(nil, RawTrusted("(")),
					RtText(segment.Text),
					el.Rp(nil, RawTrusted(")")),
				)
			}),
		)

		return rbt.RenderElementWithOptions(w, opts)
	}
}
