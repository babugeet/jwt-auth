package db

import (
	"jwt-auth/mocks"
	"jwt-auth/models"
	"jwt-auth/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MySQLdb struct {
	db *gorm.DB
}

type Table struct {
	Username  string `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	Password  string
	Gender    string
	Age       int64
	Height    int64
	Weight    int64
	BMI       int64
}

// Custom table name method
func (Table) TableName() string {
	return "user_database" // Custom table name
}

func NewMySQLdb() models.Database {
	dsn := "host=localhost user=postgres password=Abhi@1234 dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Table{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}

	return &MySQLdb{db}
}

func (m MySQLdb) GetUser(username string) (bool, error) {
	_, ok := utils.Struct2Map()[username]
	return ok, nil
}

func (m MySQLdb) AddUser(user models.User) error {
	mocks.Users = append(mocks.Users, user)
	return nil
}
