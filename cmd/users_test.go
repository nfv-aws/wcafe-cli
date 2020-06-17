package cmd

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserListOk(t *testing.T) {
	body := `{
		{"id": "3eaf434f-6eb8-4931-8e8e-e16f9267188e", "number": 2374, "name": "Sano Shinichiro","address": "Okinawa", "email": "test@text.com"},
		{"id": "4bb76b8e-83de-428e-b6b7-80173133c6c0", "number": 7194, "name": "Suzuki Chihiro", "address": "Okinawa", "email": "test@example.com"},
	}`

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)

	hundlePath := fmt.Sprintf("/api/users")

	// mockのパターンをセット
	mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	})

	list, err := client.UserList(context.Background())
	if err != nil {
		t.Fatalf("UserList was failed:list = %+v, err = %+v", list, err)
	}

	if !reflect.DeepEqual(list, body) {
		t.Errorf("list = %+v, body = %+v", list, body)
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
