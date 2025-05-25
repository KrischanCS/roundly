package attributes

import (
	"log/slog"
	"regexp"
	"strings"
)

var enumPattern = regexp.MustCompile(`^".*?"(;".*?")*(;the empty string)?$`)

func classifyAttributes(attrs []*attribute) attributes {
	slog.Debug("Grouping attributes by type...")

	attrsClassified := newAttributes()

	for _, attr := range attrs {
		classifyAttribute(&attrsClassified, attr)
	}

	slog.Info("Grouping attributes by type.")

	return attrsClassified
}

var enumSeperator = regexp.MustCompile(`; ?`)

func classifyAttribute(attrs *attributes, attr *attribute) {
	switch {
	default:
		attrs.Text = append(attrs.Text, *attr)

	case attr.Value == "[Boolean attribute]":
		attrs.Bool = append(attrs.Bool, *attr)

	case enumPattern.MatchString(attr.Value):
		attr.Values = enumSeperator.Split(attr.Value, -1)
		attrs.Enum = append(attrs.Enum, *attr)

	case strings.HasPrefix(attr.Value, "[input type keyword]"):
		attr.Values = inputTypes
		attrs.InputType = append(attrs.InputType, *attr)

	case isCommaSeparatedList(attr.Value):
		attrs.ListComma = append(attrs.ListComma, *attr)
		return

	case strings.HasPrefix(attr.Value, "[Valid list of floating-point numbers]"):
		attrs.ListCommaFloat = append(attrs.ListCommaFloat, *attr)

	case isSpaceSeparatedList(attr.Value):
		attrs.ListSpace = append(attrs.ListSpace, *attr)

	case strings.HasPrefix(attr.Value, "[Valid floating-point number]"):
		attrs.Float = append(attrs.Float, *attr)

	case strings.Contains(attr.Value, "[Valid integer]"):
		attrs.Int = append(attrs.Int, *attr)

	case strings.HasPrefix(attr.Value, "[Valid non-negative integer]"):
		attrs.Uint = append(attrs.Uint, *attr)
	}
}

func isCommaSeparatedList(value string) bool {
	if strings.HasPrefix(value, "[Set of comma-separated tokens]") {
		return true
	}

	return strings.HasPrefix(value, "Comma-separated list of")
}

func isSpaceSeparatedList(value string) bool {
	if strings.HasPrefix(value, "[Ordered set of unique space-separated tokens]") {
		return true
	}

	if strings.HasPrefix(value, "[Unordered set of unique space-separated tokens]") {
		return true
	}

	return strings.HasPrefix(value, "[Set of space-separated tokens]")
}
