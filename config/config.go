package config

import (
	"fmt"
)

const (
	DBUser     = "postgres"
	DBPassword = "nopass123"
	DBName     = "dvd_rentals"
	DBHost     = "localhost"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}
