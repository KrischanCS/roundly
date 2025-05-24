package hfcomp

import (
	"github.com/KrischanCS/htmfunc"
	. "github.com/KrischanCS/htmfunc/element"
	. "github.com/KrischanCS/htmfunc/logic"
	. "github.com/KrischanCS/htmfunc/text"
)

type RubySegment struct {
	Text        string
	Explanation string
}

func RubyText(segments []RubySegment) htmfunc.Element {
	return func(w htmfunc.Writer) error {
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

func RubyExplanation(t string) htmfunc.Element {
	return Group(
		Rp(nil, RawTrusted("(")),
		Rt(nil, RawTrusted(t)),
		Rp(nil, RawTrusted(")")),
	)
}
