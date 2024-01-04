package methods

import (
	"fmt"
	"math/rand"
	dataPopulation "root/Gen"
	"sort"
	"sync"
)

/*
   Метод ранжирования
   Задание| Реализовать метод ранжирования на основе создания популяции
*/

// Кроссоверы
// !(требуется глобальная доработка)


func KPointCrossover(parent1, parent2, parent3 []dataPopulation.DTO, k int) ([]dataPopulation.DTO, []dataPopulation.DTO, []dataPopulation.DTO) {
	child1 := make([]dataPopulation.DTO, len(parent1))
	child2 := make([]dataPopulation.DTO, len(parent2))
	child3 := make([]dataPopulation.DTO, len(parent3))

	
	crossoverPoints := make([]int, k)
	for i := 0; i < k; i++ {
		crossoverPoints[i] = rand.Intn(len(parent1))
	}

	sort.Ints(crossoverPoints)

	copy(child1, parent1)
	copy(child2, parent2)
	copy(child3, parent3)

	for i := 0; i < len(crossoverPoints)-1; i++ {
		start := crossoverPoints[i]
		end := crossoverPoints[i+1]

		for j := start; j < end; j++ {
			child1[j] = parent2[j]
			child2[j] = parent3[j]
			child3[j] = parent1[j]
		}
	}

	return child1, child2, child3
}


// Метод отбора 
// !(требуется легкая доработка)

func SelectArraysSync(arr1, arr2, arr3 []dataPopulation.DTO) {
	var wg sync.WaitGroup
	errorOccurred := false

	if len(arr1) != len(arr2) || len(arr1) != len(arr3) {
		fmt.Println("Длины массивов не совпадают")
	}

	wg.Add(len(arr1))

	for i := 0; i < len(arr1); i++ {
		go func(i int) {
			defer wg.Done()
			switch {
				case arr1[i] == arr2[i] && 
					 arr1[i] == arr3[i] &&
					 arr2[i] == arr1[i]	&&
					 arr2[i] == arr3[i]: 
					 fmt.Println("Ошибка: элементы равны во всех массивах")
					 errorOccurred = true

				case arr1[0] == arr2[0] || arr1[0] == arr3[0] ||
					 arr2[0] == arr1[0] || arr2[0] == arr3[0] || 
					 arr3[0] == arr1[0] || arr3[0] == arr2[0]:
					fmt.Println("Ошибка: элементы равны на 0ом индекси среди каких либо массивов")
					errorOccurred = true

				case arr1[1] == arr2[1] || arr1[1] == arr3[1] ||
					 arr2[1] == arr1[1] || arr2[1] == arr3[1] || 
					 arr3[1] == arr1[1] || arr3[1] == arr2[1]:
				   fmt.Println("Ошибка: элементы равны на 1ом индекси среди каких либо массивов")
				   errorOccurred = true

				case arr1[2] == arr2[2] || arr1[2] == arr3[2] ||
				     arr2[2] == arr1[2] || arr2[2] == arr3[2] || 
				     arr3[2] == arr1[2] || arr3[2] == arr2[2]:
				  fmt.Println("Ошибка: элементы равны на 2ом индекси среди каких либо массивов")
				  errorOccurred = true
			}
		}(i)
	}

	wg.Wait()

	if !errorOccurred {
		fmt.Println("ok")
		// arrays := crossover(arr1, arr2, names)
		fmt.Println("Array 1:", arr1)
		fmt.Println("Array 2:", arr2)
		fmt.Println("Array 3:", arr3)	
		// fmt.Println("Array child1:", arrays["arr1"])
		// fmt.Println("Array child2:", arrays["arr2"])
		// fmt.Println("Array child3:", arrays["arr3"])
		// fmt.Println("Array child4:", arrays["arr4"])
	}	
}