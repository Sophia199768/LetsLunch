// Code generated by DIGEN; DO NOT EDIT.
// This file was generated by Dependency Injection Container Generator 0.1.0 (built at 2023-10-21T19:51:59Z).
// See docs at https://github.com/strider2038/digen

package internal

import (
	"context"
)

type Container struct {
	err error
}

func NewContainer() *Container {
	c := &Container{}

	return c
}

// Error returns the first initialization error, which can be set via SetError in a service definition.
func (c *Container) Error() error {
	return c.err
}

// SetError sets the first error into container. The error is used in the public container to return an initialization error.
func (c *Container) SetError(err error) {
	if err != nil && c.err == nil {
		c.err = err
	}
}

func (c *Container) Close() {}
