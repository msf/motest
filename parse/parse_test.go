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

	urls, err := parse.URLs("monzo.com", content)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"https://monzo.com/",
		"https://monzo.com/about",
		"https://monzo.com/blog",
		"https://monzo.com/community",
		"https://monzo.com/faq",
		"https://monzo.com/download",
		"https://monzo.com/-play-store-redirect",
		"https://monzo.com/features/apple-pay",
		"https://monzo.com/features/travel",
		"https://monzo.com/features/switch",
		"https://monzo.com/features/overdrafts",
		"https://monzo.com/-play-store-redirect",
		"https://monzo.com/cdn-cgi/l/email-protection#2b434e475b6b464445514405484446",
		"https://monzo.com/community",
		"https://monzo.com/-play-store-redirect",
		"https://monzo.com/about",
		"https://monzo.com/blog",
		"https://monzo.com/press",
		"https://monzo.com/careers",
		"https://web.monzo.com",
		"https://monzo.com/community",
		"https://monzo.com/community/making-monzo",
		"https://monzo.com/transparency",
		"https://monzo.com/blog/how-money-works",
		"https://monzo.com/tone-of-voice",
		"https://monzo.com/faq",
		"https://monzo.com/legal/terms-and-conditions",
		"https://monzo.com/legal/fscs-information",
		"https://monzo.com/legal/privacy-policy",
		"https://monzo.com/legal/cookie-policy",
		"https://monzo.com/-play-store-redirect",
		"https://monzo.com/cdn-cgi/l/email-protection#bbd3ded7cbfbd6d4d5c1d495d8d4d6",
	},
		urls,
	)
}