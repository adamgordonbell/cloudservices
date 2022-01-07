package server

import (
	"encoding/json"
	"fmt"
	"log"
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
	id, err := s.Activities.Insert(req.Activity)
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
	log.Printf("Returning %v\n", activity)
	res := api.ActivityDocument{Activity: activity}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) handleList(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleList")
	var query api.ActivityQueryDocument
	var err error
	if r.Body != http.NoBody {
		fmt.Println("Have Body")
		err = json.NewDecoder(r.Body).Decode(&query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	activities, err := s.Activities.List(query.Offset)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Returning %d items", len(activities))
	err = json.NewEncoder(w).Encode(activities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewHTTPServer(addr string) *http.Server {
	var acc *Activities
	var err error
	if acc, err = NewActivities(); err != nil {
		log.Fatal(err)
	}
	server := &httpServer{
		Activities: acc,
	}
	r := mux.NewRouter()
	r.HandleFunc("/", server.handleInsert).Methods("POST")
	r.HandleFunc("/", server.handleGetByID).Methods("GET")
	r.HandleFunc("/list", server.handleList).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
