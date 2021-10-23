package app

import (
	"github.com/xdimedrolx/moly/pkg/log"
	"github.com/xdimedrolx/moly/pkg/platform/http"
)

type Kernel struct {
	httpServer *http.Server
	logger     log.Logger
}
