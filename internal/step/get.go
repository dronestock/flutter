package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/dronestock/flutter/internal/internal"
	"github.com/goexl/gox/args"
)

type Get struct {
	drone.Base

	source string
}

func NewGet(base drone.Base, source string) *Get {
	return &Get{
		Base: base,

		source: source,
	}
}

func (g *Get) Runnable() bool {
	return true
}

func (g *Get) Run(ctx context.Context) (err error) {
	_args := args.New().Build().Subcommand("pub", "get").Build()
	_, err = g.Command(internal.CommandFlutter).Args(_args).Dir(g.source).Context(ctx).Build().Exec()

	return
}
