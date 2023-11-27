package db

import (
	"fmt"

	projectConfig "projectname/config"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func New() (*Client, error) {
	// build up the connection string
	// Example: root:@tcp(localhost:3306)/test
	hostString := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s",
		projectConfig.DbUser,
		projectConfig.DbPassword,
		projectConfig.DbHost,
		projectConfig.DbPort,
		projectConfig.DbName,
	)

	// Open new connection
	db, err := sql.Open("mysql", hostString)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.MySQL, db.DB())

	// create the client with our configured driver
	return NewClient(Driver(drv)), nil
}

func NewX() *Client {
	client, err := New()
	if err != nil {
		panic("could not connect to database: " + err.Error())
	}

	return client
}
