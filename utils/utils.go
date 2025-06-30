package utils

import "os"

func WaitForKeypress(verbose bool) string {
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	s := string(b)
	if verbose {
		if s != "\n" {
			print(s)
		}
	}
	return s
}

func Contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
