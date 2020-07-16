package cmd

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	st = &Store{
		Id:          "sa5bafac-b35c-4852-82ca-b272cd79f2f3",
		Name:        "store_controller_test",
		Tag:         "Board game",
		Address:     "Shinagawa",
		StrongPoint: "helpful",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	now = time.Now()
)

func TestStoreListOk(t *testing.T) {
	cases := []struct {
		body string
	}{
		{
			body: `[
				{
				    "id": "f7de9114-32c3-48c7-a371-e22c28387495",
-                   "name": "Chiba Pets",
-                   "tag": "wcafe",
-                   "address": "Chiba",
-                   "strong_point": "",
					"created_time": "2020-05-15T06:14:56Z",
					"updated_time": "2020-06-15T06:55:28Z"
				},
				{
				    "id": "fd87c3a2-84f9-4170-a30b-5225cbb1d97e",
-                   "name": "Tokyo Pets",
-                   "tag": "wcafe",
-                   "address": "Tokyo",
-                   "strong_point": "",
					"created_time":"2020-05-15T06:14:56Z",
					"updated_time":"2020-06-15T06:55:28Z"
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
		hundlePath := fmt.Sprintf("/api/stores")

		// mockのパターンをセット
		mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, tc.body)
		})

		res, err := client.StoreList(context.Background())
		if err != nil {
			t.Fatalf("StoreList was failed:list = %+v, err = %+v", res, err)
		}

		if !reflect.DeepEqual(res, tc.body) {
			t.Errorf("list = %+v, body = %+v", res, tc.body)
		}
	}
}

func TestStoreCreateOk(t *testing.T) {
	body := `{"name":"` + st.Name + `","tag":"` + st.Tag + `", "address":"` + st.Address + `"}`

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)
	hundlePath := fmt.Sprintf("/api/stores")

	// Mockパターンをセット
	mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	})

	create, err := client.StoreCreate(context.Background(), st)
	if err != nil {
		t.Fatalf("StoreCreate was failed:create = %+v, err = %+v", create, err)
	}

	if !reflect.DeepEqual(create, body) {
		t.Errorf("create = %+v, body = %+v", create, body)
	}
}

func TestStoreDeleteOk(t *testing.T) {
	body := `{}`

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)
	hundlePath := fmt.Sprintf("/api/stores/")

	// Mockパターンをセット
	mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	})

	// get, err := client.StoreDelete(context.Background(), "777")
	// if err != nil {
	// 	t.Fatalf("StoreDelete was failed:get = %+v, err = %+v", get, err)
	// }

	delete, err := client.StoreDelete(context.Background(), "6a8a6122-7565-4cdf-b8ba-c2b183e4a177")
	if err != nil {
		t.Fatalf("StoreDelete was failed:delete = %+v, err = %+v", delete, err)
	}

	if !reflect.DeepEqual(delete, body) {
		t.Errorf("delete = %+v, body = %+v", delete, body)
	}
}
