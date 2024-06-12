package step

import (
	"context"
	"os"
	"path/filepath"

	"github.com/dronestock/flutter/internal/internal/command"
	"github.com/dronestock/flutter/internal/internal/constant"
	"github.com/goexl/args"
)

type Build struct {
	flutter *command.Flutter
	typ     constant.Type
}

func NewBuild(flutter *command.Flutter, typ constant.Type) *Build {
	return &Build{
		flutter: flutter,
		typ:     typ,
	}
}

func (b *Build) Runnable() bool {
	return true
}

func (b *Build) Run(ctx *context.Context) (err error) {
	target := "apk"
	switch b.typ {
	case constant.TypeAndroid:
		target = "apk"
		b.flutter.Remove("Gradle锁文件", filepath.Join(os.Getenv("GRADLE_USER_HOME"), "caches/journal-1"))
	case constant.TypeWeb:
		target = "web"
	}

	arguments := args.New().Build().Subcommand("build", target)
	err = b.flutter.Exec(ctx, arguments.Build())

	return
}
