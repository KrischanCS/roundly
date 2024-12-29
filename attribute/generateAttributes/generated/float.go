package generated

import (
    "github.com/ch-schulz/htmfunc"
    "github.com/ch-schulz/htmfunc/attribute"
)

// High creates the high attribute - Low limit of high range
//
// It can be applied to the following elements:
//   - [[meter]]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-high
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func High(high string) htmfunc.AttributeRenderer {
    return attribute.Attribute("high", high)
}

// Low creates the low attribute - High limit of low range
//
// It can be applied to the following elements:
//   - [[meter]]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-low
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Low(low string) htmfunc.AttributeRenderer {
    return attribute.Attribute("low", low)
}

// Max creates the max attribute - Upper bound of range
//
// It can be applied to the following elements:
//   - [[meter]]
//   - [[progress]]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-max
// [progress]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-max
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Max(max string) htmfunc.AttributeRenderer {
    return attribute.Attribute("max", max)
}

// Min creates the min attribute - Lower bound of range
//
// It can be applied to the following elements:
//   - [[meter]]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-min
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Min(min string) htmfunc.AttributeRenderer {
    return attribute.Attribute("min", min)
}

// Optimum creates the optimum attribute - Optimum value in gauge
//
// It can be applied to the following elements:
//   - [[meter]]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-optimum
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Optimum(optimum string) htmfunc.AttributeRenderer {
    return attribute.Attribute("optimum", optimum)
}

// Step creates the step attribute - Granularity to be matched by the form control's value
//
// It can be applied to the following elements:
//   - [[input]]
//
// Value constraints: [Valid floating-point number] greater than zero, or "any"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-step
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Step(step string) htmfunc.AttributeRenderer {
    return attribute.Attribute("step", step)
}

// Value creates the value attribute - Current value of the element
//
// It can be applied to the following elements:
//   - [[meter]]
//   - [[progress]]
//
// Value constraints: [Valid floating-point number]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-value
// [progress]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-value
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value(value string) htmfunc.AttributeRenderer {
    return attribute.Attribute("value", value)
}
