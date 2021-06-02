package controller

import (
	"fmt"
	"github.com/nghiatrandev/sample_project/dataservice"
)

type controller struct {
	kafka *dataservice.Kafka
}

func NewController(kafka *dataservice.Kafka) Controller {
	return &controller{
		kafka: kafka,
	}
}

func (c *controller) Test() {
	fmt.Println(1)
}
