package middleware

import "net/http"

type AddHeaders struct {
	Headers   map[string]string
	Transport http.RoundTripper
}

func (h AddHeaders) RoundTrip(r *http.Request) (*http.Response, error) {
	reqCopy := r.Clone(r.Context())
	for k, v := range h.Headers {
		reqCopy.Header.Add(k, v)
	}
	return http.DefaultTransport.RoundTrip(reqCopy)
}
