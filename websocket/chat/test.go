package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)
	m["January"] = 1
	m["February"] = 2

	myfunc(m)
}

func myfunc(m map[string]int){
	if left, right := m["March"]; !right {
		fmt.Println(left)
		fmt.Println(right)
	}
}