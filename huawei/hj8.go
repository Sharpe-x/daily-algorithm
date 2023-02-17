package main

import (
	"fmt"
	"sort"
)

func main() {
	var num int
	fmt.Scan(&num)

	var (
		index, value int64
		table        = make(map[int64]int64)
	)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d %d", &index, &value)
		if v, ok := table[index]; ok {
			table[index] = v + value
		} else {
			table[index] = value
		}
	}

	indexs := make([]int, 0, 0)
	for index := range table {
		indexs = append(indexs, int(index))
	}
	sort.Ints(indexs)

	for _, v := range indexs {
		fmt.Printf("%d %d", v, table[int64(v)])
	}

}
