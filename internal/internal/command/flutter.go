package command

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/goexl/args"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type Flutter struct {
	base   *drone.Base
	binary string
	source string
	_      gox.CannotCopy
}

func NewFlutter(base *drone.Base, binary string, source string) *Flutter {
	return &Flutter{
		base:   base,
		binary: binary,
		source: source,
	}
}

func (f *Flutter) Exec(ctx *context.Context, arguments *args.Arguments) (err error) {
	if f.base.Verbose {
		arguments = arguments.Rebuild().Flag("verbose").Build()
	}

	fields := gox.Fields[any]{
		field.New("binary", f.binary),
		field.New("arguments", arguments.Cli()),
	}
	f.base.Debug("命令执行开始", fields...)
	if _, err = f.base.Command(f.binary).Arguments(arguments).Dir(f.source).Context(*ctx).Build().Exec(); nil != err {
		f.base.Warn("命令执行出错", fields.Add(field.Error(err))...)
	} else {
		f.base.Debug("命令执行成功", fields...)
	}

	return
}

func (f *Flutter) Remove(name string, path string) {
	f.base.Cleanup().File(path).Name(name).Build()
}
