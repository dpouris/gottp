package main

import (
	"bytes"
	"net/http"
)

func postUrl(url string, payload []byte) (http.Response, error) {
	// fr := fileReader{fileSrc: "rt", byteStream: make([]byte, 20*1024)}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))

	// endpoint := make([]byte, 30*1024)
	// fmt.Println(fr.Read(endpoint))

	return *resp, err
}
