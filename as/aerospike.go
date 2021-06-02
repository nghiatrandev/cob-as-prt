package as

import (
	as "github.com/aerospike/aerospike-client-go"
	. "github.com/nghiatrandev/sample_project/common"
)

// NewClient generates a new Client instance.
func NewClient(cfg *AerospikeConfig) (*as.Client, error) {
	//init policy
	p := as.NewClientPolicy()
	p.Timeout = cfg.ConnectTimeout
	p.IdleTimeout = cfg.IdleTimeout
	p.MinConnectionsPerNode = cfg.MinConnectionsPerNode
	p.ConnectionQueueSize = cfg.ConnectionQueueSize

	h := as.NewHost(cfg.Host, cfg.Port)

	return as.NewClientWithPolicyAndHost(p, h)
}
