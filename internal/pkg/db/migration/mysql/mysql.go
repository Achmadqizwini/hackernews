package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Db *sql.DB

func InitDB() {
	// format --> <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	var connectionString = os.Getenv("DB_Connection")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection", err.Error())
	}
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db", errPing.Error())
	} else {
		fmt.Println("koneksi berhasil")
	}
	Db = db

}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migration/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
