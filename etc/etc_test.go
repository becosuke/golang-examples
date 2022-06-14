package etc

import (
	"flag"
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

func TestFlag(t *testing.T) {
	flagString := flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	flag.Parse()
	t.Logf("%+v", *flagString)
}
