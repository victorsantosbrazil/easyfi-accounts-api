package integration

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
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

func RunLocalstack(config LocalstackConfig) (tearDownFn func() error) {
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
		PortBindings: map[docker.Port][]docker.PortBinding{
			"4566/tcp": {{HostIP: "localhost", HostPort: "4566/tcp"}},
		},
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

	tearDownFn = func() error {
		return pool.Purge(resource)
	}

	err = pool.Retry(func() error {
		_, err = http.Get("http://localhost:4566/_localstack/health")
		return err
	})

	if err != nil {
		tearDownFn()
		log.Fatalf("Could not connect to localstack: %s", err)
	}

	return tearDownFn
}
