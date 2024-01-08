package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type DTO struct {
	Teacher
	subject string
	hour    int16
}
type Teacher struct{
	name string
}

func main() {
	subjects := []string{
		"Математика", "Русский", "Информатика", "Физика", 
		"История", "Литература", "МДК 08.02", "Архитектура АП",
		"Общество", "ОБЖ", "Физра", "Англ. яз.", "УИПД", "География"}
	teachers := []string{
		"Трушина И.Ю", "Пахилько О.Н", "Морозова М.В", "Имашева К.Б",
		"Швиндина Н.А", "Зернова Е.А", "Имаев М.А", "Нургалиева И.Ю",
		"Чебрукова Т.А", "Новиков Р.Е", "Чернышова Е.А", "Силищева О.И",
		"Новикова М.Н"}

	math := DTO{
		Teacher: Teacher{
			name: teachers[0],
		},
		subject: subjects[0],
		hour:    10,
	}
	rus := DTO{
		Teacher: Teacher{
			name: teachers[1],
		},
		subject: subjects[1],
		hour:    11,
	}

	inf:= DTO{
		Teacher: Teacher{
			name: teachers[2],
		},
		subject: subjects[2],
		hour:    12,
	}
	fiz:= DTO{
		Teacher: Teacher{
			name: teachers[3],
		},
		subject: subjects[3],
		hour:    13,
	}
	hist:= DTO{
		Teacher: Teacher{
			name: teachers[4],
		},
		subject: subjects[4],
		hour:    14,
	}
	lit:= DTO{
		Teacher: Teacher{
			name: teachers[1],
		},
		subject: subjects[5],
		hour:    15,
	}
	mdk:= DTO{
		Teacher: Teacher{
			name: teachers[5],
		},
		subject: subjects[6],
		hour:    16,
	}
	arch:= DTO{
		Teacher: Teacher{
			name: teachers[6],
		},
		subject: subjects[7],
		hour:    17,
	}
	soc:= DTO{
		Teacher: Teacher{
			name: teachers[7],
		},
		subject: subjects[8],
		hour:    18,
	}
	obg:= DTO{
		Teacher: Teacher{
			name: teachers[8],
		},
		subject: subjects[9],
		hour:    19,
	}
	fizra:= DTO{
		Teacher: Teacher{
			name: teachers[9],
		},
		subject: subjects[10],
		hour:    20,
	}
	ang:= DTO{
		Teacher: Teacher{
			name: teachers[10],
		},
		subject: subjects[11],
		hour:    21,
	}
	uipd:= DTO{
		Teacher: Teacher{
			name: teachers[11],
		},
		subject: subjects[12],
		hour:    22,
	}
	geo := DTO{
		Teacher: Teacher{
			name: teachers[12],
		},
		subject: subjects[13],
		hour:    23,
	}


	
	
	
	fmt.Println(data)
}

switch {
	case reflect.DeepEqual(arr1, arr2) && reflect.DeepEqual(arr1, arr3) && reflect.DeepEqual(arr2, arr3): 
		fmt.Println("Ошибка: все элементы повторяются")
		errorOccurred = true

	case reflect.DeepEqual(arr1, arr2):
		fmt.Println("Ошибка: одинаковые элементы в arr1 и arr2")
		errorOccurred = true

	case reflect.DeepEqual(arr1, arr3):
		fmt.Println("Ошибка: одинаковые элементы в arr1 и arr3")
		errorOccurred = true

	case reflect.DeepEqual(arr2, arr3):
		fmt.Println("Ошибка: одинаковые элементы в arr2 и arr3")
		errorOccurred = true

	case arr1[0] == arr1[i] || arr1[1] == arr1[i] || arr1[2] == arr1[i]:
		fmt.Println("Ошибка: какой либо элемент повторяется на 0 1 2ом индексе в массиве arr1")
		errorOccurred = true
	case arr2[0] == arr2[i] || arr2[1] == arr2[i] || arr2[2] == arr2[i]:
		fmt.Println("Ошибка: какой либо элемент повторяется на 0 1 2ом индексе в массиве arr2")
		errorOccurred = true
	case arr3[0] == arr3[i] || arr3[1] == arr3[i] || arr3[2] == arr3[i]:
		fmt.Println("Ошибка: какой либо элемент повторяется на 0 1 2ом индексе в массиве arr3")
		errorOccurred = true
}
switch {
	case arr1[0] == arr2[0] || arr1[0] == arr3[0] || 
		 arr1[1] == arr2[1] || arr1[1] == arr3[1] ||
		 arr1[2] == arr2[2] || arr1[2] == arr3[2]:
		fmt.Println("Ошибка: одинаковые предметы на 0 1 2ом индексе в массиве arr1 и arr2 или arr1 и arr3")  
		errorOccurred = true
	case arr2[0] == arr1[0] || arr2[0] == arr3[0] || 
		 arr2[1] == arr1[1] || arr2[1] == arr3[1] ||
		 arr3[2] == arr1[2] || arr2[2] == arr3[2]:
		fmt.Println("Ошибка: одинаковые предметы на 0 1 2ом индексе в массиве arr1 и arr2 или arr1 и arr3")
		errorOccurred = true
}

switch {
	case arr1[0] == arr1[i] || arr1[1] == arr1[i] || arr1[2] == arr1[i]:
		fmt.Println("Ошибка: какой либо элемент повторяется на 0 1 2ом индексе в массиве arr1")
		errorOccurred = true
	case arr2[0] == arr2[i] || arr2[1] == arr2[i] || arr2[2] == arr2[i]:
		fmt.Println("Ошибка: какой либо элемент повторяется на 0 1 2ом индексе в массиве arr2")
		errorOccurred = true
	case arr3[0] == arr3[i] || arr3[1] == arr3[i] || arr3[2] == arr3[i]:
		fmt.Println("Ошибка: какой либо элемент повторяется на 0 1 2ом индексе в массиве arr3")
		errorOccurred = true
}
