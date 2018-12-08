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

	expected := []string{"/", "/about", "/blog", "/community", "/faq", "/download", "https://itunes.apple.com/gb/app/mondo/id1052238659", "/-play-store-redirect", "https://www.theguardian.com/technology/2017/dec/17/monzo-facebook-of-banking", "https://www.telegraph.co.uk/personal-banking/current-accounts/monzo-atom-revolut-starling-everything-need-know-digital-banks/", "https://www.thetimes.co.uk/article/tom-blomfield-the-man-who-made-monzo-g8z59dr8n", "https://www.standard.co.uk/tech/monzo-prepaid-card-current-accounts-challenger-bank-a3805761.html", "/features/apple-pay", "/features/travel", "https://www.fscs.org.uk/", "/features/switch", "/features/overdrafts", "https://itunes.apple.com/gb/app/mondo/id1052238659", "/-play-store-redirect", "/cdn-cgi/l/email-protection#2b434e475b6b464445514405484446", "https://monzo.com/community", "https://itunes.apple.com/gb/app/mondo/id1052238659", "/-play-store-redirect", "/about", "/blog", "/press", "/careers", "https://web.monzo.com", "/community", "/community/making-monzo", "/transparency", "/blog/how-money-works", "/tone-of-voice", "/faq", "/legal/terms-and-conditions", "/legal/fscs-information", "/legal/privacy-policy", "/legal/cookie-policy", "https://itunes.apple.com/gb/app/mondo/id1052238659", "/-play-store-redirect", "https://twitter.com/monzo", "https://www.facebook.com/monzobank", "https://www.linkedin.com/company/monzo-bank", "https://www.youtube.com/monzobank", "/cdn-cgi/l/email-protection#bbd3ded7cbfbd6d4d5c1d495d8d4d6"}
	hrefs, err := findHrefs(content)
	assert.Nil(t, err)
	assert.Equal(t, expected, hrefs)
}
