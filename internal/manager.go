package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	sg "github.com/wakeapp/go-sql-generator"
)

type config struct {
	Host     string
	Username string
	Pass     string
	Port     string
	DBName   string
}

// SQLManager - manage connect to db
type SQLManager struct {
	conn *sql.DB
}

// InitManager - init connect to db
func InitManager() *SQLManager {
	var m = &SQLManager{}

	m.open(&config{
		Host:     os.Getenv("MYSQL_HOST"),
		Username: os.Getenv("MYSQL_USER"),
		Pass:     os.Getenv("MYSQL_PASSWORD"),
		Port:     "3306",
		DBName:   os.Getenv("MYSQL_DATABASE"),
	})

	return m
}

// Close - close connect to db
func (m *SQLManager) Close() {
	_ = m.conn.Close()
}

func (m *SQLManager) GetConnection() *sql.DB {
	return m.conn
}

func (m *SQLManager) Insert(dataInsert *sg.InsertData) (int, error) {
	if err := m.conn.Ping(); err != nil {
		return 0, err
	}

	if len(dataInsert.ValuesList) == 0 {
		return 0, nil
	}

	sqlGenerator := sg.MysqlSqlGenerator{}

	query, args, err := sqlGenerator.GetInsertSql(*dataInsert)
	if err != nil {
		return 0, err
	}

	var stmt *sql.Stmt
	stmt, err = m.conn.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	var result sql.Result
	result, err = stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	ra, _ := result.RowsAffected()

	return int(ra), nil
}

func (m *SQLManager) Upsert(dataUpsert *sg.UpsertData) (int, error) {
	if len(dataUpsert.ValuesList) == 0 {
		return 0, nil
	}

	sqlGenerator := sg.MysqlSqlGenerator{}

	query, args, err := sqlGenerator.GetUpsertSql(*dataUpsert)
	if err != nil {
		return 0, fmt.Errorf("on upsert: on Generate: %s", err)
	}

	var stmt *sql.Stmt
	stmt, err = m.conn.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("on upsert: on Prepare: %s, query: %s", err, query)
	}
	defer func() {
		_ = stmt.Close()
	}()

	var result sql.Result
	result, err = stmt.Exec(args...)
	if err != nil {
		return 0, fmt.Errorf("on upsert: on exec: %s", err)
	}

	ra, _ := result.RowsAffected()

	return int(ra), nil
}

func (m *SQLManager) open(c *config) {
	var conn *sql.DB
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?collation=utf8_unicode_ci", c.Username, c.Pass, c.Host, c.Port, c.DBName)
	if conn, err = sql.Open("mysql", dsn); err != nil {
		log.Printf("on insert: on open connection to db: %s \n", err.Error())

		os.Exit(1)
	}

	m.conn = conn
}
