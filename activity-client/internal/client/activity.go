package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Activities struct {
	ctx    context.Context
	client api.Activity_LogClient
	Cancel context.CancelFunc
}

func NewActivities(URL string) Activities {
	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := api.NewActivity_LogClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return Activities{ctx: ctx, client: client, Cancel: cancel}
}

func (c *Activities) Insert(activity *api.Activity) (int, error) {
	resp, err := c.client.Insert(c.ctx, activity)
	if err != nil {
		return 0, fmt.Errorf("Insert failure: %w", err)
	}
	return int(resp.GetId()), nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (*api.Activity, error) {
	resp, err := c.client.Retrieve(c.ctx, &api.RetrieveRequest{Id: int32(id)})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return &api.Activity{}, fmt.Errorf("Unexpected Insert failure: %w", err)
		}
		if st.Code() == codes.NotFound {
			return &api.Activity{}, ErrIDNotFound
		}
	}
	return resp, nil

}

func (c *Activities) List(offset int) ([]*api.Activity, error) {
	resp, err := c.client.List(c.ctx, &api.ListRequest{Offset: int32(offset)})
	if err != nil {
		return nil, fmt.Errorf("List failure: %w", err)
	}
	return resp.Activities, nil
}
