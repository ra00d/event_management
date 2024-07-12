package configs

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver)
	"github.com/jmoiron/sqlx"
)

var (
	AppDB  *sqlx.DB
	Ql     *QueryLogger
	logger *log.Logger
)

func DataBasaInit() /* *sqlx.DB  */ {
	db, err := sqlx.Connect("mysql", "root:@(127.0.0.1:3306)/events_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	AppDB = db
	logger = new(log.Logger)
	Ql = &QueryLogger{db, logger}
	// return db
	// defer db.Close()
}

type QueryLogger struct {
	queryer sqlx.Queryer
	logger  *log.Logger
}

func (p *QueryLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	log.Printf(strings.ReplaceAll(query, "?", "%v"), args...)
	return p.queryer.Query(query, args...)
}

func (p *QueryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	log.Printf(strings.ReplaceAll(query, "?", "%v"), args...)

	return p.queryer.Queryx(query, args...)
}

func (p *QueryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	log.Printf(strings.ReplaceAll(query, "?", "%v"), args...)

	return p.queryer.QueryRowx(query, args...)
}
