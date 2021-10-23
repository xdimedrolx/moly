package app

import "github.com/xdimedrolx/moly/pkg/log"

// These interfaces are aliased so that the module code is separated from the rest of the application.
// If the module is moved out of the app, copy the aliased interfaces here.

// Logger is the fundamental interface for all log operations.
type Logger = log.Logger

// ErrorHandler handles an error.
type ErrorHandler = log.ErrorHandler
