package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	numbers := generateMegaSenaNumbers()

	sort.Ints(numbers)

	fmt.Println(numbers)
}

func generateMegaSenaNumbers() []int {
	rand.Seed(time.Now().UnixNano())

	var numbers []int
	for len(numbers) < 6 {
		n := rand.Intn(60) + 1

		exists := false
		for _, num := range numbers {
			if num == n {
				exists = true
				break
			}
		}

		if !exists {
			numbers = append(numbers, n)
		}
	}

	return numbers
}
