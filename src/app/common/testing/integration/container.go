package integration

import (
	"github.com/ory/dockertest/v3"
)

type Container struct {
	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func (c *Container) Stop() error {
	return c.pool.Purge(c.resource)
}

func (c *Container) GetPort(id string) string {
	return c.resource.GetPort(id)
}
