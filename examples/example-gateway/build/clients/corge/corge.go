// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
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

package corgeClient

import (
	"context"
	"errors"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/uber/tchannel-go"
	"github.com/uber/zanzibar/runtime"

	clientsCorgeCorge "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/corge/corge"
)

// Client defines corge client interface.
type Client interface {
	EchoString(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsCorgeCorge.Corge_EchoString_Args,
	) (string, map[string]string, error)
}

// NewClient returns a new TChannel client for service corge.
func NewClient(gateway *zanzibar.Gateway) Client {
	serviceName := gateway.Config.MustGetString("clients.corge.serviceName")
	sc := gateway.Channel.GetSubChannel(serviceName, tchannel.Isolated)

	ip := gateway.Config.MustGetString("sidecarRouter.default.tchannel.ip")
	port := gateway.Config.MustGetInt("sidecarRouter.default.tchannel.port")
	sc.Peers().Add(ip + ":" + strconv.Itoa(int(port)))

	timeout := time.Millisecond * time.Duration(
		gateway.Config.MustGetInt("clients.corge.timeout"),
	)
	timeoutPerAttempt := time.Millisecond * time.Duration(
		gateway.Config.MustGetInt("clients.corge.timeoutPerAttempt"),
	)

	client := zanzibar.NewTChannelClient(gateway.Channel,
		&zanzibar.TChannelClientOption{
			ServiceName:       serviceName,
			Timeout:           timeout,
			TimeoutPerAttempt: timeoutPerAttempt,
		},
	)

	return &corgeClient{
		client: client,
		logger: gateway.Logger,
	}
}

// corgeClient is the TChannel client for downstream service.
type corgeClient struct {
	client zanzibar.TChannelClient
	logger *zap.Logger
}

// EchoString is a client RPC call for method "Corge::echoString"
func (c *corgeClient) EchoString(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsCorgeCorge.Corge_EchoString_Args,
) (string, map[string]string, error) {
	var result clientsCorgeCorge.Corge_EchoString_Result
	var resp string

	var fields []zapcore.Field
	fields = append(fields, zap.String("Downstream-Client", "corgeClient"))
	fields = append(fields, zap.String("Downstream-Method", "EchoString"))
	fields = append(fields, zap.Time("timestamp", time.Now().UTC()))
	for k, v := range reqHeaders {
		fields = append(fields, zap.String("Downstream-Request-Header-"+k, v))
	}
	fields = append(fields, zap.Any("Downstream-Request-Body", args))
	success, respHeaders, err := c.client.Call(
		ctx, "Corge", "echoString", reqHeaders, args, &result,
	)

	for k, v := range respHeaders {
		fields = append(fields, zap.String("Downstream-Response-Header-"+k, v))
	}
	fields = append(fields, zap.Any("Downstream-Response-Body", result))
	fields = append(fields, zap.Time("timestamp-finished", time.Now().UTC()))
	c.logger.Info(
		"Finished a downstream TChannel request",
		fields...)
	if err == nil && !success {
		switch {
		default:
			err = errors.New("corgeClient received no result or unknown exception for EchoString")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsCorgeCorge.Corge_EchoString_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}
