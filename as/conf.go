package as

import (
	"time"
)

type Config struct {
	AsConf *AerospikeConfig `mapstructure:"aerospike"`
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
