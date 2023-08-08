package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/dronestock/flutter/internal/internal"
	"github.com/goexl/gox/args"
)

type Build struct {
	drone.Base

	source string
	typ    internal.Type
}

func NewBuild(base drone.Base, source string, typ internal.Type) *Build {
	return &Build{
		Base: base,

		source: source,
		typ:    typ,
	}
}

func (b *Build) Runnable() bool {
	return true
}

func (b *Build) Run(ctx context.Context) (err error) {
	target := "apk"
	switch b.typ {
	case internal.TypeAndroid:
		target = "apk"
	}
	_args := args.New().Build().Subcommand("build", target)
	if b.Verbose {
		_args.Flag("verbose")
	}
	_, err = b.Command(internal.CommandFlutter).Args(_args.Build()).Dir(b.source).Context(ctx).Build().Exec()

	return
}
