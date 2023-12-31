package methods

import (
	dataPopulation "root/Gen"
	"sort"
)

// Метод ранжирования
func RankingMethod(arr *[]dataPopulation.DTO) {
    // Присваиваем ранги на основе качества каждого элемента
    ranks := make([]int, len(*arr))
    for i := 0; i < len(*arr); i++ {
        ranks[i] = 1
        for j := 0; j < len(*arr); j++ {
            if j != i && (*arr)[j].Quality < (*arr)[i].Quality {
                ranks[i]++
            }
        }
    }

    // Сортируем элементы на основе их рангов
    sortedArr := make([]dataPopulation.DTO, len(*arr))
    copy(sortedArr, *arr)
    sort.Slice(sortedArr, func(i, j int) bool {
        return ranks[i] < ranks[j]
    })

    // Заменяем исходный массив отсортированным массивом
    *arr = sortedArr
}