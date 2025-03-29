package middleware

import "context"

type OpenMiddleware interface {
	PrepareService(ctx context.Context, middlewareName, middlewareNamespace string) error
	PrepareHeadlessService(ctx context.Context, middlewareName, middlewareNamespace string) error
	PreparePodService(ctx context.Context, middlewareName, middlewareNamespace string) error
	PrepareInitConfigMap(ctx context.Context, middlewareName, middlewareNamespace string) error
	PrepareMiddlewareConfigMap(ctx context.Context, middlewareName, middlewareNamespace string) error
	UpdateMiddlewareStatus(ctx context.Context, middlewareName, middlewareNamespace string) error
}
