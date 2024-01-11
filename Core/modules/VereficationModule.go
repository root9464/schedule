package modules

import (
	//"errors"
	"errors"
	"fmt"
	"reflect"
	dataPopulation "root/Gen"
	"sync"
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
func compareMapValues(m map[string][]dataPopulation.DTO) error {
    var wg sync.WaitGroup
    errCh := make(chan error, len(m))

    for _, values := range m {
        wg.Add(1)
        go func(v []dataPopulation.DTO) {
            defer wg.Done()
            for i := 0; i < len(v)-1; i++ {
                if reflect.DeepEqual(v[i].Teacher, v[i+1].Teacher) {
                    errCh <- errors.New("ошибка: содержание ключа равно другому содержанию ключа")
                    return
                }
            }
        }(values)
		
    }

    go func() {
        wg.Wait()
        close(errCh)
    }()

    for err := range errCh {
        if err != nil {
            return err
        }
    }

    return nil
}

func SliceCompare(arrays ...[]dataPopulation.DTO) (map[string][]dataPopulation.DTO, bool) {
    minLength := minlen(arrays...)
    state := make(map[string][]dataPopulation.DTO, 3)
	errorOccurred := false // Произошла ошибка "ложь"

    for i := range arrays {
        arrays[i] = arrays[i][:minLength]
        subject := fmt.Sprintf("Group%d", i+1)
        state[subject] = arrays[i]
    }
    err := compareMapValues(state)
    if err != nil {
        fmt.Println(err)
		errorOccurred = false
		
    } else {
        fmt.Println("Сравнение успешно выполнено")
		errorOccurred = true
		return state, errorOccurred
    }
    return map[string][]dataPopulation.DTO{}, errorOccurred
}
