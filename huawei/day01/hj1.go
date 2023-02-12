package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//two()
	three()
}

func one() {
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		fmt.Println(lenLastWord(string(bytes)))
	}
}

func lenLastWord(str string) int {
	if len(str) == 0 {
		return 0
	}

	strs := strings.Split(str, " ")
	return len(strs[len(strs)-1])
}

func lenLastWordForStr(str string) int {
	// if len(str) == 0 {
	// 	return 0
	// }

	var res int
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' {
			res++
		} else {
			break
		}
	}

	return res
}

func two() {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		//fmt.Println(str)
		fmt.Println(lenLastWordForStr(str) - 1)
	}
}

func three() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(lenLastWord(str))
	}
}
