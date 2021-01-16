package arguments

import "fmt"

// console.log
func PrintHoge() {
	fmt.Println("Hoge")
}

// for ~ of
func ForArgument() {
	arr := [3]string{"hoge", "fuga", "piyo"}
	for i, e := range arr {
		fmt.Println(i, e)
	}
}

// if
func IfArgument(i int) bool {
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}
