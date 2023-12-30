package main

import (
	"fmt"
	dataPopulation "root/Gen"
	pop "root/Population"
	"sync"
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

			// if arr1[i] == arr2[i] && arr1[i] == arr3[i] {
			// 	fmt.Println("Ошибка: элементы равны во всех массивах")
			// 	errorOccurred = true
			// } else if arr1[i] == arr2[i] {
			// 	fmt.Println("Ошибка: элементы равны в arr1 и arr2")
			// 	errorOccurred = true
			// } 

			switch {
				case arr1[i] == arr2[i] && 
					 arr1[i] == arr3[i] &&
					 arr2[i] == arr1[i]	&&
					 arr2[i] == arr3[i]: 
					 fmt.Println("Ошибка: элементы равны во всех массивах")
					 errorOccurred = true

				case arr1[i] == arr2[i]: 
					fmt.Println("Ошибка: элементы равны в arr1 и arr2")
					errorOccurred = true

				case arr1[i] == arr3[i]: 
					fmt.Println("Ошибка: элементы равны в arr1 и arr3")
					errorOccurred = true

				case arr2[i] == arr1[i]: 
					fmt.Println("Ошибка: элементы равны в arr2 и arr1")
					errorOccurred = true

				case arr2[i] == arr3[i]: 
					fmt.Println("Ошибка: элементы равны в arr2 и arr3")
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
	arr1 := []dataPopulation.DTO{} //группа 1вб1
	arr2:= []dataPopulation.DTO{} //группа 1вб2
	arr3 := []dataPopulation.DTO{} //группа 1вб3
	var wg sync.WaitGroup

	wg.Add(1)
	go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, &wg)
	wg.Wait()
	

	for i:= 0; i <= 1000 ; i++ {
		
		asyncLoop(arr1, arr2, arr3)
		
	}
	

}