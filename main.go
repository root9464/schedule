package main

import (
	dataPopulation "root/Gen"
	"sync"

	methods "root/Methods"
	pop "root/Population"
)

func main() {
    arr1 := []dataPopulation.DTO{} // группа 1вб1
    arr2 := []dataPopulation.DTO{} // группа 1вб2
	arr3 := []dataPopulation.DTO{} // группа 1вб3

	var wg sync.WaitGroup
	singlecal(arr1, arr2, arr3, &wg)
   
}

//одиночный вызов
func singlecal(arr1, arr2, arr3 []dataPopulation.DTO, wg *sync.WaitGroup) {
	
    wg.Add(1)
    go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, wg)
    wg.Wait()

    methods.SelectArraysSync(arr1, arr2, arr3)
}

//множественный вызов
func multiplecall(arr1, arr2, arr3 []dataPopulation.DTO, wg *sync.WaitGroup) {
    for i:=0; i < 15000; i++ {
        wg.Add(1)
        go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, wg)
        wg.Wait()

        methods.SelectArraysSync(arr1, arr2, arr3)
        
        arr1 = []dataPopulation.DTO{}
        arr2 = []dataPopulation.DTO{}
        arr3 = []dataPopulation.DTO{}   
        
        // count := i
        // fmt.Print(count, "\n")
    }
}