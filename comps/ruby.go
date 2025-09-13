package comps

import (
	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/html"
)

type RubySegment struct {
	Text        string
	Explanation string
}

func RubyText(segments []RubySegment) roundly.Element {
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		rbt := Ruby(nil,
			Range(segments, func(_ int, segment RubySegment) roundly.Element {
				return Group(
					Text(segment.Text),
					Rp(nil, RawTrusted("(")),
					RtText(segment.Text),
					Rp(nil, RawTrusted(")")),
				)
			}),
		)

		return rbt.RenderElementWithOptions(w, opts)
	}
}
