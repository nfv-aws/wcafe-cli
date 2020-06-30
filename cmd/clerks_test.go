package cmd

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestClerkListOk(t *testing.T) {
	cases := []struct {
		body string
	}{
		{
			body: `[
				{
					"NameId": "cc5bafac-b35c-4852-82ca-b272cd79f2f3",
					"Name": "kato"
				},
				{
					"NameId": "cc2jgodl-f03d-7593-83ya-b645cg64f2f5", 
					"Name": "kosaka"
				}
			]`,
		},
		{
			body: "[]",
		},
	}

	for _, tc := range cases {
		mux, mockServerURL := newMockServer()
		client := newTestClient(mockServerURL)
		hundlePath := fmt.Sprintf("/api/clerks")

		// mockのパターンをセット
		mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, tc.body)
		})

		res, err := client.ClerkList(context.Background())
		if err != nil {
			t.Fatalf("ClerkList was failed:list = %+v, err = %+v", res, err)
		}

		if !reflect.DeepEqual(res, tc.body) {
			t.Errorf("list = %+v, body = %+v", res, tc.body)
		}
	}
}

func TestClerkCreateOk(t *testing.T) {
	body := `{"NameId": "sa5bafac-b35c-4852-82ca-b272cd79f2f3", "Name": "sasaki"}`

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)
	hundlePath := fmt.Sprintf("/api/clerks")

	// Mockパターンをセット
	mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	})

	create, err := client.ClerkCreate(context.Background())
	if err != nil {
		t.Fatalf("ClerkCreate was failed:create = %+v, err = %+v", create, err)
	}

	if !reflect.DeepEqual(create, body) {
		t.Errorf("create = %+v, body = %+v", create, body)
	}
}
