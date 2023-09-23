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
// Service exported by server reflection

// Warning: this entire file is deprecated. Use this instead:
// https://github.com/grpc/grpc-proto/blob/master/grpc/reflection/v1/reflection.proto

// Code generated by protoc-gen-defaults. DO NOT EDIT.

package reflectionv1alpha

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	_ *timestamppb.Timestamp
	_ *durationpb.Duration
	_ *wrapperspb.BoolValue
)

func (x *ServerReflectionRequest) Default() {
}

func (x *ExtensionRequest) Default() {
}

func (x *ServerReflectionResponse) Default() {
}

func (x *FileDescriptorResponse) Default() {
}

func (x *ExtensionNumberResponse) Default() {
}

func (x *ListServiceResponse) Default() {
}

func (x *ServiceResponse) Default() {
}

func (x *ErrorResponse) Default() {
}
