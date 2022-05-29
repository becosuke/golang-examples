package etc

import (
	"net/url"
	"testing"
)

func TestSW(t *testing.T) {
	SW(0)
	SW(1)
	SW(2)
	SW(3)
}

func TestURL(t *testing.T) {
	u := &url.URL{}
	t.Logf("url:%s", u.String())
}
