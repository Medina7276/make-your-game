package model

import "database/sql"

type PlayerDB struct {
	*sql.DB
}
