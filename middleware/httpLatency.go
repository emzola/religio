package middleware

import (
	"log"
	"net/http"
	"time"
)

type HttpLatencyClient struct {
	Logger *log.Logger
	Transport http.RoundTripper
}

func (c HttpLatencyClient) RoundTrip(r *http.Request) (*http.Response, error) {
	startTime := time.Now()
	resp, err := http.DefaultTransport.RoundTrip(r)
	c.Logger.Printf("Request URL: %s\nLatency: %f seconds\n", r.URL, time.Since(startTime).Seconds())
	return resp, err
}