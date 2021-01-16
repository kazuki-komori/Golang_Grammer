package main

import "fmt"

func main() {
	fmt.Println("hello word")
	forArgument()
}

// for ~ of
func forArgument() {
	arr := [3]string{"hoge", "fuga", "piyo"}
	for i, e := range arr {
		fmt.Println(i, e)
	}
}
