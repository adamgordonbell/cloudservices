package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"testing/quick"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	input := strings.NewReader(`{"activity": {"description": "sledding with nephew", "time":"2021-12-09T16:56:23Z"}}`)
	request := httptest.NewRequest(http.MethodGet, "/", input)
	responseWriter := httptest.NewRecorder()
	server := &httpServer{
		Activities: &Activities{},
	}
	server.handleInsert(responseWriter, request)
	var req IDDocument
	r := responseWriter.Result()
	json.NewDecoder(r.Body).Decode(&req)
	defer request.Body.Close()
	assert.Equal(t, req, IDDocument{ID: 0})
}

func (Activity) Generate(r *rand.Rand, size int) reflect.Value {
	a := Activity{}
	a.Time = time.Unix(rand.Int63(), 0)
	a.Description = Generate(r, size)
	return reflect.ValueOf(a)
}

func Generate(r *rand.Rand, size int) string {
	bs := make([]byte, size)
	for i := range bs {
		bs[i] = byte(r.Intn(128))
	}
	a := string(bs)
	return a
}

func canRoundTrip(a Activity) bool {
	print("1")
	fmt.Printf("Giving:%+v\n", a)
	// create server
	server := &httpServer{
		Activities: &Activities{},
	}

	// insert
	res := ActivityDocument{Activity: a}
	r, _ := json.Marshal(res)
	request := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(r))
	responseWriter := httptest.NewRecorder()
	server.handleInsert(responseWriter, request)

	// retrieve
	request = httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{ "id":0 }`))
	responseWriter = httptest.NewRecorder()
	server.handleGetByID(responseWriter, request)

	var req ActivityDocument
	err := json.NewDecoder(responseWriter.Body).Decode(&req)
	print(err)
	fmt.Printf("got:%+v\n", req.Activity)
	return req.Activity.Time == a.Time && req.Activity.Description == a.Description && req.Activity.ID == a.ID
}

func TestRoundTrip(t *testing.T) {
	c := quick.Config{}
	if err := quick.Check(canRoundTrip, &c); err != nil {
		t.Error(err)
	}
}
