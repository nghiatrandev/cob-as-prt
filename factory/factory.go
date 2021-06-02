package factory

import "github.com/nghiatrandev/sample_project/handler"

type Factory interface {
	BuildHandler() handler.Handler
}
