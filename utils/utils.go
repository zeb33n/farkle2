package utils

import "os"

func WaitForKeypress(verbose bool) string {
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	s := string(b)
	if verbose {
		print(s)
	}
	return s
}
