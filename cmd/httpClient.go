package cmd

import (
	"errors"
	"net"
	"net/http"
	"time"
)

// httpClient returns an HTTP client.
func httpClient() *http.Client {
	redirectPolicyFunc := func (req *http.Request, via []*http.Request) error {
		if len(via) >= 1 {
			return errors.New("stopped after 1 redirect")
		}
		return nil
	}

	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2: true,
		MaxIdleConns: 25,
		IdleConnTimeout: 90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	} 

	return &http.Client{
		CheckRedirect: redirectPolicyFunc,
		Transport: t,
	}
}
