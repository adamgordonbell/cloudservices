package client

import (
	"context"
	"errors"
	"fmt"
	"log"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type Activities struct {
	client api.ActivityLogServiceClient
}

func NewActivities(URL string) Activities {
	tlsCreds, err := credentials.NewClientTLSFromFile("../activity-log/certs/ca.pem", "")
	if err != nil {
		log.Fatalf("No cert found: %v", err)
	}
	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := api.NewActivityLogServiceClient(conn)
	return Activities{client: client}
}

func (c *Activities) Insert(ctx context.Context, activity *api.Activity) (int, error) {
	resp, err := c.client.Insert(ctx, &api.InsertRequest{Activity: activity})
	if err != nil {
		return 0, fmt.Errorf("Insert failure: %w", err)
	}
	return int(resp.GetId()), nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(ctx context.Context, id int) (*api.Activity, error) {
	resp, err := c.client.Retrieve(ctx, &api.RetrieveRequest{Id: int32(id)})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.NotFound {
			return &api.Activity{}, ErrIDNotFound
		} else {
			return &api.Activity{}, fmt.Errorf("Unexpected Insert failure: %w", err)
		}
	}
	return resp.Activity, nil
}

func (c *Activities) List(ctx context.Context, offset int) ([]*api.Activity, error) {
	resp, err := c.client.List(ctx, &api.ListRequest{Offset: int32(offset)})
	if err != nil {
		return nil, fmt.Errorf("List failure: %w", err)
	}
	return resp.Activities, nil
}
