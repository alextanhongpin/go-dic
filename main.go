// +build wireinject

package main

import (
	"context"

	"dic/foo"

	"github.com/google/go-cloud/wire"
)

func initializeBaz(ctx context.Context) (foo.Baz, error) {
	wire.Build(foo.SuperSet)
	return foo.Baz{}, nil
}
