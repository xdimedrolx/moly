package service

import (
	"github.com/xdimedrolx/moly/pkg/log"
)

type Services interface {
	Logger() log.Logger
	ErrorHandler() log.ErrorHandler
}

type ServicesImpl struct {
	logger log.Logger
	errorHandler log.ErrorHandler
}

func NewContainer(
	logger log.Logger,
	errorHandler log.ErrorHandler,
) *ServicesImpl {
	return &ServicesImpl{
		logger,
		errorHandler,
	}
}

func (s *ServicesImpl) Logger() log.Logger {
	return s.logger
}

func (s *ServicesImpl) ErrorHandler() log.ErrorHandler {
	return s.errorHandler
}

var _ Services = (*ServicesImpl)(nil)
