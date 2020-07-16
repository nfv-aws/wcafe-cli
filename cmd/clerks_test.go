package cmd

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var (
	cl = &Clerk{
		Id:   "cl5bafac-b35c-4852-82ca-b272cd79f2f3",
		Name: "sasaki",
	}
)

func TestClerkListOk(t *testing.T) {
	cases := []struct {
		body string
	}{
		{
			body: `[
				{
			      	"Id": "cc5bafac-b35c-4852-82ca-b272cd79f2f3",
-                   "Name": "kato"
				},
				{
					"Id": "cc2jgodl-f03d-7593-83ya-b645cg64f2f5", 
-                   "Name": "kosaka"
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
	body := `{"Id": "` + cl.Id + `", "Name": "` + cl.Name + `"}`

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)
	hundlePath := fmt.Sprintf("/api/clerks")

	// Mockパターンをセット
	mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	})

	create, err := client.ClerkCreate(context.Background(), cl)
	if err != nil {
		t.Fatalf("ClerkCreate was failed:create = %+v, err = %+v", create, err)
	}

	if !reflect.DeepEqual(create, body) {
		t.Errorf("create = %+v, body = %+v", create, body)
	}
}
