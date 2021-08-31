package main

import (
	"fmt"
	"strings"
)

var abc = []string{
	"a", "b", "c", "d",
}

func repeditor(abecedario []string) {
	for i := 0; i < len(abecedario); i++ {
		fmt.Print(strings.Repeat(abecedario[i], 10) + "\n")
	}
}
