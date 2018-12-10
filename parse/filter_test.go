package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsoluteURLsForDomainEmpty(t *testing.T) {
	urls, err := absoluteURLsForDomain("monzo.com", "", []string{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(urls))
}

func TestAbsoluteURLsForDomainNoFilter(t *testing.T) {
	urls, err := absoluteURLsForDomain("", "", []string{"/anything"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(urls))
}

func TestAbsoluteURLsForDomain(t *testing.T) {
	urls, err := absoluteURLsForDomain("monzo.com", "https://monzo.com/", []string{
		"/relative-url",
		"https://facebook.com/something",
		"http://monzo.com/insecure",
		"https://monzo.com/non-relative-url",
		"https://other.site.com/monzo.com/partnership",
	})
	assert.Nil(t, err)

	expected := []string{
		"http://monzo.com/insecure",
		"https://monzo.com/non-relative-url",
		"https://monzo.com/relative-url",
	}
	assert.Equal(t, expected, urls)
}

func TestAbsoluteURLsForDomainMonzo(t *testing.T) {
	hrefs := []string{"/", "/about", "/blog", "/community", "/faq", "/download", "https://itunes.apple.com/gb/app/mondo/id1052238659", "/-play-store-redirect", "https://www.theguardian.com/technology/2017/dec/17/monzo-facebook-of-banking", "https://www.telegraph.co.uk/personal-banking/current-accounts/monzo-atom-revolut-starling-everything-need-know-digital-banks/", "https://www.thetimes.co.uk/article/tom-blomfield-the-man-who-made-monzo-g8z59dr8n", "https://www.standard.co.uk/tech/monzo-prepaid-card-current-accounts-challenger-bank-a3805761.html", "/features/apple-pay", "/features/travel", "https://www.fscs.org.uk/", "/features/switch", "/features/overdrafts", "https://itunes.apple.com/gb/app/mondo/id1052238659", "/-play-store-redirect", "/cdn-cgi/l/email-protection#2b434e475b6b464445514405484446", "https://monzo.com/community", "https://itunes.apple.com/gb/app/mondo/id1052238659", "/-play-store-redirect", "/about", "/blog", "/press", "/careers", "https://web.monzo.com", "/community", "/community/making-monzo", "/transparency", "/blog/how-money-works", "/tone-of-voice", "/faq", "/legal/terms-and-conditions", "/legal/fscs-information", "/legal/privacy-policy", "/legal/cookie-policy", "https://itunes.apple.com/gb/app/mondo/id1052238659", "/-play-store-redirect", "https://twitter.com/monzo", "https://www.facebook.com/monzobank", "https://www.linkedin.com/company/monzo-bank", "https://www.youtube.com/monzobank", "/cdn-cgi/l/email-protection#bbd3ded7cbfbd6d4d5c1d495d8d4d6"}
	urls, err := absoluteURLsForDomain("monzo.com", "https://monzo.com/", hrefs)
	assert.Nil(t, err)

	assert.Equal(t, []string{
		"https://monzo.com/",
		"https://monzo.com/-play-store-redirect",
		"https://monzo.com/about",
		"https://monzo.com/blog",
		"https://monzo.com/blog/how-money-works",
		"https://monzo.com/careers",
		"https://monzo.com/cdn-cgi/l/email-protection#2b434e475b6b464445514405484446",
		"https://monzo.com/cdn-cgi/l/email-protection#bbd3ded7cbfbd6d4d5c1d495d8d4d6",
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
	},
		urls,
	)
}
