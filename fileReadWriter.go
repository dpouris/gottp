package main

import "io/ioutil"

type fileReader struct {
	fileSrc    string
	byteStream []byte
}

type fileWriter struct {
	fileDst string
}

func (fr fileReader) Read(bs []byte) (int, error) {
	for i, b := range fr.byteStream {
		if b == 0 {
			return i + 1, nil
		}
		bs[i] = b
	}
	return len(fr.byteStream), nil
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
