// Generated file. DO NOT EDIT.

package attribute

import (
	"github.com/KrischanCS/roundly"
)

func nopAttribute(w roundly.Writer) error {
	return nil
}

// Attributes joins all given attributes space separated.
func Attributes(attributes ...roundly.Attribute) roundly.Attribute {
	if attributes == nil {
		return nopAttribute
	}

	switch len(attributes) {
	case 0:
		return nopAttribute
	case 1:
		return attributes[0]
	default:
		return join(attributes)
	}
}

func join(attributes []roundly.Attribute) roundly.Attribute {
	return func(w roundly.Writer) error {
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
