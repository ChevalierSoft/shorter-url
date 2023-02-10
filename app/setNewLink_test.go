package main

import (
  "net/http"
	"net/http/httptest"
	"testing"

  "github.com/stretchr/testify/assert"
)

func TestSetNewLink(t *testing.T) {
  r := SetRouter()

  w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/l", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\n\"data\": null\n}", w.Body.String())
}
