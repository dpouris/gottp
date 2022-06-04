package main

func zeroByteStripper(bs []byte) []byte {
	sbs := make([]byte, 0)
	for _, b := range bs {
		if b == 0 {
			break
		}
		sbs = append(sbs, b)
	}
	return sbs
}
