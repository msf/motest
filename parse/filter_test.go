package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsoluteURLsForDomainEmpty(t *testing.T) {
	urls, err := absoluteURLsForDomain("monzo.com", []string{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(urls))
}

func TestAbsoluteURLsForDomainNoFilter(t *testing.T) {
	urls, err := absoluteURLsForDomain("", []string{"/anything"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(urls))
}

func TestAbsoluteURLsForDomain(t *testing.T) {
	urls, err := absoluteURLsForDomain("monzo.com", []string{
		"/relative-url",
		"https://facebook.com/something",
		"http://monzo.com/insecure",
		"https://monzo.com/non-relative-url",
		"https://other.site.com/monzo.com/partnership",
	})
	assert.Nil(t, err)

	expected := []string{
		"https://monzo.com/relative-url",
		"http://monzo.com/insecure",
		"https://monzo.com/non-relative-url",
	}
	assert.Equal(t, expected, urls)
}
