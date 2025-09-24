package main

import (
	"fmt"
	"sync"
	"time"
)

// Расскажи подробно что происходит
//
// Создается Wait Group и 10 000 горутин, которые имитируют сетевые запросы. Каждый такой запрос выполняется
// 1 миллисекунду. После создания горутин код ожидает их завершения - wg.Wait() - и затем выводит счетчик выполненных
// сетевых запросов.
//
// Сколько времени будет выполняться?
//
// Будет выполняться 1 миллисекунду + время на переключение (конкурентность) между горутинами при
// time.Sleep(time.Millisecond) + время на выполнение остальных участков кода.
const numRequests = 10000

var count int

var m sync.Mutex

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	m.Lock()
	count++
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
