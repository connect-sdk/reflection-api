// Copyright 2016 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Service exported by server reflection.  A more complete description of how
// server reflection works can be found at
// https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
//
// The canonical version of this proto can be found at
// https://github.com/grpc/grpc-proto/blob/master/grpc/reflection/v1/reflection.proto

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: grpc/reflection/v1/reflection.proto

package reflectionv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/connect-sdk/reflection-api/proto/grpc/reflection/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// ServerReflectionName is the fully-qualified name of the ServerReflection service.
	ServerReflectionName = "grpc.reflection.v1.ServerReflection"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServerReflectionServerReflectionInfoProcedure is the fully-qualified name of the
	// ServerReflection's ServerReflectionInfo RPC.
	ServerReflectionServerReflectionInfoProcedure = "/grpc.reflection.v1.ServerReflection/ServerReflectionInfo"
)

// ServerReflectionClient is a client for the grpc.reflection.v1.ServerReflection service.
type ServerReflectionClient interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(context.Context) *connect.BidiStreamForClient[v1.ServerReflectionRequest, v1.ServerReflectionResponse]
}

// NewServerReflectionClient constructs a client for the grpc.reflection.v1.ServerReflection
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServerReflectionClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServerReflectionClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serverReflectionClient{
		serverReflectionInfo: connect.NewClient[v1.ServerReflectionRequest, v1.ServerReflectionResponse](
			httpClient,
			baseURL+ServerReflectionServerReflectionInfoProcedure,
			opts...,
		),
	}
}

// serverReflectionClient implements ServerReflectionClient.
type serverReflectionClient struct {
	serverReflectionInfo *connect.Client[v1.ServerReflectionRequest, v1.ServerReflectionResponse]
}

// ServerReflectionInfo calls grpc.reflection.v1.ServerReflection.ServerReflectionInfo.
func (c *serverReflectionClient) ServerReflectionInfo(ctx context.Context) *connect.BidiStreamForClient[v1.ServerReflectionRequest, v1.ServerReflectionResponse] {
	return c.serverReflectionInfo.CallBidiStream(ctx)
}

// ServerReflectionHandler is an implementation of the grpc.reflection.v1.ServerReflection service.
type ServerReflectionHandler interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(context.Context, *connect.BidiStream[v1.ServerReflectionRequest, v1.ServerReflectionResponse]) error
}

// NewServerReflectionHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServerReflectionHandler(svc ServerReflectionHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serverReflectionServerReflectionInfoHandler := connect.NewBidiStreamHandler(
		ServerReflectionServerReflectionInfoProcedure,
		svc.ServerReflectionInfo,
		opts...,
	)
	return "/grpc.reflection.v1.ServerReflection/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServerReflectionServerReflectionInfoProcedure:
			serverReflectionServerReflectionInfoHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServerReflectionHandler returns CodeUnimplemented from all methods.
type UnimplementedServerReflectionHandler struct{}

func (UnimplementedServerReflectionHandler) ServerReflectionInfo(context.Context, *connect.BidiStream[v1.ServerReflectionRequest, v1.ServerReflectionResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("grpc.reflection.v1.ServerReflection.ServerReflectionInfo is not implemented"))
}
