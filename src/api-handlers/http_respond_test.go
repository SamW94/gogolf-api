package apiHandlers

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SamW94/gogolf-api/helpers"
)

func TestRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()

	respondWithError(rr, 400, "bad request")

	if status := rr.Code; status != 400 {
		t.Errorf("expected status %d, got %d", 400, status)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", ctype)
	}

	expected := `{"error":"bad request"}`
	body := strings.TrimSpace(rr.Body.String())
	if body != expected {
		t.Errorf("expected body %s, got %s", expected, body)
	}
}

func TestRespondWithJSON_StructPayload(t *testing.T) {
	type respBody struct {
		Message string `json:"message"`
	}

	payload := respBody{
		Message: "This is the response body struct/payload passed to the respondWithJSON function.",
	}

	rr := httptest.NewRecorder()
	respondWithJSON(rr, 200, payload)

	if rr.Code != 200 {
		t.Errorf("expected status %d, got %d", 200, rr.Code)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", rr.Header().Get("Content-Type"))
	}

	var responseBody respBody
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("error decoding JSON: %v", err)
	}

	if responseBody != payload {
		t.Errorf("expected %+v, got %+v", payload, responseBody)
	}
}

func TestRespondWithJSON_MapPayload(t *testing.T) {
	payload := map[string]interface{}{
		"status": "success",
		"id":     123,
	}

	rr := httptest.NewRecorder()
	respondWithJSON(rr, 201, payload)

	if rr.Code != 201 {
		t.Errorf("expected status %d, got %d", 201, rr.Code)
	}

	var respBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &respBody); err != nil {
		t.Fatalf("error decoding JSON: %v", err)
	}

	if respBody["status"] != "success" || int(respBody["id"].(float64)) != 123 {
		t.Errorf("unexpected JSON body: %v", respBody)
	}
}

func TestRespondWithJSON_MarshalError(t *testing.T) {
	badPayload := make(chan int)

	rr := httptest.NewRecorder()

	// Capture and check for expected log output from the respondWithJSON function
	logOutput := helpers.CaptureLogOutput(func() {
		respondWithJSON(rr, 200, badPayload)
	})
	if !strings.Contains(logOutput, "Error marshalling JSON to response body:") {
		t.Errorf("expected log output to contain message, got %q", logOutput)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", rr.Header().Get("Content-Type"))
	}

	if rr.Code != 200 {
		t.Errorf("expected status %d, got %d", 200, rr.Code)
	}

	body := rr.Body.String()
	if len(body) == 0 {
		t.Log("body is empty — expected since marshalling failed")
	} else {
		t.Errorf("body returned despite marshalling error: %q", body)
	}
}
