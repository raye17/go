package client

import (
	"crypto/tls"
	"net/http"
)

type NopTransport struct {
}

func (n *NopTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusTeapot,
	}, nil
}
func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}
	if nop {
		c.Transport = &NopTransport{}
	}
	http.DefaultClient = c
	return c
}
