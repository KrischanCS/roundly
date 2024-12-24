package element

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ch-schulz/htmfunc"
	attr "github.com/ch-schulz/htmfunc/attribute"
)

type elementFunc func(attributes htmfunc.Attribute, children ...htmfunc.Element) htmfunc.Element

func elementTest(
	t *testing.T,
	elementFunc elementFunc,
) {
	t.Helper()

	tag := strings.ToLower(getFunctionName(elementFunc))

	want := fmt.Sprintf(`<%s class="test other"></%s>`, tag, tag)

	w := htmfunc.NewWriter(256)
	element := elementFunc(attr.Join(attr.Class(attr.JoinValues("test", "other"))))

	err := element.RenderHTML(w)

	got := w.String()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func getFunctionName(i interface{}) string {
	funcPath := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(funcPath, ".")

	return parts[len(parts)-1]
}
