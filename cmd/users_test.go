package cmd

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserListOk(t *testing.T) {
	cases := []struct {
		body string
	}{
		{
			body: `[
				{
					"id": "f8c390f5-d2cf-48ce-bfea-f0ca718cc6b2",
					"number": 123,
					"name": "tom",
					"address": "Tokyo",
					"email": "sample@mail.com",
					"created_time": "2020-05-19T00:46:15Z",
					"updated_time": "2020-05-19T00:46:15Z"
				},
				{
					"id": "f42d5d23-0ba7-4127-89d7-13a5021f467d",
					"number": 456,
					"name": "elie",
					"address": "",
					"email": "example@mail.com",
					"created_time": "2020-05-26T04:52:10Z",
					"updated_time": "2020-05-26T04:52:10Z"
				},
				{
					"id": "f42d5d23-0ba7-4127-89d7-13a5021f467d",
					"number": 456,
					"name": "",
					"address": "Tokyo",
					"email": "example@mail.com",
					"created_time": "2020-05-26T04:52:10Z",
					"updated_time": "2020-05-26T04:52:10Z"
				},
			]`,
		},
		{
			body: "[]",
		},
	}

	for _, tc := range cases {
		mux, mockServerURL := newMockServer()
		client := newTestClient(mockServerURL)
		hundlePath := fmt.Sprintf("/api/users")

		// mockのパターンをセット
		mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, tc.body)
		})

		res, err := client.UserList(context.Background())
		if err != nil {
			t.Fatalf("UserList was failed:list = %+v, err = %+v", res, err)
		}

		if !reflect.DeepEqual(res, tc.body) {
			t.Errorf("list = %+v, body = %+v", res, tc.body)
		}
	}
}

func TestUserCreateOk(t *testing.T) {
	body := `{"id": "3eaf434f-6eb8-4931-8e8e-e16f9267188e", "number": 2374, "name": "Sano Shinichiro","address": "Okinawa", "email": "test@text.com"}`

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)
	hundlePath := fmt.Sprintf("/api/users")

	// Mockパターンをセット
	mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	})

	create, err := client.UserCreate(context.Background())
	if err != nil {
		t.Fatalf("UserCreate was failed:create = %+v, err = %+v", create, err)
	}

	if !reflect.DeepEqual(create, body) {
		t.Errorf("create = %+v, body = %+v", create, body)
	}
}
