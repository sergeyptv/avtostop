package main

import (
	"fmt"
	"time"
)

// 1
// Что происходит?
// Как сделать так, чтобы работало?
func main() {
	var m map[string]int
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 2
// После изучения многопоточки
// Расскажи подробно что происходит?(спойлер: почему будет panic)
// Как сделать так, чтобы работало?
func main() {
	x := make(map[int]int, 1)
	go func() { x[1] = 2 }()
	go func() { x[3] = 7 }()
	go func() { x[123] = 10 }()
	go func() { x[1] = 2 }()
	go func() { x[34] = 7 }()
	go func() { x[1432] = 10 }()
	go func() { x[1] = 2 }()
	go func() { x[100] = 7 }()
	go func() { x[34] = 10 }()
	go func() { x[1] = 2 }()
	time.Sleep(100 * time.Millisecond) //блокируемся на 100 миллисекунд
	fmt.Println("x[1] =", x[1])
}
