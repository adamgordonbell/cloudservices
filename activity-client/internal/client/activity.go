package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	server "github.com/adamgordonbell/cloudservices/activity-log/internal/server"
)

// type Activity struct {
// 	Time        time.Time `json:"time"`
// 	Description string    `json:"description"`
// 	ID          int       `json:"id"`
// }

// func (a Activity) String() string {
// 	return fmt.Sprintf("ID:%d\t\"%s\"\t%d-%d-%d",
// 		a.ID, a.Description, a.Time.Year(), a.Time.Month(), a.Time.Day())
// }

// type ActivityDocument struct {
// 	Activity Activity `json:"activity"`
// }

// type IDDocument struct {
// 	ID int `json:"id"`
// }

type Activities struct {
	URL string
}

func (c *Activities) Insert(activity server.Activity) (int, error) {
	activityDoc := server.ActivityDocument{server.Activity: activity}
	jsBytes, err := json.Marshal(activityDoc)
	if err != nil {
		return -1, err
	}
	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewReader(jsBytes))
	if err != nil {
		return -1, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	var document server.IDDocument
	err = json.Unmarshal(body, &document)
	if err != nil {
		return -1, err
	}
	return document.ID, nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (server.Activity, error) {
	var document server.ActivityDocument
	idDoc := server.IDDocument{ID: id}
	jsBytes, err := json.Marshal(idDoc)
	if err != nil {
		return document.Activity, err
	}
	req, err := http.NewRequest(http.MethodGet, c.URL, bytes.NewReader(jsBytes))
	if err != nil {
		return document.Activity, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return document.Activity, err
	}
	if res.StatusCode == 404 {
		return document.Activity, errors.New("Not Found")
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return document.Activity, err
	}
	err = json.Unmarshal(body, &document)
	if err != nil {
		return document.Activity, err
	}
	return document.Activity, nil
}
