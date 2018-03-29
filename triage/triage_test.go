package triage

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/nylar/triage/config"
	"gopkg.in/ory-am/dockertest.v3"
)

var db *sqlx.DB

func rootPath() string {
	return filepath.Join(os.Getenv("GOPATH"), "src/github.com/nylar/triage")
}

func setUp(t *testing.T) func() error {
	t.Logf("Setup for %s", t.Name())

	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		t.Fatalf("Couldn't open migration driver: %v", err)
	}
	defer driver.Unlock()

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://"+filepath.Join(rootPath(), "migrations"),
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
		t.Logf("Tearing down for %s", t.Name())
		return migrations.Down()
	}
}

func loadFixtures(t *testing.T) {
	if db == nil {
		panic("DB must be initialised")
	}

	fixturesPath := filepath.Join(rootPath(), "test_fixtures.sql")

	fixtures, err := ioutil.ReadFile(fixturesPath)
	if err != nil {
		t.Fatalf("Couldn't read fixtures file: %v", err)
	}

	_, err = db.Exec(string(fixtures))
	if err != nil {
		t.Fatalf("Couldn't run fixtures: %v", err)
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
			Params: map[string]string{
				"multiStatements": "true",
				"parseTime":       "true",
			},
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
