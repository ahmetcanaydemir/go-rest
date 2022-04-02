package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ahmetcanaydemir/go-rest/pkg/api/service"
	"github.com/ahmetcanaydemir/go-rest/pkg/model"
)

type MongoController struct {
	Service service.RecordsService
}

// NewMongoController creates new instance of MongoController
func NewMongoController() {
	ct := &MongoController{
		Service: service.NewRecordsService(),
	}
	http.HandleFunc("/mongo", ct.MongoHandler)
}

// MongoHandler handles /mongo POST requests
func (ct MongoController) MongoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		response, statusCode := ct.PostMongo(w, r)
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

// PostMongo reads MongoRequest from body and gets records from mongodb via records service.
func (ct MongoController) PostMongo(w http.ResponseWriter, r *http.Request) (string, int) {
	var body model.MongoRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}
	if body.StartDate == nil || body.EndDate == nil || body.MinCount == nil || body.MaxCount == nil {
		return "you must send all requested fields", http.StatusBadRequest
	}

	mongoResponse := ct.Service.GetRecords(body)

	json, err := json.Marshal(mongoResponse)
	if err != nil {
		log.Println("json parse error", err)
		return err.Error(), http.StatusInternalServerError
	}

	return string(json), http.StatusOK
}
