package reflectionv1sdk

import (
	connect "connectrpc.com/connect"
	grpcreflect "connectrpc.com/grpcreflect"
	interceptor "github.com/connect-sdk/interceptor"
	middleware "github.com/connect-sdk/middleware"
	chi "github.com/go-chi/chi/v5"
)

// ReflectionServiceController represents an instance for grpc.reflection.v1.ReflectionServiceHandler handler.
type ReflectionServiceHandler struct {
	// Reflector implements the underlying logic for gRPC's protobuf server reflection.
	Reflector *Reflector
}

// Mount mounts the controller to a given router.
func (x *ReflectionServiceHandler) Mount(r chi.Router) {
	var options []connect.HandlerOption
	// prepare the options
	options = append(options, interceptor.WithTracer())
	options = append(options, interceptor.WithLogger())
	options = append(options, interceptor.WithRecovery())
	options = append(options, interceptor.WithValidator())
	options = append(options, interceptor.WithContext())

	r.Group(func(r chi.Router) {
		// mount the middleware
		r.Use(middleware.WithLogger())
		// create the handler
		path, handler := grpcreflect.NewHandlerV1(x.Reflector, options...)
		// mount the handler
		r.Mount(path, handler)
	})
}
