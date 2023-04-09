package apidata

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/emzola/religio/middleware"
)

func Client(headers map[string]string) *http.Client {
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          25,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	h := middleware.AddHeaders{
		Headers:   headers,
		Transport: t,
	}

	return &http.Client{
		Transport: h,
	}
}

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return io.ReadAll(r.Body)
}
