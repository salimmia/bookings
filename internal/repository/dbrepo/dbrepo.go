package dbrepo

import (
	"database/sql"
	"github.com/salimmia/bookings/internal/config"
	"github.com/salimmia/bookings/internal/repository"
)

type MysqlDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewMysqlRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &MysqlDBRepo{
		App: a,
		DB:  conn,
	}
}
