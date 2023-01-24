package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// Gerar os números da Mega-Sena
	numbers := generateMegaSenaNumbers()

	// Ordenar os números
	sort.Ints(numbers)

	// Imprimir os números
	fmt.Println(numbers)
}

func generateMegaSenaNumbers() []int {
	// Inicializar o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	var numbers []int
	for len(numbers) < 6 {
		// Gerar um número aleatório entre 1 e 60
		n := rand.Intn(60) + 1

		// Verificar se o número já existe no slice
		exists := false
		for _, num := range numbers {
			if num == n {
				exists = true
				break
			}
		}

		// Adicionar o número ao slice se ele ainda não existir
		if !exists {
			numbers = append(numbers, n)
		}
	}

	return numbers
}
