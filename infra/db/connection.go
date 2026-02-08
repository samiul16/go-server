package db

import (
	"fmt"
	"go-server/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DbConfig) string {
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", cnf.DB_USER,cnf.DB_PASSWORD, cnf.DB_HOST, cnf.DB_PORT, cnf.DB_NAME)

	if !cnf.SSL_MODE  {
connectionStr += " sslmode=disable" 
	}

	return "user=postgres password=1234 host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection(dbConfig *config.DbConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(dbConfig)
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		return nil, err
	}

 return dbCon, nil
}