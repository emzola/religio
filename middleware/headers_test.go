package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emzola/religio/apidata"
)

func startHTTPServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			w.Header().Set(k, v[0])
		}
	}))
	return ts
}

func TestAddHeaderMiddleware(t *testing.T) {
	testHeaders := map[string]string{
		"X-ClientId":  "test-client",
		"X-Auth-Hash": "random$string",
	}

	client := apidata.Client(testHeaders)

	ts := startHTTPServer()
	defer ts.Close()

	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Fatalf("Expected non-nil error, got: %v", err)
	}

	for k, v := range testHeaders {
		if resp.Header.Get(k) != testHeaders[k] {
			t.Fatalf("Expected header: %s:%s, Got: %s:%s", k, v, k, testHeaders[k])
		}
	}
}
