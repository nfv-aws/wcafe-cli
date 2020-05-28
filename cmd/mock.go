package cmd

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

func newMockServer() (*http.ServeMux, *url.URL) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	mockServerURL, _ := url.Parse(server.URL)
	return mux, mockServerURL
}
