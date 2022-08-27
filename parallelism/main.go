package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

/**
Код подставленный из коментариев Stepik для проверки
*/
func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

/**
Passed. OK. Время работы: 11.1µs.
Поведение не случайное, ощибок нет
*/
func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	inputSliceOne, inputSliceTwo := make([]int, n, n), make([]int, n, n)
	var inputSequence int

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(in1 <-chan int, wg *sync.WaitGroup, mu *sync.Mutex, n int) {
			mu.Lock()
			inputSliceOne[inputSequence], inputSliceTwo[inputSequence] = <-in1, <-in2
			inputSequence++
			mu.Unlock()
			wg.Done()
		}(in1, wg, mu, i)
	}

	go func(inputSliceOne []int, inputSliceTwo []int, wg *sync.WaitGroup, mu *sync.Mutex) {
		wg.Wait()

		for key, val := range inputSliceOne {
			mu.Lock()
			inputSliceOne[key] = fn(val)
			mu.Unlock()
		}

		for key, val := range inputSliceTwo {
			mu.Lock()
			inputSliceTwo[key] = fn(val)
			mu.Unlock()
		}

	}(inputSliceOne, inputSliceTwo, wg, mu)

	go func(outputChannel chan<- int, wg *sync.WaitGroup) {
		wg.Wait()
		defer close(outputChannel)
		for key := range inputSliceOne {
			outputChannel <- inputSliceOne[key] + inputSliceTwo[key]
		}
	}(out, wg)
}
