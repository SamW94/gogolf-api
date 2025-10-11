package apiHandlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateGolferHandlerSuccessful(t *testing.T) {
	apiCfg, cleanup := setupTestAPI(t)
	defer cleanup()

	body := `{"email":"test@example.com","username":"tester","password":"secure123"}`
	req := httptest.NewRequest(http.MethodPost, "/golfer", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()

	apiCfg.CreateGolferHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d, body: %s", rr.Code, rr.Body.String())
	}
}
