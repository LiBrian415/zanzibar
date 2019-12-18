// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package workflow

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"

	endpointsServerlessServerless "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/serverless/serverless"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/serverless/module"
	"go.uber.org/zap"
)

// ServerlessServerlessArgWithHeadersDummyWorkflow defines the interface for ServerlessServerlessArgWithHeaders workflow
type ServerlessServerlessArgWithHeadersDummyWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsServerlessServerless.Serverless_ServerlessArgWithHeaders_Args,
	) (*endpointsServerlessServerless.Response, zanzibar.Header, error)
}

// NewServerlessServerlessArgWithHeadersDummyWorkflow creates a workflow
func NewServerlessServerlessArgWithHeadersDummyWorkflow(deps *module.Dependencies) ServerlessServerlessArgWithHeadersDummyWorkflow {

	return &serverlessServerlessArgWithHeadersDummyWorkflow{
		Logger: deps.Default.Logger,
	}
}

// serverlessServerlessArgWithHeadersDummyWorkflow calls thrift client .
type serverlessServerlessArgWithHeadersDummyWorkflow struct {
	Logger *zap.Logger
}

// Handle processes the request without a downstream
func (w serverlessServerlessArgWithHeadersDummyWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsServerlessServerless.Serverless_ServerlessArgWithHeaders_Args,
) (*endpointsServerlessServerless.Response, zanzibar.Header, error) {
	response := convertServerlessArgWithHeadersDummyResponse(r)

	serverlessHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		serverlessHeaders["X-Deputy-Forwarded"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		serverlessHeaders["X-Uuid"] = h
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}
	h, ok = serverlessHeaders["X-Token"]
	if ok {
		resHeaders.Set("X-Token", h)
	}
	h, ok = serverlessHeaders["X-Uuid"]
	if ok {
		resHeaders.Set("X-Uuid", h)
	}

	return response, resHeaders, nil
}

func convertServerlessArgWithHeadersDummyResponse(in *endpointsServerlessServerless.Serverless_ServerlessArgWithHeaders_Args) *endpointsServerlessServerless.Response {
	out := &endpointsServerlessServerless.Response{}

	return out
}
