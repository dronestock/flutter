package internal

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/flutter/internal/internal/constant"
	"github.com/dronestock/flutter/internal/step"
)

type plugin struct {
	drone.Base

	// 执行程序
	Binary string `default:"${BINARY=flutter}"`
	// 源代码目录
	Source string `default:"${SOURCE=.}" validate:"dirpath"`
	// 类型
	Type constant.Type `default:"${TYPE=android}" validate:"required,oneof=android web"`
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewBoost(p.Source, p.Type)).Name("加速").Interrupt().Build(),
		drone.NewStep(step.NewClean(p.Base, p.Binary, p.Source)).Name("清理").Interrupt().Build(),
		drone.NewStep(step.NewGet(p.Base, p.Binary, p.Source)).Name("依赖").Interrupt().Build(),
		drone.NewStep(step.NewBuild(p.Base, p.Binary, p.Source, p.Type)).Name("打包").Interrupt().Build(),
	}
}
