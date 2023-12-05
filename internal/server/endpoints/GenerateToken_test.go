package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	h := &handlers{}

	req, err := http.NewRequest("GET", "/doctor", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	ID := int64(123)
	tokenString, err := h.GenerateToken(w, req, ID, true)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)


	assert.Equal(t, "application/jwt", w.Header().Get("Content-Type"))

	assert.Equal(t, http.StatusOK, w.Code)
}
