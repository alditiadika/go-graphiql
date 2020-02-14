package database

import (
	"database/sql"
	"fmt"

	//lib/pq
	_ "github.com/lib/pq"

	"github.com/alditiadika/go-graphiql/app/config"
)

//RunQuery func
func RunQuery(q string) (*sql.Rows, error) {
	conf := config.Config
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error when calling database!", err)
		return &sql.Rows{}, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error when ping database!", err)
		return &sql.Rows{}, err
	}
	defer db.Close()
	rows, err := db.Query(q)
	return rows, err
}
