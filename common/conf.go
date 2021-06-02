package common

import "time"

type Config struct {
	HttpConfig *HttpConfig      `mapstructure:"http"`
	KafkaConf  *KafkaConfig     `mapstructure:"kafka"`
	AsConf     *AerospikeConfig `mapstructure:"aerospike"`
}

type AerospikeConfig struct {
	Host                  string        `mapstructure:"host"`
	Port                  int           `mapstructure:"port"`
	ConnectTimeout        time.Duration `mapstructure:"connectTimeout"`
	IdleTimeout           time.Duration `mapstructure:"idleTimeout"`
	MinConnectionsPerNode int           `mapstructure:"minConnectionsPerNode"`
	ConnectionQueueSize   int           `mapstructure:"connectionQueueSize"`
	UdfPath               string        `mapstructure:"udfPath"`
	Namespace             string        `mapstructure:"namespace"`
}

type HttpConfig struct {
	Port int `mapstructure:"port"`
}

type KafkaConfig struct {
	Addr   string `mapstructure:"addr"`
	Topics Topics `mapstructure:"topics"`
}

type Topics struct {
	TestTopic string `mapstructure:"testTopic"`
}
