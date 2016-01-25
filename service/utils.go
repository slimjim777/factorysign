package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ConfigSettings defines the parsed config file settings.
type ConfigSettings struct {
	PrivateKeyPath string `yaml:"privateKeyPath"`
}

func formatAssertion(assertions *Assertions) string {
	dataToSign := fmt.Sprintf("%s||%s||%s", assertions.PublicKey, assertions.Model, assertions.SerialNumber)

	return dataToSign
}

// Return the armored private key as a string
func getPrivateKey(privateKeyFilePath string) ([]byte, error) {
	privateKey, err := ioutil.ReadFile(privateKeyFilePath)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func formatSignResponse(success bool, message, signature string, w http.ResponseWriter) {
	response := SignResponse{Success: success, ErrorMessage: message, Signature: signature}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
