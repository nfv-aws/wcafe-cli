package cmd

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

// モックサーバの作成
func newMockServer() (*http.ServeMux, *url.URL) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	mockServerURL, _ := url.Parse(server.URL)
	return mux, mockServerURL
}

// テスト用クライアント
func newTestClient(mockServerURL *url.URL) *Client {
	endpointURL := mockServerURL.String() + "/api"
	httpClient := &http.Client{}
	client, _ := newClient(endpointURL, httpClient)
	return client
}
