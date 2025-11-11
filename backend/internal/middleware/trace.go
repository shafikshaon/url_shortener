package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shafikshaon/url_shortener/internal/logger"
)

const (
	// TraceIDHeader is the HTTP header for trace ID
	TraceIDHeader = "X-Trace-ID"
	// ContextKey for storing context in gin.Context
	ContextKey = "ctx"
)

// TraceMiddleware adds a trace ID to each request
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if trace ID exists in header, otherwise create new one
		traceID := c.GetHeader(TraceIDHeader)

		var ctx context.Context
		if traceID != "" {
			ctx = logger.WithTraceIDValue(c.Request.Context(), traceID)
		} else {
			ctx = logger.WithTraceID(c.Request.Context())
			traceID = logger.GetTraceID(ctx)
		}

		// Store context in gin.Context for later use
		c.Set(ContextKey, ctx)

		// Set trace ID in response header
		c.Header(TraceIDHeader, traceID)

		// Update request context
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// GetContext retrieves the context from gin.Context
func GetContext(c *gin.Context) context.Context {
	if ctx, exists := c.Get(ContextKey); exists {
		if ctxValue, ok := ctx.(context.Context); ok {
			return ctxValue
		}
	}
	// Fallback to request context
	return c.Request.Context()
}
