package step

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/dronestock/flutter/internal/internal"
	"github.com/magiconair/properties"
)

type Boost struct {
	source string
	typ    internal.Type
}

func NewBoost(source string, typ internal.Type) *Boost {
	return &Boost{
		source: source,
		typ:    typ,
	}
}

func (b *Boost) Runnable() bool {
	return true
}

func (b *Boost) Run(ctx context.Context) (err error) {
	switch b.typ {
	case internal.TypeAndroid:
		err = b.android(ctx)
	}

	return
}

func (b *Boost) android(ctx context.Context) (err error) {
	gradleConfigPath := filepath.Join(b.source, "android", "gradle", "wrapper", "gradle-wrapper.properties")
	if config, lfe := properties.LoadFile(gradleConfigPath, properties.UTF8); nil != lfe {
		err = lfe
	} else if lpe := b.linkPlatforms(ctx); nil != lpe {
		err = lpe
	} else {
		err = b.boostGradle(config, gradleConfigPath)
	}

	return
}

func (b *Boost) linkPlatforms(_ context.Context) (err error) {
	link := filepath.Join(os.Getenv(internal.EnvAndroidHome), internal.DirPlatforms)
	modules := filepath.Join(os.Getenv(internal.EnvFlutterCache), internal.DirPlatforms)
	if _, se := os.Stat(modules); nil != se && os.IsNotExist(se) {
		err = os.MkdirAll(modules, os.ModePerm)
	}
	if _, se := os.Lstat(link); nil != se && os.IsNotExist(se) && nil == err {
		err = os.Symlink(modules, link)
	}

	return
}

func (b *Boost) boostGradle(prop *properties.Properties, path string) (err error) {
	if version, ve := b.gradleVersion(prop.MustGetString(internal.KeyDistributionUrl)); nil != ve {
		err = ve
	} else if file, oe := os.OpenFile(path, os.O_WRONLY, os.ModePerm); nil != oe {
		err = oe
	} else {
		url := fmt.Sprintf("https://mirrors.cloud.tencent.com/gradle/gradle-%s.zip", version)
		_, _ = prop.MustSet(internal.KeyDistributionUrl, url)
		_, err = prop.Write(file, properties.UTF8)
	}

	return
}

func (b *Boost) gradleVersion(url string) (version string, err error) {
	re := regexp.MustCompile(".*gradle-(.+)\\.zip")
	match := re.FindStringSubmatch(url)
	if 1 < len(match) {
		version = match[1]
	}

	return
}
