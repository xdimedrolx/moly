package app

import (
	"context"
	"github.com/xdimedrolx/moly/pkg/platform/correlation"
	"go.opencensus.io/trace"
)

// ContextExtractor extracts fields from a context.
func ContextExtractor(ctx context.Context) map[string]interface{} {
	fields := make(map[string]interface{})

	if correlationID, ok := correlation.FromContext(ctx); ok {
		fields["correlation_id"] = correlationID
	}

	if span := trace.FromContext(ctx); span != nil {
		spanCtx := span.SpanContext()

		fields["trace_id"] = spanCtx.TraceID.String()
		fields["span_id"] = spanCtx.SpanID.String()
	}

	return fields
}
