package main

import (
	//"fmt"
	dataPopulation "root/Gen"
	pop "root/Population"
	//"sync"
	//methods "root/Methods"
)


func main() {
    // arr1 := []dataPopulation.DTO{} // группа 1вб1
    // arr2 := []dataPopulation.DTO{} // группа 1вб2
    // arr3 := []dataPopulation.DTO{} // группа 1вб3
    // var wg sync.WaitGroup

	// wg.Add(1)
	// go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, &wg)
    // wg.Wait()
	// methods.SelectArraysSync(arr1, arr2, arr3)
	// fmt.Print("работает")
	// fmt.Print("Array 1:", arr1, " Длина:", len(arr1), "\n")
	// fmt.Print("Array 2:", arr2, " Длина:", len(arr2), "\n")
	// fmt.Print("Array 3:", arr3, " Длина:", len(arr3), "\n")
	pop.FillArrays(dataPopulation.ArrObjs)

	
}