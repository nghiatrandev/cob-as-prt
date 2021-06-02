package factory

import (
	. "github.com/nghiatrandev/sample_project/common"
	"github.com/nghiatrandev/sample_project/controller"
	"github.com/nghiatrandev/sample_project/dataservice"
	"github.com/nghiatrandev/sample_project/handler"
)

const (
	KeyDataExtractor = "DataExtractor"
	KeyDataService   = "DataService"
	KeyKafkaService  = "KafkaService"
	KeyController    = "Controller"
	KeyHandler       = "Handler"
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
		f.container[HandlerKey] = handler.NewHandler()
	}
	return f.container[KeyHandler].(handler.Handler)
}

func (f *defaultFactory) BuildController() {
	if _, exist := f.container[KeyController]; !exist {
		f.container[KeyController] = controller.NewController(f.BuildKafkaService())
	}
}

func (f *defaultFactory) BuildKafkaService() *dataservice.Kafka {
	if _, exist := f.container[KeyKafkaService]; !exist {
		//build config
		conf := KafkaConfig{
			Addr:   f.conf.KafkaConf.Addr,
			Topics: f.conf.KafkaConf.Topics,
		}

		f.container[KeyKafkaService] = dataservice.NewKafka(conf)
	}

	return f.container[KeyKafkaService].(*dataservice.Kafka)
}
