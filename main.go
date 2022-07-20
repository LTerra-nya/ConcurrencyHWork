package main

import (
	"fmt"
	"sync"
	"time"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
		{512, 421, 51, 34},
		{4, 1, 124, 100, 1, 124, 100, 1, 124, 100},
		{1},
	}
	var wait sync.WaitGroup
	for i := 0; i < len(n); i++ {
		wait.Add(1)

		go func(i2 int) {

			sum(n[i2], i2)
			defer wait.Done()
		}(i)

	}
	wait.Wait()
	// Ваша реалізація
}

func sum(b []int, num int) {
	sum := 0
	i := 0
	for i < len(b) {
		sum += b[i]
		i++
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("slice %d with %v element(s): %v \n", num, i, sum)

}
