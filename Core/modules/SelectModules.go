package modules

import (
	data "root/Core/data"
	dataPopulation "root/Gen"
	pop "root/Population"
	"sync"
)

func SecondIndex(arr1, arr2, arr3 []dataPopulation.DTO) {
	result := PermutationsCount(data.Group1wb1, data.Group1wb2, data.Group1wb3)
	var wg sync.WaitGroup
	if result == 6 {
		wg.Add(1)
		go pop.AddRandomElementsAsyncIndex(&data.Group2wb1, &data.Group2wb2, &data.Group2wb3, 3,  &wg)
		wg.Wait()
		
	}
}