package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGroup := sort.StringSlice{"Jeff", "Britta", "Troy", "Abed"}
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
}
