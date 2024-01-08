package population

import (
	"math/rand"
	dataPopulation "root/Gen"
	"sync"
)

var all = []dataPopulation.DTO{
	dataPopulation.Data.Math,
	dataPopulation.Data.Rus,
	dataPopulation.Data.Inf,
	// dataPopulation.Data.Fiz,
	// dataPopulation.Data.Hist,
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


// generateRandomNumber генерирует случайный элемент из среза all.
// Входной параметр all - срез элементов типа dataPopulation.DTO.
// Возвращает случайный элемент типа dataPopulation.DTO.
func generateRandomNumber(all []dataPopulation.DTO) dataPopulation.DTO {
	randomIndex := rand.Intn(len(all))
	randomElement := all[randomIndex]
	return randomElement
}

// getRandomElements генерирует два случайных элемента из среза all.
// Входной параметр all - срез элементов типа dataPopulation.DTO.
// Возвращает два случайных элемента типа dataPopulation.DTO.
func getRandomElements(all []dataPopulation.DTO) (dataPopulation.DTO, dataPopulation.DTO) {
	randomElement1 := generateRandomNumber(all)
	randomElement2 := generateRandomNumber(all)
	return randomElement1, randomElement2
}

// AddRandomElementsAsync добавляет случайные элементы в три среза arr1, arr2 и arr3 асинхронно.
// Входные параметры:
// - arr1, arr2, arr3: указатели на срезы элементов типа dataPopulation.DTO.
// - wg: указатель на sync.WaitGroup, используемый для синхронизации горутин.
func AddRandomElementsAsync(arr1, arr2, arr3 *[]dataPopulation.DTO, wg *sync.WaitGroup) {
	defer wg.Done()

	mu := sync.Mutex{}

	for i := 0; i < 3; i++ {
		wg.Add(3)

		go func() {
			defer wg.Done()

			randomElement1, randomElement2 := getRandomElements(all)
			higherHourElement := getHigherHourElement(randomElement1, randomElement2)

			mu.Lock()
			*arr1 = append(*arr1, higherHourElement)
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			randomElement1, randomElement2 := getRandomElements(all)
			higherHourElement := getHigherHourElement(randomElement1, randomElement2)

			mu.Lock()
			*arr2 = append(*arr2, higherHourElement)
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			randomElement1, randomElement2 := getRandomElements(all)
			higherHourElement := getHigherHourElement(randomElement1, randomElement2)

			mu.Lock()
			*arr3 = append(*arr3, higherHourElement)
			mu.Unlock()
		}()
	}
}

// getHigherHourElement возвращает элемент с более высоким значением поля Hour из двух элементов.
// Входные параметры:
// - element1, element2: элементы типа dataPopulation.DTO.
func getHigherHourElement(element1, element2 dataPopulation.DTO) dataPopulation.DTO {
	if element1.Hour > element2.Hour {
		return element1
	} else {
		return element2
	}
}