package elements

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/html"
)

type elementFunc func(attributes roundly.Attribute, children ...roundly.Element) roundly.Element

func elementTest(
	t *testing.T,
	elementFunc elementFunc,
) {
	t.Helper()

	tag := strings.ToLower(getFunctionName(elementFunc))

	want := fmt.Sprintf(`<%s class="test other"></%s>`, tag, tag)

	w := roundly.NewWriter()
	element := elementFunc(Attributes(Class("test", "other")))

	err := element.RenderElement(w)

	got := w.String()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func getFunctionName(i interface{}) string {
	funcPath := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(funcPath, ".")

	return parts[len(parts)-1]
}
