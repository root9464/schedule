package modules

import (
	dataPopulation "root/Gen"
	"sort"
)

func PermutationsCount(arr1, arr2, arr3 []dataPopulation.DTO) int {
	count := 0

	// Создаем слайсы с уникальными предметами для каждого массива
	subjects1 := getUniqueSubjects(arr1)
	subjects2 := getUniqueSubjects(arr2)
	subjects3 := getUniqueSubjects(arr3)

	// Сортируем слайсы предметов
	sort.Strings(subjects1)
	sort.Strings(subjects2)
	sort.Strings(subjects3)

	// Перебираем все возможные комбинации предметов
	for _, subject1 := range subjects1 {
		for _, subject2 := range subjects2 {
			for _, subject3 := range subjects3 {
				// Проверяем, что предметы на соответствующих индексах не совпадают
				if subject1 != subject2 && subject1 != subject3 && subject2 != subject3 {
					count++
				}
			}
		}
	}

	return count
}

func getUniqueSubjects(arr []dataPopulation.DTO) []string {
	// Создаем мапу для хранения уникальных предметов
	subjectsMap := make(map[string]bool)

	// Перебираем массив и добавляем уникальные предметы в мапу
	for _, student := range arr {
		subjectsMap[student.Subject] = true
	}

	// Создаем слайс с уникальными предметами из мапы
	var uniqueSubjects []string
	for subject := range subjectsMap {
		uniqueSubjects = append(uniqueSubjects, subject)
	}

	return uniqueSubjects
}