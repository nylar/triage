package triage

import (
	"log"
	"os"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
	"github.com/nylar/triage/config"
	"gopkg.in/ory-am/dockertest.v3"
)

var db *sqlx.DB

func setUp(t *testing.T) func() error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		t.Fatalf("Couldn't open migration driver: %v", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		t.Fatalf("Couldn't start migrations: %v", err)
	}
	if err := migrations.Up(); err != nil {
		t.Fatalf("Couldn't run migrations: %v", err)
	}

	return func() error {
		return migrations.Down()
	}
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Couldn't connect to Docker: %v", err)
	}

	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Couldn't start resource: %v", err)
	}

	if err := pool.Retry(func() error {
		port, err := strconv.Atoi(resource.GetPort("3306/tcp"))
		if err != nil {
			return err
		}

		sqlConfig := &config.SQL{
			Hostname: "localhost",
			Port:     port,
			Username: "root",
			Password: "secret",
			Database: "mysql",
		}

		db, err = sqlx.Open("mysql", sqlConfig.DataSourceName())
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("Couldn't connect to Docker: %v", err)
	}

	code := m.Run()

	if err := db.Close(); err != nil {
		log.Fatalf("Couldn't close DB connection: %v", err)
	}

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Couldn't purge resource: %v", err)
	}

	os.Exit(code)
}
