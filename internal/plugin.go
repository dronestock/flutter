package internal

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/flutter/internal/internal"
	"github.com/dronestock/flutter/internal/step"
)

type Plugin struct {
	drone.Base

	// 源代码目录
	Source string `default:"${SOURCE=.}" validate:"dirpath"`
	// 类型
	Type internal.Type `default:"${TYPE=android}" validate:"required,oneof=android"`
}

func NewPlugin() drone.Plugin {
	return new(Plugin)
}

func (p *Plugin) Config() drone.Config {
	return p
}

func (p *Plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewBoost(p.Source, p.Type)).Name("加速").Interrupt().Build(),
		drone.NewStep(step.NewClean(p.Base, p.Source)).Name("清理").Interrupt().Build(),
		drone.NewStep(step.NewGet(p.Base, p.Source)).Name("依赖").Interrupt().Build(),
		drone.NewStep(step.NewBuild(p.Base, p.Source, p.Type)).Name("打包").Interrupt().Build(),
	}
}
