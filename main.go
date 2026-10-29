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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	result := make([]int, 0, size)

	for i := 0; i < size; i++ {
		result = append(result, rand.Int())
	}

	return result
}


// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]

	for _, num := range data[1:] {
		if num > max {
			max = num
		}
	}

	return max

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup

	if len(data) == 0 {
		return 0
	}
//len1:=len(data)
len:=SIZE/CHUNKS
el:=SIZE-(len*CHUNKS)
	max1:=make([]int, CHUNKS)



wg.Add(1)
go func() {
defer wg.Done()
	for i:=0; i<CHUNKS; i++{
	

	max1[i] = data[(0+i)]

//	for num :=i; num<(len1-el); num=num+CHUNKS {
//
//		if max1[i] < data[num] {
//			max1[i] = data[num]
//
//		}

//	}
y:=len+(len*i)
i1:=len*i
for num :=i1; num<y; num++ {

		if max1[i] < data[num] {
			max1[i] = data[num]

		}

	}



}}()
wg.Wait()
for num :=(SIZE-el); num<SIZE; num++ {
		
max1 = append(max1, data[(num)])
	}

	max2 := data[0]

	for _, num5 := range max1[1:] {


		if num5 > max2 {
			max2 = num5
		}
	}

	return max2

}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data:=generateRandomElements(SIZE)
//elapsed:=3
	fmt.Println("Ищем максимальное значение в один поток")
startLoop := time.Now()
	max:=maximum(data)
elapsed := time.Since(startLoop).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
startLoop1 := time.Now()
	max1:=maxChunks(data)
elapsed = time.Since(startLoop1).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max1, elapsed)
}
