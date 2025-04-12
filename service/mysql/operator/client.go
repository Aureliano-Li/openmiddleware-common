package operator

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

// Manage Mysql Connections

var mysqlDriverName = "mysql"

type Info struct {
	Path     string
	Port     string
	Username string
	Password string
}

type Client struct {
	Mysql Info
	DBs   map[string]*sql.DB
}

func (c *Client) GetConnection(dbName string) (*sql.DB, error) {
	conn, err := c.getConnection(dbName)
	return conn, err
}

func (c *Client) RemoveConnection(dbName string) (*sql.DB, error) {
	conn, err := c.getConnection(dbName)
	return conn, err
}

func (c *Client) getConnection(connectionName string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	db = c.DBs[connectionName]

	if db == nil {
		// If db connection never been created
		source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Mysql.Username, c.Mysql.Password, c.Mysql.Path, c.Mysql.Port, connectionName)
		if connectionName == "default" {
			source = fmt.Sprintf("%s:%s@tcp(%s:%s)/", c.Mysql.Username, c.Mysql.Password, c.Mysql.Path, c.Mysql.Port)
		}
		logrus.Debug(fmt.Sprintf("Source Mysql: %s", source))
		db, err = sql.Open(mysqlDriverName, source)
		if err != nil {
			return nil, err
		}
		// Sets the maximum amount of time a connection may be reused,
		// Expired connections may be closed lazily before reuse.
		db.SetConnMaxLifetime(time.Minute * 3)
		// sets the maximum number of open connections to the database.
		db.SetMaxOpenConns(10)
		// sets the maximum number of connections in the idle
		db.SetMaxIdleConns(5)
		c.DBs[connectionName] = db
		logrus.Debug(fmt.Sprintf("Created mysql: %s:%s ,database: %s connection.", c.Mysql.Path, c.Mysql.Port, connectionName))
	}
	return db, nil
}

func (c *Client) CloseDb() {
	// close all db
	for _, db := range c.DBs {
		err := db.Close()
		if err != nil {
			logrus.Error(err)
		}
	}
}
