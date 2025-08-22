// Generated file. DO NOT EDIT.

package attribute

import (
	"github.com/KrischanCS/roundly"
)

func nopAttribute(_ roundly.Writer, _ *roundly.RenderOptions) error {
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
	return func(w roundly.Writer, opts *roundly.RenderOptions) error {
		if opts == nil {
			return joinNoOpts(w, attributes)
		}

		return joinWithOptions(w, attributes, opts)
	}
}

func joinNoOpts(w roundly.Writer, attributes []roundly.Attribute) error {
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

func joinWithOptions(w roundly.Writer, attributes []roundly.Attribute, opts *roundly.RenderOptions) error {
	if opts.Pretty && len(attributes) >= 3 {
		opts.Indent = "  "
		defer opts.DecreaseIndent()
		err := joinWithLineBrakes(w, attributes, opts)
		return err
	}

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

func joinWithLineBrakes(w roundly.Writer, attributes []roundly.Attribute, opts *roundly.RenderOptions) (err error) {
	for _, attr := range attributes {
		err = w.WriteByte('\n')

		err = attr.RenderAttribute(w)
		if err != nil {
			return err
		}
	}

	return nil
}
