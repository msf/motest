package motest_test

import (
	"testing",

	"github.com/stretchr/testify/assert"
)

func TestFindHrefs(t *testing.T) {

}

func TestFilterByDomain(t *testing.T) {
	
	urls, err := motest.FilterByDomain("monzo.com", []string{
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