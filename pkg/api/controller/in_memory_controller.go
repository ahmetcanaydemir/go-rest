package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/ahmetcanaydemir/go-rest/pkg/model"
)

var InMemoryDB sync.Map

type InMemoryController struct {
}

// NewInMemoryController creates new instance of InMemoryController
func NewInMemoryController() {
	ct := &InMemoryController{}
	http.HandleFunc("/in-memory", ct.InMemoryHandler)
}

// InMemoryHandler handles /in-memory GET and POST requests
func (ct InMemoryController) InMemoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		response, statusCode := ct.GetInMemory(w, r)
		log.Println(r.Method, r.URL, statusCode)

		if statusCode == http.StatusOK {
			_, err := w.Write([]byte(response))
			if err != nil {
				log.Println(err)
			}
		} else {
			http.Error(w, response, statusCode)

		}
	case "POST":
		response, statusCode := ct.PostInMemory(w, r)
		log.Println(r.Method, r.URL, statusCode)

		if statusCode == http.StatusOK {
			_, err := w.Write([]byte(response))
			if err != nil {
				log.Println(err)
			}
		} else {
			http.Error(w, response, statusCode)
		}
	default:
		log.Println(r.Method, r.URL, http.StatusMethodNotAllowed)

		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(`{"message": "this method is not allowed"}`))
		if err != nil {
			log.Println(err)
		}
	}
}

// GetInMemory finds value of given key in InMemoryDB and returns key and value as response
func (ct InMemoryController) GetInMemory(w http.ResponseWriter, r *http.Request) (string, int) {
	keys, ok := r.URL.Query()["key"]
	if !ok || keys[0] == "" {
		return "you must send key as query param", http.StatusBadRequest
	}

	key := keys[0]
	value, ok := InMemoryDB.Load(key)
	if !ok {
		return key + " not found in in memory database", http.StatusNotFound
	}

	response := model.InMemoryResponse{
		Key:   key,
		Value: value.(string),
	}

	json, _ := json.Marshal(response)

	return string(json), http.StatusOK
}

// PostInMemory adds given key, value to InMemoryDB and echoes the body value as response
func (ct InMemoryController) PostInMemory(w http.ResponseWriter, r *http.Request) (string, int) {
	var body model.InMemoryRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("json parse error", err)
		return err.Error(), http.StatusBadRequest
	}

	if body.Key == nil || body.Value == nil {
		return "you must send key and value fields in body", http.StatusBadRequest
	}

	if *body.Key == "" {
		return "key cannot be empty", http.StatusBadRequest
	}

	json, _ := json.Marshal(body)

	InMemoryDB.Store(*body.Key, *body.Value)
	return string(json), http.StatusOK
}
