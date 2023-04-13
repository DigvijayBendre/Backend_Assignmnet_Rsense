package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// response is a struct that represents the response message returned from the API.
type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// CreateEmployee is an API handler function that creates a new Employee record in the database.
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// Initialize the response message with a default success status.
	resp := response{
		Code: http.StatusOK,
	}

	// Set the content type of the response to JSON.
	w.Header().Set("Content-Type", "application/json")

	// Decode the request body to an Employee struct.
	var emp Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		// If there is an error during decoding, update the response message with an error status and message.
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Save the name of the Employee to use in the response message.
	name := emp.Name

	// Save the Employee record to the database.
	Database.Create(&emp)

	// Update the response message with a success status and message.
	resp.Message = fmt.Sprintf("record saved successfully for %s", name)

	// Encode and write the Employee struct to the response.
	json.NewEncoder(w).Encode(emp)

	// Encode and write the response message to the response.
	json.NewEncoder(w).Encode(resp)

	// Set the HTTP status code of the response to success.
	w.WriteHeader(http.StatusOK)
}
