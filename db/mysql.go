package db

import (
	"fmt"
	"log"

	"github.com/Ashis-007/hms/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(env config.Env) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("host=%s port=%s user=%s password=% dbname=%s", env.DBHost, env.DBPort, env.DBUser, env.DBPassword, env.DBName))
	if err != nil {
		log.Fatal("failed to connect to mysql db: ", err)
		return db, err
	}

	return db, nil
}
