package main

import "fmt"

// Что выведет код и почему?
//
// Выведет false, т.к. в функции A() происходит return с nil, а в функции B() - return с типом *impl, который
// имплементирует интерфейс I. Таким образом, у нас будет a = nil типа и nil значения, а b = *impl типа и nil значения.
type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil

}
func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	fmt.Println(a == b)
}
