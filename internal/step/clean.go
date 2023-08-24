package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/dronestock/flutter/internal/internal"
	"github.com/goexl/gox/args"
)

type Clean struct {
	drone.Base

	source string
}

func NewClean(base drone.Base, source string) *Clean {
	return &Clean{
		Base: base,

		source: source,
	}
}

func (b *Clean) Runnable() bool {
	return true
}

func (b *Clean) Run(ctx context.Context) (err error) {
	_args := args.New().Build().Subcommand("clean")
	if b.Verbose {
		_args.Flag("verbose")
	}
	_, err = b.Command(internal.CommandFlutter).Args(_args.Build()).Dir(b.source).Context(ctx).Build().Exec()

	return
}
