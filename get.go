package main

import (
	"io/ioutil"
	"net/http"
)

// func main() {
// 	// respBody := make([]byte,20000)
// 	resp, _ := http.Get("https://google.com")

// 	// resp.Body.Read(respBody)

// 	// newRes := make([]byte,0)

// 	// for _, Byte := range respBody {
// 	// 	if Byte == 0 {
// 	// 		break
// 	// 	}
// 	// 	newRes = append(newRes, Byte)
// 	// }

// 	fr := fileWriter{
// 		"hello.txt",
// 	}

// 	// nw, err := fr.Write(newRes)

// 	io.Copy(fr, resp.Body)

// 	// fmt.Println(nw,err)

// }

type fileWriter struct {
	fileDst string
}

func (fr fileWriter) Write(p []byte) (n int, err error) {
	bw := 0
	for index, Byte := range p {
		if Byte == 0 {
			bw = index + 1
			break
		}
		if index+1 == len(p) {
			bw = index + 1
		}
	}
	er := ioutil.WriteFile(fr.fileDst, p, 0755)

	return bw, er
}

func getUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	bodyByteStream := make([]byte, 2000000)

	if err != nil {
		return bodyByteStream, err
	}

	resp.Body.Read(bodyByteStream)

	return (bodyByteStream), nil
}
