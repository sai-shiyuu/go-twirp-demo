package db

import (
	"database/sql"
	"errors"
	"fmt"

	"go-web/common/config"
	"go-web/common/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = log.GetLogger()
}

func getDB() (*sql.DB, error) {
	dbcon := config.Configs.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local",
		dbcon.Username,
		dbcon.Password,
		dbcon.Host,
		dbcon.Port,
		dbcon.Database)
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Error(err)
		return db, errors.New("open db error")
	}
	return db, err
}

func QueryRow(sql string, args ...interface{}) (*sql.Row, error) {
	db, err := getDB()
	if err != nil {
		logger.Error(err)
		return nil, errors.New("get db error")
	}
	defer db.Close()
	row := db.QueryRow(sql, args...)
	return row, err
}

func QueryRows(sql string, args ...interface{}) (*sql.Rows, error) {
	db, err := getDB()
	if err != nil {
		logger.Error(err)
		return nil, errors.New("get db error")
	}
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		logger.Error(err)
		return nil, errors.New("prepare error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	return rows, err
}

func Exec(sql string, args ...interface{}) (sql.Result, error) {
	db, err := getDB()
	if err != nil {
		logger.Error(err)
		return nil, errors.New("get db error")
	}
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		logger.Error(err)
		return nil, errors.New("prepare error")
	}
	defer stmt.Close()
	result, err := stmt.Exec(args...)
	return result, err
}
