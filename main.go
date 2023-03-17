package main 

import (
	"fmt"
)

func main() {

	word := "selamat malam"

	count := make(map[string]int)

	for _, v := range word {
		i := string(v)

		fmt.Println(i)

		count[i]++
	}

	fmt.Println(count)
}