// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
)
var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)
func initializeBaz(ctx context.Context) (Baz, error) {
	panic(wire.Build(SuperSet))
	return Baz{}, nil
}