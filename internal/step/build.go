package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/dronestock/flutter/internal/internal/constant"
	"github.com/goexl/gox/args"
)

type Build struct {
	drone.Base

	binary string
	source string
	typ    constant.Type
}

func NewBuild(base drone.Base, binary string, source string, typ constant.Type) *Build {
	return &Build{
		Base: base,

		binary: binary,
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
	case constant.TypeAndroid:
		target = "apk"
	case constant.TypeWeb:
		target = "web"
	}
	_args := args.New().Build().Subcommand("build", target)
	if b.Verbose {
		_args.Flag("verbose")
	}
	_, err = b.Command(b.binary).Args(_args.Build()).Dir(b.source).Context(ctx).Build().Exec()

	return
}
