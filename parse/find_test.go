package parse

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHrefsBasic(t *testing.T) {
	content := []byte(`<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`)
	hrefs, err := findHrefs(content)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"foo",
		"/bar/baz",
	},
		hrefs,
	)
}

func TestFindHrefsMonzoHomepage(t *testing.T) {
	content, err := ioutil.ReadFile("fixtures/monzo.com.html")
	assert.Nil(t, err)

	hrefs, err := findHrefs(content)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"foo",
		"/bar/baz",
	},
		hrefs,
	)
}
