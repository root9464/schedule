package modules

import (
	"errors"
	"fmt"
	dataPopulation "root/Gen"
	//"sync"
)

func countmodule(arr ...[]dataPopulation.DTO) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		count++
	}
	return count
}

func Log(arrays ...[]dataPopulation.DTO) {
	count := countmodule(arrays...)
    fmt.Println(count)

    for i := 0; i < count; i++ {
        fmt.Println("Массив", arrays[i])
        fmt.Println("Количество элементов", len(arrays[i]))

        for j := 0; j < len(arrays[i]); j++ {
            fmt.Println("Элемент", j+1, arrays[i][j])
        }
    }
}


func maxlen(arrays ...[]dataPopulation.DTO) int {
	max := 0
	for i := 0; i < len(arrays); i++ {
		if len(arrays[i]) > max {
			max = len(arrays[i])
		}
	}
	return max

}

func minlen(arrays ...[]dataPopulation.DTO) int {
	if len(arrays) == 0 {
		return 0
	}

	minLength := len(arrays[0])
	for _, arr := range arrays {
		if len(arr) < minLength {
			minLength = len(arr)
		}
	}

	return minLength
}
func CompareArrays(arrays ...[]dataPopulation.DTO) (<-chan error, bool) {
	errCh := make(chan error)
	flag := false
	go func() {
		defer close(errCh)

		// Находим массив с максимальной длиной
		maxLength := maxlen(arrays...)
		minLength := minlen(arrays...)
		fmt.Println("Максимальная длина", maxLength)

		// условие проверки и готово
	}()


	return errCh, flag
}
