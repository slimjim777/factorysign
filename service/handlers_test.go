package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignHandlerNoData(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/v1/sign", nil)
	http.HandlerFunc(SignHandler).ServeHTTP(w, r)

	// Check the JSON response
	result := SignResponse{}
	err := json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding the signed response: %v", err)
	}
	if result.Success {
		t.Error("Expected an error, got success response")
	}
}

func TestSignHandler(t *testing.T) {
	const assertions = `
  {
    "model":"聖誕快樂",
    "serial":"A1234/L",
    "publickey":"NNhqloxPyIYXiTP+3JTPWV/mNoBar2geWIf"
  }`

	Config = &ConfigSettings{PrivateKeyPath: "../TestKey.asc"}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/v1/sign", bytes.NewBufferString(assertions))
	http.HandlerFunc(SignHandler).ServeHTTP(w, r)

	// Check the JSON response
	result := SignResponse{}
	err := json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding the signed response: %v", err)
	}
	if !result.Success {
		t.Errorf("Error generated in signing the device: %s", result.ErrorMessage)
	}
	if result.Signature == "" {
		t.Errorf("Empty signed data returned.")
	}
}
