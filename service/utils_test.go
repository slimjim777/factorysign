package service

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"gopkg.in/yaml.v2"
)

const TestPrivateKeyPath = "../TestKey.asc"

func TestFormatAssertion(t *testing.T) {
	assertions := Assertions{Model: "聖誕快樂", SerialNumber: "ABC1234",
		PublicKey: "NNhqloxPyIYXiTP+3JTPWV/mNoBar2geWIf/TKTNraWeyGL49TDxunDkf5T8yfCWbOaQCWFsr8yK2oawp3DNBjC4C9eYVN"}

	var identity DeviceAssertion

	response := formatAssertion(&assertions)
	yaml.Unmarshal([]byte(response), &identity)

	if identity.PublicKey != assertions.PublicKey || identity.Model != assertions.Model || identity.Serial != assertions.SerialNumber {
		t.Error("Formatted assertion not as expected.")
	}
}

func TestFormatSignResponse(t *testing.T) {
	const signature = "聖誕快樂NNhqloxPyIYXiTP+3JTPWV/mNoBar2geWIf/TKTNraWeyGL49TDxun"

	w := httptest.NewRecorder()
	formatSignResponse(true, "", signature, w)

	var result SignResponse
	err := json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding the signed response: %v", err)
	}

	if result.Signature != signature || !result.Success || result.ErrorMessage != "" {
		t.Errorf("Signed response not as expected: %v", result)
	}
}

func TestGetPrivateKey(t *testing.T) {
	key, err := getPrivateKey(TestPrivateKeyPath)
	if err != nil {
		t.Errorf("Error reading the private key file: %v", err)
	}
	if len(key) == 0 {
		t.Errorf("Empty private key returned.")
	}
}
