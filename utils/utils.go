package utils

import "os"

func WaitForKeypress() string {
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	return string(b)
}
