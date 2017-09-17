package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"Jeff", "Britta", "Troy", "Abed"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
}
