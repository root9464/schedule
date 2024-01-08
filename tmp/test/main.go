package main 

import "fmt"

func crossover(parent1, parent2 []string, names []string) map[string][]string {
    arrays := make(map[string][]string)
    for i := 0; i < len(parent1); i++ {
        array := make([]string, len(parent1))
        for j := 0; j <= i; j++ {
            array[j] = parent1[j]
        }

        for j := i + 1; j < len(parent1); j++ {
            array[j] = parent2[j]
        }
        arrays[names[i]] = array
    }

    return arrays
}

func main() {
    parent1 := []string{"a", "b", "c", "d"}
    parent2 := []string{"e", "f", "g", "h"}
    names := []string{"arr1", "arr2", "arr3", "arr4"}

    arrays := crossover(parent1, parent2, names)

    fmt.Println("Array 1:", arrays["arr1"])
    fmt.Println("Array 2:", arrays["arr2"])
    fmt.Println("Array 3:", arrays["arr3"])
    fmt.Println("Array 4:", arrays["arr4"])
}