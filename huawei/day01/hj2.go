package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, substr string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	scanner.Scan()
	substr = scanner.Text()
	fmt.Println(length(s, substr))
}

func length(str, subStr string) int {
	if len(str) == 0 || len(subStr) == 0 {
		return 0
	}

	return strings.Count(strings.ToUpper(str), strings.ToUpper(subStr))
}
