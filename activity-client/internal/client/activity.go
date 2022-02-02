package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func (c *Activities) Insert(activity api.Activity) (int, error) {
	resp, err := c.client.Insert(c.ctx, &activity)
	if err != nil {
		return 0, fmt.Errorf("Insert failure: %w", err)
	}
	return int(resp.GetId()), nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (api.Activity, error) {
	resp, err := c.client.Retrieve(c.ctx, &api.RetrieveRequest{Id: int32(id)})
	if err != nil {
		return api.Activity{}, fmt.Errorf("Insert failure: %w", err)
	}
	//todo: handle 404
	return *resp, nil
	// 	var document api.ActivityDocument
	// 	idDoc := api.IDDocument{ID: id}
	// 	jsBytes, err := json.Marshal(idDoc)
	// 	if err != nil {
	// 		return document.Activity, err
	// 	}
	// 	req, err := http.NewRequest(http.MethodGet, c.URL, bytes.NewReader(jsBytes))
	// 	if err != nil {
	// 		return document.Activity, err
	// 	}
	// 	res, err := http.DefaultClient.Do(req)
	// 	if err != nil {
	// 		return document.Activity, err
	// 	}
	// 	if res.StatusCode == 404 {
	// 		return document.Activity, errors.New("Not Found")
	// 	}
	// 	if res.Body != nil {
	// 		defer res.Body.Close()
	// 	}
	// 	body, err := ioutil.ReadAll(res.Body)
	// 	if err != nil {
	// 		return document.Activity, err
	// 	}
	// 	err = json.Unmarshal(body, &document)
	// 	if err != nil {
	// 		return document.Activity, err
	// 	}
	// 	return document.Activity, nil
}

func (c *Activities) List(offset int) ([]*api.Activity, error) {
	resp, err := c.client.List(c.ctx, &api.ListRequest{Offset: int32(offset)})
	if err != nil {
		return nil, fmt.Errorf("List failure: %w", err)
	}
	return resp.Activities, nil
	// 	var list []api.Activity
	// 	queryDoc := api.ActivityQueryDocument{Offset: offset}
	// 	jsBytes, err := json.Marshal(queryDoc)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	req, err := http.NewRequest(http.MethodGet, c.URL+"/list", bytes.NewReader(jsBytes))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	res, err := http.DefaultClient.Do(req)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if res.StatusCode == 404 {
	// 		return nil, errors.New("Not Found")
	// 	}
	// 	if res.Body != nil {
	// 		defer res.Body.Close()
	// 	}
	// 	body, err := ioutil.ReadAll(res.Body)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	err = json.Unmarshal(body, &list)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return list, nil
}
