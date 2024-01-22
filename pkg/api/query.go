package api

import (
	"fmt"
	"net/http"
	"back_go/pkg/handler"
	"log"
	"io/ioutil"
	"encoding/json"
)

func HandlerPostQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a POST request at", r.URL.Path)

	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error decoding JSON from the body", http.StatusBadRequest)
		return
	}

	searchTerm, ok := data["searchTerm"].(string)
	if !ok {
		http.Error(w, "Field 'searchTerm' not found or not a string", http.StatusBadRequest)
		return
	}

	res, err := handler.SearchDocuments(searchTerm)
	if err != nil {
		log.Println("Error processing query request", err)
		http.Error(w, "Error processing query request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}