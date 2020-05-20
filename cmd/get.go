package cmd

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newGetCmd())
}

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newGetPetsCmd(),
	)

	return cmd
}

func newGetPetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pets",
		Short: "Get Pets list",
		RunE:  runGetPetsCmd,
	}

	return cmd
}

func runGetPetsCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	if len(args) == 0 {
		return errors.New("GetID is required")
	}

	petsID, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.Wrapf(err, "failed to parse GetID: %s", args[0])
	}

	req := PetsShowRequest{
		ID: petsID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.GetPets(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "GetPets was failed: req = %+v, res = %+v", req, res)
	}

	pets := res.Pets
	fmt.Printf(
		// 		"id: %s, name: %s, inserted_at: %v, updated_at: %v\n",
		"id: %s, name: %s",
		// 		pets.ID, pets.Name, pets.InsertedAt, pets.UpdatedAt,
		pets.ID, pets.Name,
	)

	return nil
}

func (client *Client) GetPets(ctx context.Context, apiRequest PetsShowRequest) (*PetsShowResponse, error) {
	subPath := fmt.Sprintf("/app_stacks/%d", apiRequest.ID)
	httpRequest, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse PetsShowResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
