package integration

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ory/dockertest/v3"
)

type (
	LocalstackConfig struct {
		Aws LocalStackAwsConfig
	}

	LocalStackAwsConfig struct {
		DefaultRegion string
		AccessKey     string
		SecretKey     string
	}
)

func RunLocalstack(config LocalstackConfig) *Container {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "localstack/localstack",
		Env: []string{
			"DOCKER_HOST=unix:///var/run/docker.sock",
			fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", config.Aws.AccessKey),
			fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", config.Aws.SecretKey),
			fmt.Sprintf("AWS_DEFAULT_REGION=%s", config.Aws.DefaultRegion),
		},
		Mounts: []string{
			"/var/run/docker.sock:/var/run/docker.sock",
		},
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	container := &Container{
		pool:     pool,
		resource: resource,
	}

	err = pool.Retry(func() error {
		_, err = http.Get("http://localhost:4566/_localstack/health")
		return err
	})

	if err != nil {
		container.Stop()
		log.Fatalf("Could not connect to localstack: %s", err)
	}

	return container
}
