package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	_ = sql.ErrConnDone
}
