package step

import (
	"context"

	"github.com/dronestock/flutter/internal/internal/command"
	"github.com/goexl/args"
)

type Get struct {
	flutter *command.Flutter
}

func NewGet(flutter *command.Flutter) *Get {
	return &Get{
		flutter: flutter,
	}
}

func (g *Get) Runnable() bool {
	return true
}

func (g *Get) Run(ctx *context.Context) (err error) {
	return g.flutter.Exec(ctx, args.New().Build().Subcommand("pub", "get").Build())
}
