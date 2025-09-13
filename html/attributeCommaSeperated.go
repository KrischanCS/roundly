// Generated file. DO NOT EDIT.

package html

import (
	"github.com/KrischanCS/roundly"
)

// Accept creates the accept attribute - Hint for expected file type in [file upload controls]
//
// It can be applied to the following elements:
//   - [input]
//
// Value constraints: [Set of comma-separated tokens] (Additional rules apply, see elements documentation) consisting of [valid MIME type strings with no parameters] or audio/ (Additional rules apply, see elements documentation), video/ (Additional rules apply, see elements documentation), or image/ (Additional rules apply, see elements documentation)
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [input]: https://html.spec.whatwg.org/dev/input.html#attr-input-accept
// [file upload controls]: https://html.spec.whatwg.org/dev/input.html#file-upload-state-(type=file)
// [Set of comma-separated tokens]: https://html.spec.whatwg.org/dev/common-microsyntaxes.html#set-of-comma-separated-tokens
// [valid MIME type strings with no parameters]: https://html.spec.whatwg.org/dev/https://mimesniff.spec.whatwg.org/#valid-mime-type-with-no-parameters
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func Accept(accept ...string) roundly.Attribute {
	return roundly.WriteMultiValueAttribute("accept", ',', accept...)
}

// ImageSrcSet creates the imagesrcset attribute - Images to use in different situations, e.g., high-resolution displays, small monitors, etc. (for [rel]="[preload]")
//
// It can be applied to the following elements:
//   - [link]
//
// Value constraints: Comma-separated list of [image candidate strings]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [link]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-imagesrcset
// [rel]: https://html.spec.whatwg.org/dev/semantics.html#attr-link-rel
// [preload]: https://html.spec.whatwg.org/dev/links.html#link-type-preload
// [image candidate strings]: https://html.spec.whatwg.org/dev/images.html#image-candidate-string
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func ImageSrcSet(imageSrcSet ...string) roundly.Attribute {
	return roundly.WriteMultiValueAttribute("imagesrcset", ',', imageSrcSet...)
}

// SrcSet creates the srcset attribute - Images to use in different situations, e.g., high-resolution displays, small monitors, etc.
//
// It can be applied to the following elements:
//   - [img] [source]
//   - [source]
//
// Value constraints: Comma-separated list of [image candidate strings]
//
// Source: [The HTML Standard for Web Developers/Indices/Attributes]
//
// [img]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-img-srcset
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#attr-source-srcset
// [image candidate strings]: https://html.spec.whatwg.org/dev/images.html#image-candidate-string
// [The HTML Standard for Web Developers/Indices/Attributes]: https://html.spec.whatwg.org/dev/indices.html#attributes-3
func SrcSet(srcSet ...string) roundly.Attribute {
	return roundly.WriteMultiValueAttribute("srcset", ',', srcSet...)
}
