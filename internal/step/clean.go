package step

import (
	"context"

	"github.com/dronestock/flutter/internal/internal/command"
	"github.com/goexl/args"
)

type Clean struct {
	flutter *command.Flutter
}

func NewClean(flutter *command.Flutter) *Clean {
	return &Clean{
		flutter: flutter,
	}
}

func (c *Clean) Runnable() bool {
	return true
}

func (c *Clean) Run(ctx *context.Context) error {
	return c.flutter.Exec(ctx, args.New().Build().Subcommand("clean").Build())
}
