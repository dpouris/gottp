package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func postUrl(url string, p *string) (http.Response, error) {
	payload := make([]byte, 0)
	if *p != "" {
		payload, _ = ioutil.ReadFile(*p)
	}
	payload = zeroByteStripper(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))

	return *resp, err
}
