package dataservice

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go"
	"github.com/aerospike/aerospike-client-go/types"
)

// This helper waits until the index creation task is completed and
// processes the returned error if there is one
func ProcessIndexTask(provider func() (*aerospike.IndexTask, error), errorIfAlreadyExists bool) error {
	indexTask, err := provider()
	if err != nil {
		if asErr, ok := err.(types.AerospikeError); ok && asErr.ResultCode() == types.INDEX_FOUND {
			if !errorIfAlreadyExists {
				// we good, no error
				return nil
			}
		}
		return fmt.Errorf("error creating index: %s", err)
	}
	return <-indexTask.OnComplete()
}

type IndexDescription struct {
	Namespace string
	Set       string
	Index     string
	Bin       string
	Type      aerospike.IndexType
}

// This helper launches indexes creation tasks and processes
// their results
func EnsureIndexes(client *aerospike.Client, descriptions []IndexDescription) error {
	for _, d := range descriptions {
		if err := ProcessIndexTask(func() (*aerospike.IndexTask, error) {
			return client.CreateIndex(nil, d.Namespace, d.Set, d.Index, d.Bin, d.Type)
		}, false); err != nil {
			return err
		}
	}
	return nil
}
