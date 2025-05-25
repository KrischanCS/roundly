// Generated file. DO NOT EDIT.

package attribute

import (
	"github.com/KrischanCS/htmfunc"
)

// Attributes joins all given attributes space separated.
func Attributes(attributes ...htmfunc.Attribute) htmfunc.Attribute {
	if attributes == nil {
		return func(w htmfunc.Writer) error {
			return nil
		}
	}

	switch len(attributes) {
	case 0:
		return func(w htmfunc.Writer) error {
			return nil
		}
	case 1:
		return attributes[0]
	default:
		return join(attributes)
	}
}

func join(attributes []htmfunc.Attribute) htmfunc.Attribute {
	return func(w htmfunc.Writer) error {
		err := attributes[0].RenderAttribute(w)
		if err != nil {
			return err
		}

		for _, attribute := range attributes[1:] {
			err = attribute.RenderAttribute(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
