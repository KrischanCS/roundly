// Generated file. DO NOT EDIT.

package element

import (
"github.com/KrischanCS/htmfunc"
)

// Audio creates the audio element - Audio player [(More)]
//
// It belongs to the following categories:
// [flow] [phrasing] [embedded] [interactive] [palpable]
//
// It can be parent to the following elements/categories of elements:
// [source] [track] [transparent]
//
// If can itself be a child of the following elements/categories of elements:
// [phrasing]
//
// The following attributes can be added to this element:
// [globals] [src] [crossorigin] [preload] [autoplay] [loop] [muted] [controls]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/media.html#the-audio-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#the-source-element
// [track]: https://html.spec.whatwg.org/dev/media.html#the-track-element
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [src]: https://html.spec.whatwg.org/dev/media.html#attr-media-src
// [crossorigin]: https://html.spec.whatwg.org/dev/media.html#attr-media-crossorigin
// [preload]: https://html.spec.whatwg.org/dev/media.html#attr-media-preload
// [autoplay]: https://html.spec.whatwg.org/dev/media.html#attr-media-autoplay
// [loop]: https://html.spec.whatwg.org/dev/media.html#attr-media-loop
// [muted]: https://html.spec.whatwg.org/dev/media.html#attr-media-muted
// [controls]: https://html.spec.whatwg.org/dev/media.html#attr-media-controls
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Audio(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("audio", attributes, children...)
}

// Track creates the track element - Timed text track [(More)]
//
// It belongs to the following categories:
// none
//
// It can be parent to the following elements/categories of elements:
// empty
//
// If can itself be a child of the following elements/categories of elements:
// [audio] [video]
//
// The following attributes can be added to this element:
// [globals] [default] [kind] [label] [src] [srclang]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/media.html#the-track-element
// [audio]: https://html.spec.whatwg.org/dev/media.html#the-audio-element
// [video]: https://html.spec.whatwg.org/dev/media.html#the-video-element
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [default]: https://html.spec.whatwg.org/dev/media.html#attr-track-default
// [kind]: https://html.spec.whatwg.org/dev/media.html#attr-track-kind
// [label]: https://html.spec.whatwg.org/dev/media.html#attr-track-label
// [src]: https://html.spec.whatwg.org/dev/media.html#attr-track-src
// [srclang]: https://html.spec.whatwg.org/dev/media.html#attr-track-srclang
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Track(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("track", attributes, children...)
}

// Video creates the video element - Video player [(More)]
//
// It belongs to the following categories:
// [flow] [phrasing] [embedded] [interactive] [palpable]
//
// It can be parent to the following elements/categories of elements:
// [source] [track] [transparent]
//
// If can itself be a child of the following elements/categories of elements:
// [phrasing]
//
// The following attributes can be added to this element:
// [globals] [src] [crossorigin] [poster] [preload] [autoplay] [playsinline] [loop] [muted] [controls] [width] [height]
//
// Source: [The HTML Standard for Web Developers/Indices/Elements]
//
// [(More)]: https://html.spec.whatwg.org/dev/media.html#the-video-element
// [flow]: https://html.spec.whatwg.org/dev/dom.html#flow-content-2
// [phrasing]: https://html.spec.whatwg.org/dev/dom.html#phrasing-content-2
// [embedded]: https://html.spec.whatwg.org/dev/dom.html#embedded-content-category
// [interactive]: https://html.spec.whatwg.org/dev/dom.html#interactive-content-2
// [palpable]: https://html.spec.whatwg.org/dev/dom.html#palpable-content-2
// [source]: https://html.spec.whatwg.org/dev/embedded-content.html#the-source-element
// [track]: https://html.spec.whatwg.org/dev/media.html#the-track-element
// [transparent]: https://html.spec.whatwg.org/dev/dom.html#transparent
// [globals]: https://html.spec.whatwg.org/dev/dom.html#global-attributes
// [src]: https://html.spec.whatwg.org/dev/media.html#attr-media-src
// [crossorigin]: https://html.spec.whatwg.org/dev/media.html#attr-media-crossorigin
// [poster]: https://html.spec.whatwg.org/dev/media.html#attr-video-poster
// [preload]: https://html.spec.whatwg.org/dev/media.html#attr-media-preload
// [autoplay]: https://html.spec.whatwg.org/dev/media.html#attr-media-autoplay
// [playsinline]: https://html.spec.whatwg.org/dev/media.html#attr-video-playsinline
// [loop]: https://html.spec.whatwg.org/dev/media.html#attr-media-loop
// [muted]: https://html.spec.whatwg.org/dev/media.html#attr-media-muted
// [controls]: https://html.spec.whatwg.org/dev/media.html#attr-media-controls
// [width]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-width
// [height]: https://html.spec.whatwg.org/dev/embedded-content-other.html#attr-dim-height
// [The HTML Standard for Web Developers/Indices/Elements]: https://html.spec.whatwg.org/dev/indices.html#elements-3
func Video(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element {
    return htmfunc.WriteElement("video", attributes, children...)
}
