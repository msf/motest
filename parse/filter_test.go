package parse

import (
	"testing",

	"github.com/stretchr/testify/assert"
)


func TestFilterByDomainEmpty(t *testing.T) {
	urls, err := filterByDomain("monzo.com", []string{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(urls))
}

func TestFilterByDomainNoFilter(t *testing.T) {
	urls, err := filterByDomain("", []string{"anything"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(urls))
}

func TestFilterByDomain(t *testing.T) {
	urls, err := filterByDomain("monzo.com", []string{
		"/relative-url",
		"https://facebook.com/something",
		"http://monzo.com/insecure",
		"https://monzo.com/non-relative-url",
		"https://other.site.com/monzo.com/partnership".
	})
	assert.Nil(t, err)

	expected := []string{
		"/relative-url",
		"http://monzo.com/insecure",
		"https://monzo.com/non-relative-url",
	}
	assert.Equal(t, expected, urls)
}