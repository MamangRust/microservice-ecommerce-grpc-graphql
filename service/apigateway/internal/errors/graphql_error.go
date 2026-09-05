package graphqlerror

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/response"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	otelCodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ResolverFunction[T any] func(ctx context.Context) (T, error)

type resolverHandler struct {
	observability observability.TraceLoggerObservability
	logger        logger.LoggerInterface
}

func NewResolverHandler(obs observability.TraceLoggerObservability, logger logger.LoggerInterface) *resolverHandler {
	return &resolverHandler{
		observability: obs,
		logger:        logger,
	}
}

func ResolverHandle[T any](h *resolverHandler, method string, ctx context.Context, fn ResolverFunction[T]) (T, error) {
	var zero T

	ctx, span, end, status, logSuccess := h.observability.StartTracingAndLogging(
		ctx,
		method,
	)
	defer func() {
		end(status)
	}()

	result, err := fn(ctx)
	if err != nil {
		status = "error"
		h.handleResolverError(err, span, method)
		return zero, err
	}

	logSuccess("Resolver completed successfully")
	return result, nil
}

func (h *resolverHandler) handleResolverError(err error, span trace.Span, method string) {
	traceID := span.SpanContext().TraceID().String()

	h.logger.Error(
		fmt.Sprintf("Resolver error in %s", method),
		zap.Error(err),
		zap.String("trace.id", traceID),
	)

	span.SetAttributes(
		attribute.String("trace.id", traceID),
		attribute.String("error", err.Error()),
	)
	span.RecordError(err)
	span.SetStatus(otelCodes.Error, err.Error())
}

func ToGraphqlErrorFromErrorResponse(err error) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return fmt.Errorf("graphql error: %v", err)
	}

	httpCode := grpcToHttpCode(st.Code())
	statusText := http.StatusText(httpCode)
	message := st.Message()

	return NewGraphqlError(statusText, message, httpCode)
}

func NewGraphqlError(statusText string, message string, code int) error {
	errResp := &response.ErrorResponse{
		Status:  statusText,
		Message: message,
		Code:    code,
	}

	return fmt.Errorf("graphql error: [%d] %s - %s", errResp.Code, errResp.Status, errResp.Message)
}

func grpcToHttpCode(code codes.Code) int {
	switch code {
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	default:
		return http.StatusInternalServerError
	}
}
