package core

import (
	data "root/Core/data"
	"sync"
)

func Core() {
	var wg sync.WaitGroup
	multiplecall(data.Group1wb1, data.Group1wb2, data.Group1wb3, &wg)

}


