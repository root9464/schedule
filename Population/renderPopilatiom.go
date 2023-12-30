package population

import (

	"math/rand"
	"sync"
	"time"
	dataPopulation "root/Gen"
)

var all = []dataPopulation.DTO{
	dataPopulation.Data.Math,
	dataPopulation.Data.Rus,
	dataPopulation.Data.Inf,
	dataPopulation.Data.Fiz,
	dataPopulation.Data.Hist,
	// dataPopulation.Data.Lit,
	// dataPopulation.Data.Mdk,
	// dataPopulation.Data.Arch,
	// dataPopulation.Data.Soc,
	// dataPopulation.Data.Obg,
	// dataPopulation.Data.Fizr,
	// dataPopulation.Data.Angl,
	// dataPopulation.Data.UIPD,
	// dataPopulation.Data.Geo,
}


func generateRandomNumber() dataPopulation.DTO {
	randomIndex := rand.Intn(len(all))
	randomElement := all[randomIndex]
	return randomElement
}
func AddRandomElementsAsync(arr1, arr2, arr3 *[]dataPopulation.DTO, wg *sync.WaitGroup) {
	defer wg.Done()

	mu := sync.Mutex{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		wg.Add(3)

		go func() {
			defer wg.Done()

			mu.Lock()
			*arr1 = append(*arr1, generateRandomNumber())
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			*arr2 = append(*arr2, generateRandomNumber())
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			*arr3 = append(*arr3, generateRandomNumber())
			mu.Unlock()
		}()
	}
}

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