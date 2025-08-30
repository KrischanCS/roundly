package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/homePage.html
var wantHomePage string

func TestHomePage(t *testing.T) {
	got := HomePage().StringPretty()

	assert.Equal(t, wantHomePage, got)
}
