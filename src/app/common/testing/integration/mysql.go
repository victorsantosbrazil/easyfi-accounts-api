package integration

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

type MysqlConfig struct {
	RootPassword string
	User         string
	Password     string
	Database     string
}

func RunMysql(config MysqlConfig) Container {
	var db *sql.DB

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	env := getEnv(config)

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
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
		port := resource.GetPort("3306/tcp")
		db, err = sql.Open("mysql", fmt.Sprintf("root:root@(localhost:%s)/%s", port, config.Database))
		if err != nil {
			return errors.New("Fail to open connection to mysql instance")
		}

		err = db.Ping()
		if err != nil {
			return errors.New("Fail to ping mysql instance")
		}

		return nil

	})

	if err != nil {
		container.Stop()
		log.Fatalf("Could not connect to database: %s", err)
	}

	return container
}

func getEnv(config MysqlConfig) []string {
	env := []string{
		fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", config.RootPassword),
		fmt.Sprintf("MYSQL_DATABASE=%s", config.Database),
	}

	if config.User != "" {
		env = append(env, fmt.Sprintf("MYSQL_USER=%s", config.User))
	}

	if config.Password != "" {
		env = append(env, fmt.Sprintf("MYSQL_PASSWORD=%s", config.Password))
	}

	return env
}
