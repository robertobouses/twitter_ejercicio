package internal

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	User     string
	Pass     string
	Host     string
	Port     string
	Database string
}

func NewPostgres(c DBConfig) (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Pass, c.Database, c.Port)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Println("Error al abrir la conexión:", err)
		return nil, err
	}
	log.Println("La conexión se abrió correctamente")
	return db, nil
}
