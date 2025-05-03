package main

import (
	"fmt"
)

// 1
// Что происходит?
// Как сделать так, чтобы работало?
//
// При инициализации мапы через var m map[string]int - она будет равна nil, т.к. не создается hmap. Для того, чтобы
// программа работала необходимо инициализировать мапу через m := make(map[string]int).
func main() {
	m := make(map[string]int, 7) // Указываем 7 потому, что заранее известна емкость мапы. Это позволит предотвратить эвакуацию данных.
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
