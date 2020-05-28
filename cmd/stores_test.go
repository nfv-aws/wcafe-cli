package cmd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoresList(t *testing.T) {
	router := RunStoresListCmd()

	req := httptest.NewRequest("GET", "http://"+dns+":8080/api/v1/stores", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

}

// func TestStoresList(t *testing.T) {
// 	testserver := httptest.NewServer(http.Handler(RunStoresListCmd))
// 	defer testserver.Close()
// 	res, err := http.Get(testserver.URL)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	stores_list, err := ioutil.ReadAll(res.Body)
// 	defer res.Body.Close()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if res.StatusCode != 200 {
// 		t.Error("a response code is not 200")
// 	}
// 	if string(stores_list) != "stores_list" {
// 		t.Error("a response is not stores_list")
// 	}
// }

// func TestStoresList(t *testing.T) {
// 	t.Helper()
// 	t.Run("/stores", func(t *testing.T) {

// 		app := RunStoresListCmd(s)
// 		e := httptest.New(t, app, httptest.URL("http://"+dns+":8080/api/v1/stores"))

// 		e.GET("/stores").WithHeaders(map[string]string{
// 			"Content-Type": "application/json",
// 		}).Expect().Status(httptest.StatusOK)

// 	})
// }
