package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	candidate := []string{"Иванов", "Абрамов"}
	votes := []int{1234, 605}

	for i := 1; i <= 1; i++ {

		wg.Add(1)
		go election(i, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for id := range ch {
		for i := 0; i < len(candidate); i++ {
			fmt.Printf("Кандидат %d: [%s]  набрал [%d] голосов\n", id, candidate[i], votes[i])
		}
	}
}

func election(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Идет подсчет голосов по Акмолинской области")
	fmt.Println("Идет подсчет голосов по Жамбыльской области")
	fmt.Println("Идет подсчет голосов по Костанайской области")
	ch <- id
	fmt.Println("подсчет голосов по Акмолинской области завершен")
	fmt.Println("подсчет голосов по Жамбыльской области завершен")
	fmt.Println("подсчет голосов по Костанайской области завершен")
}
