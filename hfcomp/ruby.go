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
	return func(w roundly.Writer) error {
		_, err := w.WriteString("<ruby>")
		if err != nil {
			return err
		}

		for _, s := range segments {
			err = Group(
				Text(s.Text),
				RubyExplanation(s.Explanation),
			).RenderElement(w)
			if err != nil {
				return err
			}
		}

		_, err = w.WriteString("</ruby>")
		if err != nil {
			return err
		}

		return nil
	}
}

func RubyExplanation(t string) roundly.Element {
	return Group(
		Rp(nil, RawTrusted("(")),
		Rt(nil, RawTrusted(t)),
		Rp(nil, RawTrusted(")")),
	)
}
