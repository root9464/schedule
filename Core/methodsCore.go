package core

import (
	"fmt"
	data "root/Core/data"
	modules "root/Core/modules"
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
    methods.SelectArraysSync(arr1, arr2, arr3, &data.Group1wb1, &data.Group1wb2, &data.Group1wb3)
    //modules.VereficationModule(arr1, arr2, arr3, data.Group1wb1, data.Group1wb2, data.Group1wb3)
}

//множественный вызов
func multiplecall(arr1, arr2, arr3 []dataPopulation.DTO, wg *sync.WaitGroup) {
    for i:=0; i < 10000; i++ {
        wg.Add(1)
        go pop.AddRandomElementsAsync(&arr1, &arr2, &arr3, wg)
        wg.Wait()

        methods.SelectArraysSync(arr1, arr2, arr3, &data.Group1wb1, &data.Group1wb2, &data.Group1wb3)
        
        arr1, arr2, arr3 = []dataPopulation.DTO{}, []dataPopulation.DTO{}, []dataPopulation.DTO{}
        modules.SecondIndex(arr1, arr2, arr3)
    }
    errCh, flag := modules.CompareArrays(data.Group1wb1, data.Group1wb2, data.Group1wb3, data.Group2wb1, data.Group2wb2, data.Group2wb3)
    wg.Add(1)

	go func() {
		defer wg.Done()

		for err := range errCh {
			fmt.Println(err)
		}
	}()

	wg.Wait()
    Log()
	fmt.Println("Все массивы проверены")
    fmt.Println(flag)

}

func Log(){
	fmt.Print("\n","\033[31m","Массивы первой группы", "\033[97m", "\n")
	fmt.Print(data.Group1wb1, "\n")
	fmt.Print(data.Group1wb2, "\n")
	fmt.Print(data.Group1wb3, "\n")

	fmt.Print("\n","\033[31m","Массивы второй группы", "\033[97m", "\n")
	fmt.Print(data.Group2wb1, "\n")
	fmt.Print(data.Group2wb2, "\n")
	fmt.Print(data.Group2wb3, "\n")
}
