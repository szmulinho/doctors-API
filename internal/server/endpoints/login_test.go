package endpoints

import (
	"bytes"
	"encoding/json"
	"github.com/szmulinho/doctors/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	h := &handlers{
	}

	requestBody := `{"login": "sampleLogin", "password": "samplePassword"}`

	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(requestBody))
	assert.NoError(t, err, "Error creating request")

	w := httptest.NewRecorder()

	h.Login(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP status OK")

	var response model.LoginResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error decoding response body")

	assert.Equal(t, "sampleLogin", response.Doctor.Login, "Unexpected doctor login")

}
