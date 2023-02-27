package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var clients = []Client{}

type Client struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Error struct {
	Message string `json:"message"`
}

func Create(w http.ResponseWriter, r *http.Request) {

	var clientToCreate *Client

	err := json.NewDecoder(r.Body).Decode(&clientToCreate)

	if err != nil {
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
	}

	clientToCreate.ID = uuid.New().String()

	clients = append(clients, *clientToCreate)

	json.NewEncoder(w).Encode(clientToCreate)

}

func GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)

}

func GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, client := range clients {

		if client.ID == params["id"] {
			json.NewEncoder(w).Encode(client)
			return
		}

	}

	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(Error{Message: "Client not found"})

}

func DeleteById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for index, client := range clients {

		if client.ID == params["id"] {
			clients = append(clients[:index], clients[index+1:]...)
			json.NewEncoder(w).Encode(client)
			return
		}

	}

	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(Error{Message: "Client not found"})

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/client", Create).Methods("POST")
	r.HandleFunc("/client", GetAll).Methods("GET")
	r.HandleFunc("/client/{id}", GetById).Methods("GET")
	r.HandleFunc("/client/{id}", DeleteById).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
