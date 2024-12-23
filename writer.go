package htmfunc

import (
	"bytes"
)

func NewWriter(initialLength int) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, initialLength))
}
