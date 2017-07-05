package main

import (
	"strings"
	"fmt"
)

func main() {

	s := "i want u u want"
	res := WordCount(s)

	fmt.Println(res)

}

func WordCount(s string) map[string]int {

	s_arr := strings.Fields(s)
	s_map := make(map[string]int)

	for i := 0; i < len(s_arr); i++ {
		if s_map[s_arr[i]] == 0 {
			s_map[s_arr[i]] =1
		}else {
			s_map[s_arr[i]] = s_map[s_arr[i]] + 1
		}
	}


	return s_map

}