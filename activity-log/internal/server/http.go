package server

import (
	"encoding/json"
	"net/http"

	api "github.com/adamgordonbell/cloudservices/activity-log"
	"github.com/gorilla/mux"
)

type httpServer struct {
	Activities *Activities
}

func (s *httpServer) handleInsert(w http.ResponseWriter, r *http.Request) {
	var req api.ActivityDocument
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := s.Activities.Insert(req.Activity)
	res := api.IDDocument{ID: id}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) handleGetByID(w http.ResponseWriter, r *http.Request) {
	var req api.IDDocument
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
	res := api.ActivityDocument{Activity: activity}
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
