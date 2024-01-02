package methods

import (
	dataPopulation "root/Gen"
	"sync"
    "fmt"
)

/*
   Метод ранжирования
   Задание| Реализовать метод ранжирования на основе создания популяции
*/

// Метод наследования
// !(требуется глобальная доработка)

func crossover(parent1, parent2 []dataPopulation.DTO, crossoverPoint int) ([]dataPopulation.DTO, []dataPopulation.DTO) {
    child1 := append(parent1[:crossoverPoint], parent2[crossoverPoint:]...)
    child2 := append(parent2[:crossoverPoint], parent1[crossoverPoint:]...)
    return child1, child2
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
		child1, child2 := crossover(arr1, arr2, 3)
		fmt.Println("Array 1:", arr1)
		fmt.Println("Array 2:", arr2)
		fmt.Println("Array 3:", arr3)	
		fmt.Println("Array child1:", child1)
		fmt.Println("Array child2:", child2)
	}	
}