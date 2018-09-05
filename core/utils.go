package core

import (
	"fmt"
	"os"
	"strings"
)

// PathExists reports whether the named file or directory exists.
func PathExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Check(err error, message string) {
	if err != nil {
		// fmt.Println("Error:")
		// _, file, line, _ := runtime.Caller(1)
		// fmt.Println(line, "\t", file, "\n", err)
		// fmt.Println(message)
		fmt.Printf("Error: %s: %s\n", message, err)
		os.Exit(1)
	}
}

func cleanArgs(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, strings.Trim(str, " "))
		}
	}
	return r
}