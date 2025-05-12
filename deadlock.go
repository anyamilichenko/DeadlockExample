package main

import (
	"fmt"
	"sync"
	"time"
)

//Стандартный пример дедлока на мьютексах

func main() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Горутина 1: Блокировка mu1")
		mu1.Lock()
		defer mu1.Unlock()

		time.Sleep(1 * time.Second)
		fmt.Println("Горутина 1: Попытка заблокировать mu2")
		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("Горутина 1: Завершена")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Горутина 2: Блокировка mu2")
		mu2.Lock()
		defer mu2.Unlock()

		time.Sleep(1 * time.Second)
		fmt.Println("Горутина 2: Попытка заблокировать mu1")
		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("Горутина 2: Завершена")
	}()

	wg.Wait()
	fmt.Println("Программа завершена")
}

//Вывод:
//Горутина 1: Блокировка mu1
//Горутина 2: Блокировка mu2
//Горутина 1: Попытка заблокировать mu2
//Горутина 2: Попытка заблокировать mu1
//Тут будет зависание
