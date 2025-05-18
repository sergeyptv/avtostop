package main

import "fmt"

// Что выведет код и почему?
//
// Выведет:
// start
// end
// 3
// 2
// 1
// Вывод будет таким, т.к. defer будут складываться как пирог и последний созданный defer будет вызван первым при
// завершении программы (LIFO).
func main() {
	fmt.Println("start")
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}
