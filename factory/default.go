package factory

import (
	. "github.com/nghiatrandev/cob-as-prt/common"
	"github.com/nghiatrandev/cob-as-prt/handler"
)

var HandlerKey = "Handler"

type defaultFactory struct {
	conf      *Config `mapstructure:"config"`
	container map[string]interface{}
}

func NewDefaultFactory(conf *Config) *defaultFactory {
	return &defaultFactory{
		conf:      conf,
		container: map[string]interface{}{},
	}
}

func (f *defaultFactory) BuildHandler() handler.Handler {
	if _, ok := f.container[HandlerKey]; !ok {
		return handler.NewHandler()
	}
	return f.container[HandlerKey].(handler.Handler)
}
