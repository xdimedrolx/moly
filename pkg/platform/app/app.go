package app

import (
	"github.com/xdimedrolx/moly/pkg/log"
)

type App interface {
	Dispose()
	Logger() log.Logger
}
