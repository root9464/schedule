package methods

import (
	"fmt"
	dataPopulation "root/Gen"
	"sync"
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
	errorOccurred := true

	if len(arr1) != len(arr2) || len(arr1) != len(arr3) {
		fmt.Println("Длины массивов не совпадают")
	}

	wg.Add(len(arr1))

	for i := 0; i < len(arr1); i++ {
		go func(i int) {
			defer wg.Done()
			if arr1[i].Subject == arr2[i].Subject || arr1[i].Subject == arr3[i].Subject || arr2[i].Subject == arr3[i].Subject {
				fmt.Printf("Ошибка: элементы равны на индексе %d во всех массивах\n", i)
				errorOccurred = false
				return
			}
		}(i)
	} //adsd

	wg.Wait()

	if errorOccurred == true {
		fmt.Println("ok")
		child1, child2 := crossover(arr1, arr2, 3)
		fmt.Println("Array 1:", arr1)
		fmt.Println("Array 2:", arr2)
		fmt.Println("Array 3:", arr3)
		fmt.Println("Array child1:", child1)
		fmt.Println("Array child2:", child2)
	}
}
