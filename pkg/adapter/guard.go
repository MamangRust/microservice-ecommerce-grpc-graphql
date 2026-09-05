package adapter

import (
	"github.com/MamangRust/microservice-ecommerce-pkg/resilience"
)

// guardSetter is implemented by every gRPC adapter so WithDependencyGuard can
// attach a guard without changing each constructor's signature for existing
// call sites.
type guardSetter interface {
	setGuard(*resilience.DependencyGuard)
}

// WithDependencyGuard attaches a dependency guard (per-call timeout + circuit
// breaker + bulkhead) to a gRPC adapter. Passing nil disables guarding.
//
// Usage:
//
//	guard := resilience.NewDependencyGuard("product", 5, 30, 100, 3*time.Second, srv.Logger)
//	productAdapter := adapter.NewProductAdapter(q, c, adapter.WithDependencyGuard(guard))
func WithDependencyGuard(guard *resilience.DependencyGuard) func(guardSetter) {
	return func(s guardSetter) {
		s.setGuard(guard)
	}
}
