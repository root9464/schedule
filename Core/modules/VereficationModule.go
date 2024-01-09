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

func CompareArrays(arrays ...[]dataPopulation.DTO) (<-chan error, bool) {
	errCh := make(chan error)
	flag := false
	go func() {
		defer close(errCh)

		// Находим массив с максимальной длиной
		maxLength := maxlen(arrays...)
		

		// Сравниваем элементы на одинаковых индексах
		for i := 0; i < maxLength; i++ {
			for j := 1; j < len(arrays); j++ {
				if i < len(arrays[j]) && arrays[0][i] == arrays[j][i] {
					errCh <- errors.New("ошибка: элементы на одинаковых индексах совпадают")
					flag = false
					return
				}else{
					flag = true
				}
				// if i != j && len(arrays[i]) > 0 && len(arrays[j]) > 1 && arrays[i][0] == arrays[j][1] {
				// 	errCh <- errors.New("ошибка: элемент на первом индексе совпадает с элементом на втором индексе другого массива")
				// 	return
				// }

				
			}
		}
	}()


	return errCh, flag
}
