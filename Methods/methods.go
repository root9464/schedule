package methods

import (
	"fmt"
	"math/rand"
	dataPopulation "root/Gen"
	"sort"
	"sync"
)

/*
   Метод ранжирования
   Задание| Реализовать метод ранжирования на основе создания популяции
*/

// Кроссоверы
// !(требуется глобальная доработка)


func KPointCrossover(parent1, parent2, parent3 []dataPopulation.DTO, k int) ([]dataPopulation.DTO, []dataPopulation.DTO, []dataPopulation.DTO) {
	child1 := make([]dataPopulation.DTO, len(parent1))
	child2 := make([]dataPopulation.DTO, len(parent2))
	child3 := make([]dataPopulation.DTO, len(parent3))

	crossoverPoints := make([]int, k)
	for i := 0; i < k; i++ {
		crossoverPoints[i] = rand.Intn(len(parent1))
	}

	sort.Ints(crossoverPoints)

	copy(child1, parent1)
	copy(child2, parent2)
	copy(child3, parent3)

	for i := 0; i < len(crossoverPoints)-1; i++ {
		start := crossoverPoints[i]
		end := crossoverPoints[i+1]

		for j := start; j < end; j++ {
			child1[j] = parent2[j]
			child2[j] = parent3[j]
			child3[j] = parent1[j]
		}
	}

	return child1, child2, child3
}

func ThreePointCrossover(parent1, parent2, parent3 []dataPopulation.DTO) ([]dataPopulation.DTO, []dataPopulation.DTO, []dataPopulation.DTO) {
    child1 := make([]dataPopulation.DTO, len(parent1))
    child2 := make([]dataPopulation.DTO, len(parent2))
    child3 := make([]dataPopulation.DTO, len(parent3))

    crossoverPoint1 := rand.Intn(len(parent1))
    crossoverPoint2 := rand.Intn(len(parent1))
    crossoverPoint3 := rand.Intn(len(parent1))

    // Сортируем точки пересечения в порядке возрастания
    crossoverPoints := []int{crossoverPoint1, crossoverPoint2, crossoverPoint3}
    sort.Ints(crossoverPoints)

    copy(child1[:crossoverPoints[0]], parent1[:crossoverPoints[0]])
    copy(child2[:crossoverPoints[0]], parent2[:crossoverPoints[0]])
    copy(child3[:crossoverPoints[0]], parent3[:crossoverPoints[0]])

    copy(child1[crossoverPoints[0]:crossoverPoints[1]], parent2[crossoverPoints[0]:crossoverPoints[1]])
    copy(child2[crossoverPoints[0]:crossoverPoints[1]], parent1[crossoverPoints[0]:crossoverPoints[1]])
    copy(child3[crossoverPoints[0]:crossoverPoints[1]], parent3[crossoverPoints[0]:crossoverPoints[1]])

    copy(child1[crossoverPoints[1]:crossoverPoints[2]], parent3[crossoverPoints[1]:crossoverPoints[2]])
    copy(child2[crossoverPoints[1]:crossoverPoints[2]], parent2[crossoverPoints[1]:crossoverPoints[2]])
    copy(child3[crossoverPoints[1]:crossoverPoints[2]], parent1[crossoverPoints[1]:crossoverPoints[2]])

    copy(child1[crossoverPoints[2]:], parent1[crossoverPoints[2]:])
    copy(child2[crossoverPoints[2]:], parent2[crossoverPoints[2]:])
    copy(child3[crossoverPoints[2]:], parent3[crossoverPoints[2]:])

    return child1, child2, child3
}



// Метод отбора 
// !(требуется легкая доработка)

func SelectArraysSync(arr1, arr2, arr3 []dataPopulation.DTO) {
	var wg sync.WaitGroup
	errorOccurred := true

	if len(arr1) != len(arr2) || len(arr1) != len(arr3) {
		fmt.Println("Длины массивов не совпадают")
	}

	wg.Add(len(arr1))

	for i := 0; i < len(arr1); i++ {
		go func(i int) {
			defer wg.Done()
			if arr1[i].Subject == arr2[i].Subject || arr1[i].Subject == arr3[i].Subject || arr2[i].Subject == arr3[i].Subject {
				errorOccurred = false
				return
			}
			// Проверка на одинаковые элементы внутри каждого массива
			for j := i + 1; j < len(arr1); j++ {
				// sort.Slice(arr1, func(i, j int) bool { return arr1[i].Subject < arr1[j].Subject })
				switch {
					case arr1[i].Subject == arr1[j].Subject:
						errorOccurred = false
						return
					case arr2[i].Subject == arr2[j].Subject:
						errorOccurred = false
						return
					case arr3[i].Subject == arr3[j].Subject:
						errorOccurred = false
						return
				}
			}	
		}(i)
	} 
	
	wg.Wait()


	if errorOccurred {
		fmt.Println("ok")
		fmt.Println("Родительские массивы:")
		fmt.Println("Array 1:", arr1)
		fmt.Println("Array 2:", arr2)
		fmt.Println("Array 3:", arr3)

		fmt.Println("Дочерние массивы:")
		child1, child2, child3 := KPointCrossover(arr1, arr2, arr3, 2)
		fmt.Println("K-точечный кроссовер:")
		fmt.Println("Child 1:", child1)
		fmt.Println("Child 2:", child2)
		fmt.Println("Child 3:", child3)

		child1_1, child2_1, child3_1 := ThreePointCrossover(arr1, arr2, arr3)
		fmt.Println("Трёхточечный кроссовер:")
		fmt.Println("Child 1:", child1_1)
		fmt.Println("Child 2:", child2_1)
		fmt.Println("Child 3:", child3_1)
	}	
}
