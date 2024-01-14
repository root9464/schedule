package main

import (
	"fmt"
	"io"
	"log"
	dataPopulation "root/Gen"
	controllers "root/database/controllers"
)

//массив предметов для 1 группы (можно переиспользовать)

var web_group = []dataPopulation.DTO{
	dataPopulation.Data.Math,
	dataPopulation.Data.Rus,
	dataPopulation.Data.Inf,
	dataPopulation.Data.Fiz,
	dataPopulation.Data.Hist,
	dataPopulation.Data.Lit,
	dataPopulation.Data.Mdk,
	dataPopulation.Data.Arch,
	dataPopulation.Data.Soc,
	dataPopulation.Data.Obg,
	dataPopulation.Data.Fizr,
	dataPopulation.Data.Angl,
	dataPopulation.Data.UIPD,
	dataPopulation.Data.Geo,
}

func main() {
	log.SetOutput(io.Discard)
	db := controllers.NewPostgresDB()
	err := controllers.Database.Connect(db)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных. \n", err)
	}

	//цикл добавления записей в таблицу + автоматическое создание самой таблицы
	/* for i := range web_group {
		addData(db, &all[i])
	}*/
	 

	//цикл чтения записей из таблицы
	for i := 1; i <= len(web_group); i++ {
		data, _ := controllers.Database.GetDataFromTableById(db, "1wb2", i)
		fmt.Println(*data)
	}
	
	
}


