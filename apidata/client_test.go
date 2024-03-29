package apidata

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}))
	return ts
}

func TestFetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()

	client := Client(nil)
	data, err := fetchRemoteResource(client, ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello World"
	got := string(data)

	if expected != got {
		t.Errorf("Expected response to be: %s, Got: %s", expected, got)
	}
}
