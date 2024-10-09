package integration

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
)

type PostgresConfig struct {
	User     string
	Password string
	Database string
}

func RunPostgres(config PostgresConfig) Container {
	var db *sql.DB

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	env := getPostgresEnv(config)

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16-bullseye",
		Env:        env,
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	container := Container{
		pool:     pool,
		resource: resource,
	}

	// exponential backoff-retry, because the application in the container might not be ready
	// to accept connections yet
	err = pool.Retry(func() error {
		var err error
		port := resource.GetPort("5432/tcp")
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", config.User, config.Password, port, config.Database))
		if err != nil {
			return errors.New("Fail to open connection to postgres instance")
		}

		err = db.Ping()
		if err != nil {
			return errors.New("Fail to ping postgres instance")
		}

		return nil
	})

	if err != nil {
		container.Stop()
		log.Fatalf("Could not connect to database: %s", err)
	}

	return container
}

func getPostgresEnv(config PostgresConfig) []string {
	return []string{
		fmt.Sprintf("POSTGRES_USER=%s", config.User),
		fmt.Sprintf("POSTGRES_PASSWORD=%s", config.Password),
		fmt.Sprintf("POSTGRES_DB=%s", config.Database),
	}
}
