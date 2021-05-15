package dataservice

import (
	"fmt"
	aerospike "github.com/aerospike/aerospike-client-go"
)

const (
	SetUserInsight = "user_insight"
)

type dataServiceAerospike struct {
	client    *aerospike.Client
	namespace string
}

func NewDataServiceAerospike(client *aerospike.Client, namespace string) (DataServiceAerospike, error) {
	// ensure indexes
	if err := EnsureIndexes(client, []IndexDescription{
		// ensure index for classification
		{namespace, "student", "idx_rule_id", "rule_id", aerospike.STRING},
		{namespace, "teacher", "idx_user_id", "userId", aerospike.STRING},
	}); err != nil {
		return nil, fmt.Errorf("Error ensuring indexes: %+v", err)
	}

	// build ds
	ds := &dataServiceAerospike{
		client:    client,
		namespace: namespace,
	}

	return ds, nil
}

func (ds *dataServiceAerospike) IsConnected() bool {
	return true
}
