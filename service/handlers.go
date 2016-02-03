package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Assertions are the details of the device
type Assertions struct {
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial"`
	Type         string `json:"type"`
	Revision     int    `json:"revision"`
	PublicKey    string `json:"device-key"`
}

// VersionResponse is the JSON response from the API Version method
type VersionResponse struct {
	Version string `json:"version"`
}

// SignResponse is the JSON response from the API Sign method
type SignResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"message"`
	Signature    string `json:"identity"`
}

// VersionHandler is the API method to return the version of the service
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := VersionResponse{Version: Config.Version}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

// SignHandler is the API method to sign assertions from the device
func SignHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Check we have some data
	if r.Body == nil {
		formatSignResponse(false, "No data supplied for signing.", "", w)
		return
	}
	defer r.Body.Close()

	assertions := new(Assertions)
	err := json.NewDecoder(r.Body).Decode(&assertions)
	if err != nil {
		errorMessage := fmt.Sprintf("Error decoding JSON: %v", err)
		formatSignResponse(false, errorMessage, "", w)
		return
	}

	// Format the assertions string
	dataToSign := formatAssertion(assertions)

	// Read the private key into a string
	privateKey, err := getPrivateKey(Config.PrivateKeyPath)
	if err != nil {
		errorMessage := fmt.Sprintf("Error reading the private key: %v", err)
		formatSignResponse(false, errorMessage, "", w)
		return
	}

	// Sign the assertions
	signedText, err := ClearSign(dataToSign, string(privateKey), "")
	if err != nil {
		fmt.Printf("Error signing the assertions: %v\n", err)
		return
	}

	// Return successful JSON response with the signed text
	formatSignResponse(true, "", string(signedText), w)
}
