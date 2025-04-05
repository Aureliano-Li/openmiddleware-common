package middleware

import "context"

type OpenMiddleware interface {
	PrepareInitConfigMap(ctx context.Context, middleware interface{}) error
	PrepareMiddlewareConfigMap(ctx context.Context, middleware interface{}) error
	PrepareStatefulSet(ctx context.Context, middleware interface{}) error
	CreateStatefulSet(ctx context.Context, middleware interface{}) error
	UpdateStatefulSet(ctx context.Context, middleware interface{}) error
	PrepareService(ctx context.Context, middleware interface{}) error
	PrepareHeadlessService(ctx context.Context, middleware interface{}) error
	PreparePodService(ctx context.Context, middleware interface{}) error
	UpdateMiddlewareStatus(ctx context.Context, middleware interface{}) error
}
