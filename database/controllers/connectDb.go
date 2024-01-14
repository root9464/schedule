package controllers

import (
	"log"
	"os"
	dataPopulation "root/Gen"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	"gorm.io/gorm/logger"
)

type Database interface {
	Connect() error
	AddData(data *dataPopulation.DTO) error
	GetDataFromTableById(Tables string, Id int) (*dataPopulation.DTO, error)
}

type Data struct {
	DB *gorm.DB
}

type PostgresDB struct {
	Data
}

var slowLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold: 1 * time.Microsecond,   
		LogLevel: logger.Silent, 
		Colorful: true,
	},
)

func (p *PostgresDB) Connect() error {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Ошибка чтения env файла. \n", err)
	}

	USER_NAME := os.Getenv("USER_NAME")
	PASSWORD_DB := os.Getenv("PASSWORD_DB")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DB_NAME := os.Getenv("DB_NAME")
	

	db, err := gorm.Open(postgres.Open(
		"host=" + HOST + " user=" + USER_NAME + " password=" + PASSWORD_DB + " dbname=" + DB_NAME + " port=" + PORT,
	), &gorm.Config{
		Logger: slowLogger,
	})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных. \n", err)
	}
	//log.Println("Подключение к базе данных выполнено успешно.")


	//db.Logger = logger.Default.LogMode(logger.Info)
	p.DB = db
	
	return nil
}

func (p *PostgresDB) AddData(data *dataPopulation.DTO) error {
	result := p.DB.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PostgresDB) GetDataFromTableById(Tables string, Id int) (*dataPopulation.DTO, error) {
	var data dataPopulation.DTO
	err := p.DB.Table(Tables).First(&data, Id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func NewPostgresDB() Database {
	return &PostgresDB{}
}