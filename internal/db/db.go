package db

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/myl7/bibak/internal/config"
)

func GetDB() (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", config.Cfg.DBPath)
}

//go:embed schema.sql
var schema string

func InitDB(db *sqlx.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}

var DB *sqlx.DB

func LoadDB() error {
	if DB != nil {
		return nil
	}

	var err error
	DB, err = GetDB()
	if err != nil {
		return err
	}

	err = InitDB(DB)
	if err != nil {
		return err
	}

	return nil
}
