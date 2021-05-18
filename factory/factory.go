package factory

import "github.com/nghiatrandev/cob-as-prt/handler"

type Factory interface {
	BuildHandler() handler.Handler
}
