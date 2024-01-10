package population

import (
	"math/rand"
	dataPopulation "root/Gen"
	//"sync"
)



func generateRandomNumber() dataPopulation.DTO {
	randomIndex := rand.Intn(len(dataPopulation.ArrObjs[0].Arr2))
	randomElement := dataPopulation.ArrObjs[0].Arr2[randomIndex]
	return randomElement
}
func FillArrays(obj []dataPopulation.ArrayObject) [][]dataPopulation.DTO {	
	newArr := [][]dataPopulation.DTO{}
    for i := 0; i < len(dataPopulation.ArrObjs); i++{
		for j := 0; j < 3; j++{
			obj[i].Arr1 = append(obj[i].Arr1,generateRandomNumber())
			//newArr=append(newArr[i], obj[i].Arr1...)
		}
	}
	for i := 0; i < len(dataPopulation.ArrObjs); i++{
		newArr=append(newArr, dataPopulation.ArrObjs[i].Arr1)
	}
	return newArr

}



// func AddRandomElementsAsync(arr1, arr2, arr3 *[]dataPopulation.DTO, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	mu := sync.Mutex{}

// 	for i := 0; i < 3; i++ {
// 		wg.Add(3)

// 		go func() {
// 			defer wg.Done()

// 			mu.Lock()
// 			*arr1 = append(*arr1, generateRandomNumber())
// 			mu.Unlock()
// 		}()

// 		go func() {
// 			defer wg.Done()

// 			mu.Lock()
// 			*arr2 = append(*arr2, generateRandomNumber())
// 			mu.Unlock()
// 		}()

// 		go func() {
// 			defer wg.Done()

// 			mu.Lock()
// 			*arr3 = append(*arr3, generateRandomNumber())
// 			mu.Unlock()
// 		}()
// 	}	
// }


// func Render () {
// 	arr1 := []po.DTO{}
// 	arr2 := []po.DTO{}
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go addRandomElementsAsync(&arr1, &arr2, &wg)
// 	wg.Wait()

// 	fmt.Println("Array 1:", arr1)
// 	fmt.Println("Array 2:", arr2)
// }