package reflectionv1alphasdk

import (
	grpcreflect "connectrpc.com/grpcreflect"
)

// Reflector implements the underlying logic for gRPC's protobuf server reflection.
// They're configurable, so they can support straightforward process-local reflection or more complex proxying.
type Reflector = grpcreflect.Reflector

// A Reflector Namer lists the fully-qualified Protobuf service names available for reflection
// (for example, "acme.user.v1.UserService"). Namers must be safe to call concurrently.
type ReflectorNamer = grpcreflect.Namer

// NewReflector constructs a highly configurable Reflector:
// it can serve a dynamic list of services, proxy reflection requests to other backends, or use an alternate source of extension information.
var NewReflector = grpcreflect.NewReflector

var _ ReflectorNamer = NameCollection{}

// NameCollection is a list of names
type NameCollection []string

// WithName returns a NameCollection
func WithName(name ...string) NameCollection {
	return NameCollection(name)
}

// Names returns the names
func (c NameCollection) Names() []string {
	return c
}
