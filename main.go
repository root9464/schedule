package main

import (
	"fmt"
	dataPopulation "root/Gen"
	pop "root/Population"
	"sync"
	mth "root/Methods"
)


func asyncLoop(arr1, arr2, arr3 []dataPopulation.DTO) {
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
		fmt.Println("Array 1:", arr1)
		fmt.Println("Array 2:", arr2)
		fmt.Println("Array 3:", arr3)	
	}
	

	
}



func main() {
	arr1 := []dataPopulation.DTO{} // группа 1вб1
	arr2 := []dataPopulation.DTO{} // группа 1вб2
	arr3 := []dataPopulation.DTO{} // группа 1вб3
	var wg sync.WaitGroup

	wg.Add(1)
	go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, &wg) // генерация начальной популяции
	wg.Wait()

	// Ранжирование массивов
	mth.RankingMethod(&arr1)
	mth.RankingMethod(&arr2)
	mth.RankingMethod(&arr3)

	// Вывод отранжированных массивов
	render := 1000
    for i := 0; i < render; i++ {
		fmt.Print("Вызов ", i, "\n")
        asyncLoop(arr1, arr2, arr3)
    }
}

