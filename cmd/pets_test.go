package cmd

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPetListOk(t *testing.T) {
	cases := []struct {
		body string
	}{
		{
			body: `[
				{
					"id": "sa5bafac-b35c-4852-82ca-b272cd79f2f3",
					"species": "Canine",
					"name": "Shiba lun", 
					"age": 2,
					"store_id": "6a8a6122-7565-4cdf-b8ba-c2b183e4a177",
					"created_time": "2020-05-15T06:14:56Z",
					"updated_time": "2020-06-15T06:55:28Z",
					"status": "PENDING_CREATE"
				},
				{
					"id": "df2jgodl-f03d-7593-83ya-b645cg64f2f5", 
					"species": "Canine",
					"name": "Shiba lun",
					"age": 3,
					"store_id": "6a8a6122-7565-4cdf-b8ba-c2b183e4a177",
					"created_time": "2020-05-15T06:14:56Z", 
					"updated_time": "2020-06-15T06:55:28Z", 
					"status": "PENDING_CREATE"
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
		hundlePath := fmt.Sprintf("/api/pets")

		// mockのパターンをセット
		mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, tc.body)
		})

		res, err := client.PetList(context.Background())
		if err != nil {
			t.Fatalf("PetList was failed:list = %+v, err = %+v", res, err)
		}

		if !reflect.DeepEqual(res, tc.body) {
			t.Errorf("list = %+v, body = %+v", res, tc.body)
		}
	}
}

func TestPetCreateOk(t *testing.T) {
	body := `{"id": "sa5bafac-b35c-4852-82ca-b272cd79f2f3", "species": "Canine","name": "Shiba lun", "age": 2, "store_id": "6a8a6122-7565-4cdf-b8ba-c2b183e4a177"}`

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)
	hundlePath := fmt.Sprintf("/api/pets")

	// Mockパターンをセット
	mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	})

	create, err := client.PetCreate(context.Background(), "6a8a6122-7565-4cdf-b8ba-c2b183e4a177")
	if err != nil {
		t.Fatalf("PetCreate was failed:create = %+v, err = %+v", create, err)
	}

	if !reflect.DeepEqual(create, body) {
		t.Errorf("create = %+v, body = %+v", create, body)
	}
}

func TestPetDeleteOk(t *testing.T) {
	body := `{}`

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)
	hundlePath := fmt.Sprintf("/api/pets/")

	// Mockパターンをセット
	mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	})

	delete, err := client.PetDelete(context.Background(), "6a8a6122-7565-4cdf-b8ba-c2b183e4a177")
	if err != nil {
		t.Fatalf("PetDelete was failed:create = %+v, err = %+v", delete, err)
	}

	if !reflect.DeepEqual(delete, body) {
		t.Errorf("delete = %+v, body = %+v", delete, body)
	}
}
