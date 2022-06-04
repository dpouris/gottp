package main

import (
	"net/http"
)

func getUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	bodyByteStream := make([]byte, 2000000)

	if err != nil {
		return bodyByteStream, err
	}

	resp.Body.Read(bodyByteStream)
	resp.Body.Close()

	return (bodyByteStream), nil
}
