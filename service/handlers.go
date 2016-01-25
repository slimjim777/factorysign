package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Assertions are the details of the device
type Assertions struct {
	Model        string `json:"model"`
	SerialNumber string `json:"serial"`
	PublicKey    string `json:"publickey"`
}

// SignResponse is the JSON response from the API Sign method
type SignResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"message"`
	Signature    string `json:"signature"`
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
