package integration

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type MysqlConfig struct {
	RootPassword string
	User         string
	Password     string
	Database     string
}

func RunMysql(config MysqlConfig) (tearDownFn func() error) {
	var db *sql.DB

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	env := getEnv(config)

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306/tcp": {{HostIP: "localhost", HostPort: "3306/tcp"}},
		},
		Env: env,
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	tearDownFn = func() error {
		return pool.Purge(resource)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", "root:root@(localhost:3306)/mysql")
		if err != nil {
			return errors.New("Fail to open connection to mysql instance")
		}

		err = db.Ping()
		if err != nil {
			return errors.New("Fail to ping mysql instance")
		}

		return nil

	}); err != nil {
		tearDownFn()
		log.Fatalf("Could not connect to database: %s", err)
	}

	return tearDownFn
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
