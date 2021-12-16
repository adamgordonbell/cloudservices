package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
)

type httpServer struct {
	Activities *Activities
}

type ActivityDocument struct {
	Activity Activity `json:"activity"`
}

type IDDocument struct {
	ID uint64 `json:"id"`
}

func (s *httpServer) handleInsert(w http.ResponseWriter, r *http.Request) {
	var req ActivityDocument
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	valid := validateJSON(newStr)
	if !valid.Valid() {
		http.Error(w, valid.Errors()[0].Description(), http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(strings.NewReader(newStr)).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := s.Activities.Insert(req.Activity)
	res := IDDocument{ID: id}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validateJSON(s string) gojsonschema.Result {
	schemaLoader := gojsonschema.NewReferenceLoader("file:///Users/adam/sandbox/cloudservices/ActivityLog/ActivitySchema.json")
	result, err := gojsonschema.Validate(schemaLoader, gojsonschema.NewStringLoader(s))
	if err != nil {
		println("error!")
		panic(err.Error())
	}
	return *result
}

func (s *httpServer) handleGetByID(w http.ResponseWriter, r *http.Request) {
	var req IDDocument
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	activity, err := s.Activities.Retrieve(req.ID)
	if err == ErrIDNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := ActivityDocument{Activity: activity}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewHTTPServer(addr string) *http.Server {
	server := &httpServer{
		Activities: &Activities{},
	}
	r := mux.NewRouter()
	r.HandleFunc("/", server.handleInsert).Methods("POST")
	r.HandleFunc("/", server.handleGetByID).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
