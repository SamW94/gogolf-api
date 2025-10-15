package apiHandlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SamW94/gogolf-api/helpers"
	_ "github.com/lib/pq"
)

func TestCreateGolferHandlerSuccessful(t *testing.T) {
	apiCfg, cleanup := setupTestAPI(t)
	defer cleanup()

	body := `{"email":"test@example.com","username":"tester","password":"secure123"}`
	req := httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()

	apiCfg.CreateGolferHandler(rr, req)

	if rr.Code != 201 {
		t.Fatalf("expected 201, got %d, body: %s", rr.Code, rr.Body.String())
	}
}

func TestCreateGolferHandlerInvalidJSON(t *testing.T) {
	apiCfg, cleanup := setupTestAPI(t)
	defer cleanup()

	body := `{email: "test@example.com",username: "tester",password:"secure123"}`
	req := httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()

	apiCfg.CreateGolferHandler(rr, req)

	if rr.Code != 500 {
		t.Fatalf("expected 500, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "Something went wrong") {
		t.Fatalf("expected body to contain 'Something went wrong.', got %s", rr.Body.String())
	}

	// GOTCHA: rr.Body in Go is an io.ReadCloser. It is read once passed to the CreateGolferHandler function.
	//  Once you read from it, the underlying stream is done, so failing to recreate the HTTP request means
	// the body just contains 'EOF' which results in the wrong error message when passed to CaptureLogOutput again.

	req = httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))

	logOutput := helpers.CaptureLogOutput(func() {
		apiCfg.CreateGolferHandler(rr, req)
	})

	if !strings.Contains(logOutput, "Error decoding request body to requestJSON") {
		t.Fatalf("expected log to contain 'Error decoding request body to requestJSON', got '%s'", logOutput)
	}
}

func TestCreateGolferHandlerNoPassword(t *testing.T) {
	apiCfg, cleanup := setupTestAPI(t)
	defer cleanup()

	body := `{"email":"test@example.com","username":"tester"}`
	req := httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()

	apiCfg.CreateGolferHandler(rr, req)

	if rr.Code != 400 {
		t.Fatalf("expected status code 400, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "No password provided in request body.") {
		t.Fatalf("expected body of 'No password provided in request body.' got %s", rr.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))

	logOutput := helpers.CaptureLogOutput(func() {
		apiCfg.CreateGolferHandler(rr, req)
	})

	if !strings.Contains(logOutput, "No password provided in request body") {
		t.Fatalf("expected log to contain 'No password provided in request body', got '%s'", logOutput)
	}
}

func TestCreateGolferHandlerNoEmailAddress(t *testing.T) {
	apiCfg, cleanup := setupTestAPI(t)
	defer cleanup()

	body := `{"username":"tester","password":"secure123"}`
	req := httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()

	apiCfg.CreateGolferHandler(rr, req)

	if rr.Code != 400 {
		t.Fatalf("expected status code 400, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "No email address provided in request body") {
		t.Fatalf("expected body of 'No email address provided in request body' got %s", rr.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))

	logOutput := helpers.CaptureLogOutput(func() {
		apiCfg.CreateGolferHandler(rr, req)
	})

	if !strings.Contains(logOutput, "No email address provided in request body") {
		t.Fatalf("expected log to contain 'No password provided in request body', got '%s'", logOutput)
	}
}

func TestCreateGolferHandlerNoUsername(t *testing.T) {
	apiCfg, cleanup := setupTestAPI(t)
	defer cleanup()

	body := `{"email":"test@example.com","password":"secure123"}`
	req := httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()

	apiCfg.CreateGolferHandler(rr, req)

	if rr.Code != 400 {
		t.Fatalf("expected status code 400, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "No username provided in request body") {
		t.Fatalf("expected body of 'No username provided in request body' got %s", rr.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))

	logOutput := helpers.CaptureLogOutput(func() {
		apiCfg.CreateGolferHandler(rr, req)
	})

	if !strings.Contains(logOutput, "No username provided in request body") {
		t.Fatalf("expected log to contain 'No username provided in request body', got '%s'", logOutput)
	}
}
