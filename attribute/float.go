// Generated file. DO NOT EDIT.

package attribute

import (
	"strconv"

	"github.com/KrischanCS/htmfunc"
)

// High creates the high attribute - Low limit of high range
//
// It can be applied to the following elements:
//   - [meter]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-high
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func High(high float64) htmfunc.Attribute {
	return htmfunc.WriteAttribute("high", strconv.FormatFloat(high, 'f', -1, 64))
}

// Low creates the low attribute - High limit of low range
//
// It can be applied to the following elements:
//   - [meter]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-low
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Low(low float64) htmfunc.Attribute {
	return htmfunc.WriteAttribute("low", strconv.FormatFloat(low, 'f', -1, 64))
}

// Max_MeterProgress creates the max attribute - Upper bound of range
//
// It can be applied to the following elements:
//   - [meter]
//   - [progress]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-max
// [progress]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-max
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Max_MeterProgress(max float64) htmfunc.Attribute {
	return htmfunc.WriteAttribute("max", strconv.FormatFloat(max, 'f', -1, 64))
}

// Min_Meter creates the min attribute - Lower bound of range
//
// It can be applied to the following elements:
//   - [meter]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-min
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Min_Meter(min float64) htmfunc.Attribute {
	return htmfunc.WriteAttribute("min", strconv.FormatFloat(min, 'f', -1, 64))
}

// Optimum creates the optimum attribute - Optimum value in gauge
//
// It can be applied to the following elements:
//   - [meter]
//
// Value constraints: [Valid floating-point number] (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-optimum
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Optimum(optimum float64) htmfunc.Attribute {
	return htmfunc.WriteAttribute("optimum", strconv.FormatFloat(optimum, 'f', -1, 64))
}

// Step creates the step attribute - Granularity to be matched by the form control's value
//
// It can be applied to the following elements:
//   - [input]
//
// Value constraints: [Valid floating-point number] greater than zero, or "any"
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-step
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Step(step float64) htmfunc.Attribute {
	return htmfunc.WriteAttribute("step", strconv.FormatFloat(step, 'f', -1, 64))
}

// Value_MeterProgress creates the value attribute - Current value of the element
//
// It can be applied to the following elements:
//   - [meter]
//   - [progress]
//
// Value constraints: [Valid floating-point number]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [meter]: https://html.spec.whatwg.org/dev/form-elements.html#attr-meter-value
// [progress]: https://html.spec.whatwg.org/dev/form-elements.html#attr-progress-value
// [Valid floating-point number]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#valid-floating-point-number
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Value_MeterProgress(value float64) htmfunc.Attribute {
	return htmfunc.WriteAttribute("value", strconv.FormatFloat(value, 'f', -1, 64))
}
