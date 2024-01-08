package core

import (
	dataPopulation "root/Gen"
	methods "root/Methods"
	pop "root/Population"
	"sync"
)

//одиночный вызов
func singlecal(arr1, arr2, arr3 []dataPopulation.DTO, wg *sync.WaitGroup) {
	
    wg.Add(1)
    go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, wg)
    wg.Wait()
    methods.SelectArraysSync(arr1, arr2, arr3)
}

//множественный вызов
func multiplecall(arr1, arr2, arr3 []dataPopulation.DTO, wg *sync.WaitGroup) {
    for i:=0; i < 10000; i++ {
        wg.Add(1)
        go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, wg)
        wg.Wait()

        methods.SelectArraysSync(arr1, arr2, arr3)
        
        arr1, arr2, arr3 = []dataPopulation.DTO{}, []dataPopulation.DTO{}, []dataPopulation.DTO{}
  
        // count := i
        // fmt.Print(count, "\n")
    }
}

