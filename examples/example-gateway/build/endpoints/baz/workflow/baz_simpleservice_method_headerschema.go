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
	"net/textproto"

	"github.com/uber/zanzibar/config"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsIDlClientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/baz/baz"
	endpointsIDlEndpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints-idl/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceHeaderSchemaWorkflow defines the interface for SimpleServiceHeaderSchema workflow
type SimpleServiceHeaderSchemaWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsIDlEndpointsBazBaz.SimpleService_HeaderSchema_Args,
	) (context.Context, *endpointsIDlEndpointsBazBaz.HeaderSchema, zanzibar.Header, error)
}

// NewSimpleServiceHeaderSchemaWorkflow creates a workflow
func NewSimpleServiceHeaderSchemaWorkflow(deps *module.Dependencies) SimpleServiceHeaderSchemaWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.baz.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.baz.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName))
		}
	}

	return &simpleServiceHeaderSchemaWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// simpleServiceHeaderSchemaWorkflow calls thrift client Baz.HeaderSchema
type simpleServiceHeaderSchemaWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// Handle calls thrift client.
func (w simpleServiceHeaderSchemaWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsIDlEndpointsBazBaz.SimpleService_HeaderSchema_Args,
) (context.Context, *endpointsIDlEndpointsBazBaz.HeaderSchema, zanzibar.Header, error) {
	clientRequest := convertToHeaderSchemaClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	var k string

	k = textproto.CanonicalMIMEHeaderKey("x-uber-foo")
	h, ok = reqHeaders.Get(k)
	if ok {
		clientHeaders[k] = h
	}
	k = textproto.CanonicalMIMEHeaderKey("x-uber-bar")
	h, ok = reqHeaders.Get(k)
	if ok {
		clientHeaders[k] = h
	}

	h, ok = reqHeaders.Get("Auth")
	if ok {
		clientHeaders["Auth"] = h
	}
	h, ok = reqHeaders.Get("Content-Type")
	if ok {
		clientHeaders["Content-Type"] = h
	}
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	h, ok = reqHeaders.Get("X-Token")
	if ok {
		clientHeaders["X-Token"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}
	for _, whitelistedHeader := range w.whitelistedDynamicHeaders {
		headerVal, ok := reqHeaders.Get(whitelistedHeader)
		if ok {
			clientHeaders[whitelistedHeader] = headerVal
		}
	}

	ctx, clientRespBody, _, err := w.Clients.Baz.HeaderSchema(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsIDlClientsBazBaz.AuthErr:
			serverErr := convertHeaderSchemaAuthErr(
				errValue,
			)

			return ctx, nil, nil, serverErr

		case *clientsIDlClientsBazBaz.OtherAuthErr:
			serverErr := convertHeaderSchemaOtherAuthErr(
				errValue,
			)

			return ctx, nil, nil, serverErr

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			return ctx, nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceHeaderSchemaClientResponse(clientRespBody)
	return ctx, response, resHeaders, nil
}

func convertToHeaderSchemaClientRequest(in *endpointsIDlEndpointsBazBaz.SimpleService_HeaderSchema_Args) *clientsIDlClientsBazBaz.SimpleService_HeaderSchema_Args {
	out := &clientsIDlClientsBazBaz.SimpleService_HeaderSchema_Args{}

	if in.Req != nil {
		out.Req = &clientsIDlClientsBazBaz.HeaderSchema{}
	} else {
		out.Req = nil
	}

	return out
}

func convertHeaderSchemaAuthErr(
	clientError *clientsIDlClientsBazBaz.AuthErr,
) *endpointsIDlEndpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsIDlEndpointsBazBaz.AuthErr{}
	return serverError
}
func convertHeaderSchemaOtherAuthErr(
	clientError *clientsIDlClientsBazBaz.OtherAuthErr,
) *endpointsIDlEndpointsBazBaz.OtherAuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsIDlEndpointsBazBaz.OtherAuthErr{}
	return serverError
}

func convertSimpleServiceHeaderSchemaClientResponse(in *clientsIDlClientsBazBaz.HeaderSchema) *endpointsIDlEndpointsBazBaz.HeaderSchema {
	out := &endpointsIDlEndpointsBazBaz.HeaderSchema{}

	return out
}
