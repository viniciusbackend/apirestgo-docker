package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringConexao := "host=localhost user=vinicius password=12 dbname=produtos port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))

	if err != nil {
		log.Panic("Erro ao conecar com banco de dados")
	}
}
