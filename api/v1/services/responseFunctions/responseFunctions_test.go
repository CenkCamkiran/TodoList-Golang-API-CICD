package responseFunctions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {

	t.Run("api can response to requests", func(t *testing.T) {
		response := httptest.NewRecorder()

		byteJSON, _ := json.MarshalIndent(bytes.NewBuffer([]byte("{'status': 'testing'")), "", "  ")
		RespondWithJSON(response, http.StatusOK, byteJSON)
		assertResponse(t, response.Code, http.StatusOK)

	})

}

func assertResponse(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	} else {
		t.Logf("response body is correct, got %q want %q", got, want)
	}
}
