package database

import (
	"database/sql"
)

type Database struct {
	conexao  string
	database string
}

func New(conexao, database string) *Database {
	return &Database{
		conexao:  conexao,
		database: database,
	}
}

func (d *Database) Conect() *sql.DB {
	db, err := sql.Open(d.database, d.conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=root dbname=loja password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}