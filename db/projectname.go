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
	hostString := projectConfig.DbUser +
		":" +
		projectConfig.DbPassword +
		"@tcp(" +
		projectConfig.DbHost +
		":" +
		fmt.Sprintf("%d", projectConfig.DbPort) +
		")/" +
		projectConfig.DbName

	// Open new connection
	db, err := sql.Open("mysql", hostString)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db.DB())

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
