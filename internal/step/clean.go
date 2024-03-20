package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/goexl/gox/args"
)

type Clean struct {
	drone.Base

	binary string
	source string
}

func NewClean(base drone.Base, binary string, source string) *Clean {
	return &Clean{
		Base: base,

		binary: binary,
		source: source,
	}
}

func (c *Clean) Runnable() bool {
	return true
}

func (c *Clean) Run(ctx context.Context) (err error) {
	_args := args.New().Build().Subcommand("clean")
	if c.Verbose {
		_args.Flag("verbose")
	}
	_, err = c.Command(c.binary).Args(_args.Build()).Dir(c.source).Context(ctx).Build().Exec()

	return
}
