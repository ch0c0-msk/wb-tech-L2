package dev09

import (
	"io"
	"net/http"
)

// DownloadSite returns html page provided by this url
func DownloadSite(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
