package cmd

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStoreList(t *testing.T) {
	cases := []struct {
		body string
	}{
		{
			body: `[
		{"id":"sa5bafac-b35c-4852-82ca-b272cd79f2f3", "name":"Sano Shinichiro", "tag":"CLI",  "address":"Okinawa"},
		{"id":"sa5bafac-b35c-4852-82ca-b272cd79f2f5", "name":"Suzuki Chihiro", "tag":"CLI",  "address":"Okinawa"}
	]`,
		},
	}

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)

	for _, tc := range cases {
		hundlePath := fmt.Sprintf("/api/stores")
		mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, tc.body)
		})

		list, err := client.StoreList(context.Background())
		if err != nil {
			t.Fatalf("StoreList was failed:list = %+v, err = %+v", list, err)
		}

		if !reflect.DeepEqual(list, tc.body) {
			t.Errorf("list = %+v, body = %+v", list, tc.body)
		}
	}
}

func TestStoreCreate(t *testing.T) {
	cases := []struct {
		body string
	}{
		{
			body: `{"name":"Sano Shinichiro","tag":"CLI", "address":"Okinawa"}`,
		},
	}

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)

	for _, tc := range cases {
		hundlePath := fmt.Sprintf("/api/stores")
		mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, tc.body)
		})

		create, err := client.StoreCreate(context.Background())
		if err != nil {
			t.Fatalf("StoreCreate was failed:create = %+v, err = %+v", create, err)
		}

		if !reflect.DeepEqual(create, tc.body) {
			t.Errorf("create = %+v, body = %+v", create, tc.body)
		}
	}
}
