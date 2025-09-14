// Package html provides functions to render HTML elements, attributes, text and
// meta-elements .
//
// Elements: A function for each element from the HTML standard, which can be
// composed with attributes and child elements.
//
// Attributes: A function for each attribute from the HTML standard. For providing
// multiple attributes to an element, there is the Attributes function which takes
// multiple attributes and combines them.
//
// Text: Functions for rendering text nodes, either escaped (Text) or raw (RawTrusted).
// Also, there are functions for elements which often contain text only, where you can
// pass a string without wrapping it in Text(). They are the elements name suffixed with
// "Text", e.g. H1Text("My Page Title").
//
// Meta: Helpers for grouping, conditionals and iterations.
package html
