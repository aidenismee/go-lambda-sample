package db

import "github.com/aidenismee/go-lambda-sample/pkg/config"

type DB interface {
	GetClient() any
}

type db struct {
	client any
}

func NewDB(config config.DatabaseConfig) DB {
	client := "DB"

	return &db{
		client: client,
	}
}

func (d *db) GetClient() any {
	return d.client
}
