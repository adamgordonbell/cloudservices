package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Activities *Activities
}

type RecordDocument struct {
	Record Record `json:"record"`
}

type IdDocument struct {
	Id uint64 `json:"id"`
}

func (s *httpServer) handleInsert(w http.ResponseWriter, r *http.Request) {
	var req RecordDocument
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := s.Activities.Insert(req.Record)
	res := IdDocument{Id: id}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) handleGetById(w http.ResponseWriter, r *http.Request) {
	var req IdDocument
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	record, err := s.Activities.Retrieve(req.Id)
	if err == ErrIdNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := RecordDocument{Record: record}
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
	r.HandleFunc("/", server.handleGetById).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
