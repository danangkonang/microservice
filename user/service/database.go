package service

import "database/sql"

type Database struct {
	Postgresql *sql.DB
}
