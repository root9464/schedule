package core

import (
	"fmt"
	data "root/Core/data"
	"sync"
)

func Core() {
	var wg sync.WaitGroup
	multiplecall(data.Group1wb1, data.Group1wb2, data.Group1wb3, &wg)
	
	fmt.Print("\n","\033[31m","Массивы первой группы", "\033[97m", "\n")
	fmt.Print(data.Group1wb1, "\n")
	fmt.Print(data.Group1wb2, "\n")
	fmt.Print(data.Group1wb3, "\n")

	fmt.Print("\n","\033[31m","Массивы второй группы", "\033[97m", "\n")
	fmt.Print(data.Group2wb1, "\n")
	fmt.Print(data.Group2wb2, "\n")
	fmt.Print(data.Group2wb3, "\n")
}


