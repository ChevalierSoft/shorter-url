package main

import (
  "net/http"
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_hello(t *testing.T) {
  tr := SetRouter
  //go tr.Run(":80")
  //defer tr.Close()

  t.Run("it should return 200 when health is ok", func(t *testing.T) {
    resp, err := http.Get(tr.Url)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode)
    //assert.Equal(t, `{"response":"hello"}`)
	})
  
}
