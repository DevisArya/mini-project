package config

import (
	"fmt"
	"log"
	"miniproject/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	DB *gorm.DB
)

func Init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	Username string
	Password string
	Port     string
	Host     string
	Name     string
}

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	config := Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		&models.Customer{},
		&models.Admin{},
		&models.Area{},
		&models.Cleaner{},
		&models.Payment{},
		&models.ServiceType{},
		&models.Store{},
		&models.Team{},
		&models.TransactionDetail{},
		&models.Transaction{},
		&models.Chart{},
	)
}
