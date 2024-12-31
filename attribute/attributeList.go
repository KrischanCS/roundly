package attribute

import (
	"github.com/ch-schulz/htmfunc"
)

// Attributes joins all given attributes space separated
func Attributes(attributes ...htmfunc.AttributeRenderer) htmfunc.AttributeRenderer {
	if attributes == nil {
		return htmfunc.AttributeWriteFunc(func(w htmfunc.Writer) error {
			return nil
		})
	}

	switch len(attributes) {
	case 0:
		return htmfunc.AttributeWriteFunc(func(w htmfunc.Writer) error {
			return nil
		})

	case 1:
		return attributes[0]

	default:
		return htmfunc.AttributeWriteFunc(func(w htmfunc.Writer) error {
			err := attributes[0].RenderAttribute(w)
			if err != nil {
				return err
			}

			for _, attribute := range attributes[1:] {
				if attribute == nil {
					continue
				}

				err = w.WriteByte(' ')
				if err != nil {
					return err
				}

				err = attribute.RenderAttribute(w)
				if err != nil {
					return err
				}
			}

			return nil
		})
	}
}
