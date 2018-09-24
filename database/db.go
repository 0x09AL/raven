package database

import "database/sql"

var db_filename = "db_file.db"

var db, _ = sql.Open("sqlite3", db_filename)
