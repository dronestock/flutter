package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/goexl/gox/args"
)

type Get struct {
	drone.Base

	binary string
	source string
}

func NewGet(base drone.Base, binary string, source string) *Get {
	return &Get{
		Base: base,

		binary: binary,
		source: source,
	}
}

func (g *Get) Runnable() bool {
	return true
}

func (g *Get) Run(ctx context.Context) (err error) {
	_args := args.New().Build().Subcommand("pub", "get").Build()
	_, err = g.Command(g.binary).Args(_args).Dir(g.source).Context(ctx).Build().Exec()

	return
}
