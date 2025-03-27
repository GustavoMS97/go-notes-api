package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GustavoMS97/go-notes-api/internal/app"
	"github.com/GustavoMS97/go-notes-api/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserE2E(t *testing.T) {
	helpers.CleanupCollection("users")
	t.Log("testing create user e2e")
	app := app.InitApp()

	payload := map[string]string{
		"name":     "Test User",
		"email":    "test-e2e@example.com",
		"password": "123456",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	// always cleanup after test
	helpers.CleanupCollection("users")
}
