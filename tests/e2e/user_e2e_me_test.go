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

func TestGetLoggedUserE2E(t *testing.T) {
	t.Log("testing get logged in user /me")

	app := app.InitApp()

	// 1. Create user
	payload := map[string]string{
		"name":     "Logged User",
		"email":    "logged-e2e@example.com",
		"password": "123456",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// 2. Login to get token
	loginPayload := map[string]string{
		"email":    "logged-e2e@example.com",
		"password": "123456",
	}
	loginBody, _ := json.Marshal(loginPayload)

	loginReq := httptest.NewRequest(http.MethodPost, "/api/users/login", bytes.NewBuffer(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")

	loginResp, err := app.Test(loginReq, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, loginResp.StatusCode)

	var loginRespBody map[string]string
	err = json.NewDecoder(loginResp.Body).Decode(&loginRespBody)
	assert.NoError(t, err)

	token := loginRespBody["access_token"]
	assert.NotEmpty(t, token)

	// 3. Call /users/me with Authorization
	meReq := httptest.NewRequest(http.MethodGet, "/api/users/me", nil)
	meReq.Header.Set("Authorization", "Bearer "+token)

	meResp, err := app.Test(meReq, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, meResp.StatusCode)

	var userData map[string]interface{}
	err = json.NewDecoder(meResp.Body).Decode(&userData)
	assert.NoError(t, err)
	assert.Equal(t, "logged-e2e@example.com", userData["email"])
	// always cleanup after test
	helpers.CleanupCollection("users")
}
