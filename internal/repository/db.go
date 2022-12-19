package repository

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbLock = &sync.Mutex{}
var db *sqlx.DB

func GetDB() *sqlx.DB {
	if db == nil {
		dbLock.Lock()
		defer dbLock.Unlock()

		if db == nil {
			var err error
			db, err = sqlx.Connect(
				"postgres",
				"password=postgres user=postgres dbname=postgres sslmode=disable search_path=kontora_schema",
			)
			if err != nil {
				log.Fatalf(fmt.Sprintf("Unable to connect to database: %v\n", err))
			}

			return db
		} else {
			return db
		}
	} else {
		return db
	}
}
