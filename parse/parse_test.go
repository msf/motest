package parse_test

import (
	"io/ioutil"
	"testing"

	"github.com/msf/motest/parse"
	"github.com/stretchr/testify/assert"
)

func TestParseMonzoHomepage(t *testing.T) {
	content, err := ioutil.ReadFile("fixtures/monzo.com.html")
	assert.Nil(t, err)

	urls, err := parse.URLs("monzo.com", "https://monzo.com/", content)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"https://monzo.com/",
		"https://monzo.com/-play-store-redirect",
		"https://monzo.com/about",
		"https://monzo.com/blog",
		"https://monzo.com/blog/how-money-works",
		"https://monzo.com/careers",
		"https://monzo.com/cdn-cgi/l/email-protection",
		"https://monzo.com/community",
		"https://monzo.com/community/making-monzo",
		"https://monzo.com/download",
		"https://monzo.com/faq",
		"https://monzo.com/features/apple-pay",
		"https://monzo.com/features/overdrafts",
		"https://monzo.com/features/switch",
		"https://monzo.com/features/travel",
		"https://monzo.com/legal/cookie-policy",
		"https://monzo.com/legal/fscs-information",
		"https://monzo.com/legal/privacy-policy",
		"https://monzo.com/legal/terms-and-conditions",
		"https://monzo.com/press",
		"https://monzo.com/tone-of-voice",
		"https://monzo.com/transparency",
		"https://web.monzo.com",
	},
		urls,
	)
}
