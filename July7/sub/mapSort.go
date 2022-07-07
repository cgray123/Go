package sub

import (
	"fmt"
	"sort"
)

func MapSort(m map[string]int) {
	tempMap := make(map[int]string)

	for k, _ := range m {
		tempMap[len(k)] = k
	}
	
	keys := make([]int, 0, len(tempMap))
	for i, _ := range tempMap {
		keys = append(keys, i)
	}

	sort.Sort((sort.IntSlice(keys)))
	for _, j := range keys {
		fmt.Println(tempMap[j], m[tempMap[j]])
	}

}
