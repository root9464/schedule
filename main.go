package main

import (
	"fmt"
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

	wg.Add(1)
	go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, &wg)
    wg.Wait()

	methods.SelectArraysSync(arr1, arr2, arr3)

	child1, child2, child3 := methods.KPointCrossover(arr1, arr2, arr3, 2)
	fmt.Println("K-точечный кроссовер:")
	fmt.Println("Child 1:", child1)
	fmt.Println("Child 2:", child2)
	fmt.Println("Child 3:", child3)



}