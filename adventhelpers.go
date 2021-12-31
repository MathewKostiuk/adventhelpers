package adventhelpers

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetInput() []string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s\n", err)
	}

	str := strings.Trim(string(content[:]), "\n")
	return strings.Split(str, "\n")
}
