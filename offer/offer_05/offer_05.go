package main

import "strings"

func main() {

}

func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
