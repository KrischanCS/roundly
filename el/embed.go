package el

import (
	"github.com/ch-schulz/htmfunc"
)

// Picture creates a [picture element].
//
// The picture element is a container which provides multiple sources to its contained img element to allow authors to
// declaratively control or give hints to the user agent about which image resource to use, based on the screen pixel
// density, viewport size, image format, and other factors. It represents its children.
//
// [picture element]: https://html.spec.whatwg.org/#the-picture-element
func Picture(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("picture", attributes, children...)
}

// Source creates a [source element].
//
// The source element allows authors to specify multiple alternative source sets for img elements or multiple
// alternative media resources for media elements. It does not represent anything on its own.
//
// The type attribute may be present. If present, the value must be a valid MIME type string.
//
// The media attribute may also be present. If present, the value must contain a valid media query list. The user agent
// will skip to the next source element if the value does not match the environment.
//
// [source element]: https://html.spec.whatwg.org/#the-source-element
func Source(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("source", attributes, children...)
}

// Img creates a [img element].
//
// An img element represents an image.
//
// An img element has a dimension attribute source, initially set to the element itself.
//
// The image given by the src and srcset attributes, and any previous sibling source elements' srcset attributes if the
// parent is a picture element, is the embedded content; the value of the alt attribute provides equivalent content for
// those who cannot process images or who have image loading disabled (i.e. it is the img element's fallback content).
//
// The requirements on the alt attribute's value are described in a separate section.
//
// The src attribute must be present, and must contain a valid non-empty URL potentially surrounded by spaces
// referencing a non-interactive, optionally animated, image resource that is neither paged nor scripted.
//
// [img element]: https://html.spec.whatwg.org/#the-img-element
func Img(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("img", attributes, children...)
}

// Iframe creates a [iframe element].
//
// The iframe element represents its content navigable.
//
// The src attribute gives the URL of a page that the element's content navigable is to contain. The attribute, if
// present, must be a valid non-empty URL potentially surrounded by spaces. If the itemprop attribute is specified on an
// iframe element, then the src attribute must also be specified.
//
// The srcdoc attribute gives the content of the page that the element's content navigable is to contain. The value of
// the attribute is used to construct an iframe srcdoc document, which is a Document whose URL matches about:srcdoc.
//
// [iframe element]: https://html.spec.whatwg.org/#the-iframe-element
func Iframe(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("iframe", attributes, children...)
}

// Embed creates a [embed element].
//
// The embed element provides an integration point for an external application or interactive content.
//
// The src attribute gives the URL of the resource being embedded. The attribute, if present, must contain a valid
// non-empty URL potentially surrounded by spaces.
//
// If the itemprop attribute is specified on an embed element, then the src attribute must also be specified.
//
// [embed element]: https://html.spec.whatwg.org/#the-embed-element
func Embed(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("embed", attributes, children...)
}

// Object creates a [object element].
//
// The object element can represent an external resource, which, depending on the type of the resource, will either be
// treated as an image or as a child navigable.
//
// The data attribute specifies the URL of the resource. It must be present, and must contain a valid non-empty URL
// potentially surrounded by spaces.
//
// The type attribute, if present, specifies the type of the resource. If present, the attribute must be a valid MIME
// type string.
//
// The name attribute, if present, must be a valid navigable target name. The given value is used to name the element's
// content navigable, if applicable, and if present when the element's content navigable is created.
//
// [object element]: https://html.spec.whatwg.org/#the-object-element
func Object(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("object", attributes, children...)
}

// Video creates a [video element].
//
// A video element is used for playing videos or movies, and audio files with captions.
//
// Content may be provided inside the video element.User agents should not show this content to the user; it is intended
// for older web browsers which do not support video, so that text can be shown to the users of these older browsers
// informing them of how to access the video contents.
//
// [video element]: https://html.spec.whatwg.org/#the-video-element
func Video(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("video", attributes, children...)
}

// Audio creates a [audio element].
//
// An audio element represents a sound or audio stream.
//
// Content may be provided inside the audio element.User agents should not show this content to the user; it is intended
// for older web browsers which do not support audio, so that text can be shown to the users of these older browsers
// informing them of how to access the audio contents.
//
// [audio element]: https://html.spec.whatwg.org/#the-audio-element
func Audio(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("audio", attributes, children...)
}

// Track creates a [track element].
//
// The track element allows authors to specify explicit external timed text tracks for media elements. It does not
// represent anything on its own.
//
// [track element]: https://html.spec.whatwg.org/#the-track-element
func Track(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("track", attributes, children...)
}

// Map creates a [map element].
//
// The map element, in conjunction with an img element and any area element descendants, defines an image map. The
// element represents its children.
//
// [map element]: https://html.spec.whatwg.org/#the-map-element
func Map(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("map", attributes, children...)
}

// Area creates a [area element].
//
// The area element represents either a hyperlink with some text and a corresponding area on an image map, or a dead
// area on an image map.
//
// [area element]: https://html.spec.whatwg.org/#the-area-element
func Area(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("area", attributes, children...)
}

// Math creates a [math element].
//
// The MathML math element falls into the embedded content, phrasing content, flow content, and palpable content
// categories for the purposes of the content models in this specification.
//
// When the MathML annotation-xml element contains elements from the HTML namespace, such elements must all be flow
// content.
//
// When the MathML token elements (mi, mo, mn, ms, and mtext) are descendants of HTML elements, they may contain
// phrasing content elements from the HTML namespace.
//
// User agents must handle text other than inter-element whitespace found in MathML elements whose content models do
// not allow straight text by pretending for the purposes of MathML content models, layout, and rendering that the text
// is actually wrapped in a MathML mtext element. (Such text is not, however, conforming.)
//
// User agents must act as if any MathML element whose contents does not match the element's content model was replaced,
// for the purposes of MathML layout and rendering, by a MathML merror element containing some appropriate error
// message.
//
// The semantics of MathML elements are defined by [MathML] and [other applicable specifications].
//
// [math element]: https://html.spec.whatwg.org/#the-math-element
// [MathML]: https://html.spec.whatwg.org/multipage/references.html#refsMATHML
// [other applicable specifications]: https://html.spec.whatwg.org/multipage/infrastructure.html#other-applicable-specifications
func Math(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("math", attributes, children...)
}

// Svg creates a [svg element].
//
// The svg element falls into the embedded content, phrasing content, flow content, and palpable content categories
// for the purposes of the content models in this specification.
//
// When the SVG foreignObject element contains elements from the HTML namespace, such elements must all be flow content.
//
// The content model for the SVG title element inside HTML documents is phrasing content.(This further constrains the
// requirements given in SVG 2.)
//
// The semantics of SVG elements are defined by [SVG 2] and [other applicable specifications].
//
// [svg element]: https://html.spec.whatwg.org/#the-svg-element
// [SVG 2]: https://html.spec.whatwg.org/multipage/references.html#refsSVG
// [other applicable specifications]: https://html.spec.whatwg.org/multipage/infrastructure.html#other-applicable-specifications
func Svg(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
	return htmfunc.WriteElement("svg", attributes, children...)
}
