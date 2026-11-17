package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements создаёт слайс заданного размера
// и заполняет его случайными положительными целыми числами.
// Если переданный размер меньше или равен нулю,
// возвращается пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}

	return data
}

// maximum находит максимальное значение в переданном слайсе.
// Если слайс пустой, возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

// maxChunks находит максимальное значение в слайсе,
// разделяя его на несколько частей и обрабатывая их параллельно.
// Для обработки каждой части запускается отдельная горутина.
// После завершения всех горутин функция находит общий максимум
// среди результатов обработки частей слайса.
// Если слайс пустой, возвращает 0.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunks := CHUNKS
	if len(data) < chunks {
		chunks = len(data)
	}

	chunkMaxima := make([]int, chunks)
	var wg sync.WaitGroup

	for i := 0; i < chunks; i++ {
		start := i * len(data) / chunks
		end := (i + 1) * len(data) / chunks

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaxima[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()
	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("\nИщем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("\nМаксимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
