package domain

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type BookRepositoryDb struct {
	client *sqlx.DB
}

func (d BookRepositoryDb) FindAll() ([]Book, error){
	var err error

	books := make([]Book, 0)
	
	findAllsql := "SELECT id, name, read_date, report, rating, img_url FROM book"

	err = d.client.Select(&books, findAllsql)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return books, nil

}

func NewBookRepositoryDb(dbClient *sqlx.DB) BookRepositoryDb {
	return BookRepositoryDb{dbClient}
}