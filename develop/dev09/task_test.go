package dev09

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadSite(t *testing.T) {
	expected := "Hello, World!"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer ts.Close()
	url := ts.URL
	actual, err := DownloadSite(url)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, string(actual))
}
