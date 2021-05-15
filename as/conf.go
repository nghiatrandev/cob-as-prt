package as

import "time"

type AerospikeConfig struct {
	Host                  string
	Port                  int
	ConnectTimeout        time.Duration
	IdleTimeout           time.Duration
	MinConnectionsPerNode int
	ConnectionQueueSize   int
	UdfPath               string
	Namespace             string
}
