package core

import "os"

func WaitForKeyPress(verbose bool) string {
	b := make([]byte, 1)
	os.Stdin.Read(b)
	s := string(b)
	if verbose {
		if s != "\n" {
			print(s)
		}
	}
	return s
}
