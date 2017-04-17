// Code generated by zanzibar
// @generated

package endpoints

import (
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow"
	"github.com/uber/zanzibar/examples/example-gateway/middlewares/example"
	"github.com/uber/zanzibar/runtime/middlewares/logger"

	"github.com/uber/zanzibar/runtime"
)

// Endpoints is a struct that holds all the endpoints
type Endpoints struct {
	ArgNotStructHandler     *bar.ArgNotStructHandler
	MissingArgHandler       *bar.MissingArgHandler
	NoRequestHandler        *bar.NoRequestHandler
	NormalHandler           *bar.NormalHandler
	TooManyArgsHandler      *bar.TooManyArgsHandler
	CallHandler             *baz.CallHandler
	SimpleHandler           *baz.SimpleHandler
	SimpleFutureHandler     *baz.SimpleFutureHandler
	SaveContactsHandler     *contacts.SaveContactsHandler
	AddCredentialsHandler   *googlenow.AddCredentialsHandler
	CheckCredentialsHandler *googlenow.CheckCredentialsHandler
}

// CreateEndpoints bootstraps the endpoints.
func CreateEndpoints(
	gateway *zanzibar.Gateway,
) interface{} {
	return &Endpoints{
		ArgNotStructHandler:     bar.NewArgNotStructEndpoint(gateway),
		MissingArgHandler:       bar.NewMissingArgEndpoint(gateway),
		NoRequestHandler:        bar.NewNoRequestEndpoint(gateway),
		NormalHandler:           bar.NewNormalEndpoint(gateway),
		TooManyArgsHandler:      bar.NewTooManyArgsEndpoint(gateway),
		CallHandler:             baz.NewCallEndpoint(gateway),
		SimpleHandler:           baz.NewSimpleEndpoint(gateway),
		SimpleFutureHandler:     baz.NewSimpleFutureEndpoint(gateway),
		SaveContactsHandler:     contacts.NewSaveContactsEndpoint(gateway),
		AddCredentialsHandler:   googlenow.NewAddCredentialsEndpoint(gateway),
		CheckCredentialsHandler: googlenow.NewCheckCredentialsEndpoint(gateway),
	}
}

// Register will register all endpoints
func Register(g *zanzibar.Gateway, router *zanzibar.Router) {
	endpoints := CreateEndpoints(g).(*Endpoints)

	router.Register(
		"POST", "/bar/arg-not-struct-path",
		zanzibar.NewEndpoint(
			g,
			"bar",
			"argNotStruct",
			endpoints.ArgNotStructHandler.HandleRequest,
		),
	)
	router.Register(
		"GET", "/bar/missing-arg-path",
		zanzibar.NewEndpoint(
			g,
			"bar",
			"missingArg",
			endpoints.MissingArgHandler.HandleRequest,
		),
	)
	router.Register(
		"GET", "/bar/no-request-path",
		zanzibar.NewEndpoint(
			g,
			"bar",
			"noRequest",
			endpoints.NoRequestHandler.HandleRequest,
		),
	)
	router.Register(
		"POST", "/bar/bar-path",
		zanzibar.NewEndpoint(
			g,
			"bar",
			"normal",
			zanzibar.NewStack([]zanzibar.MiddlewareHandle{
				example.NewMiddleWare(
					g,
					example.Options{
						Foo: "test",
					},
				),
				logger.NewMiddleWare(
					g,
					logger.Options{},
				),
			}, endpoints.NormalHandler.HandleRequest).Handle,
		),
	)
	router.Register(
		"POST", "/bar/too-many-args-path",
		zanzibar.NewEndpoint(
			g,
			"bar",
			"tooManyArgs",
			endpoints.TooManyArgsHandler.HandleRequest,
		),
	)
	router.Register(
		"POST", "/baz/call-path",
		zanzibar.NewEndpoint(
			g,
			"baz",
			"call",
			endpoints.CallHandler.HandleRequest,
		),
	)
	router.Register(
		"GET", "/baz/simple-path",
		zanzibar.NewEndpoint(
			g,
			"baz",
			"simple",
			endpoints.SimpleHandler.HandleRequest,
		),
	)
	router.Register(
		"GET", "/baz/simple-future-path",
		zanzibar.NewEndpoint(
			g,
			"baz",
			"simpleFuture",
			endpoints.SimpleFutureHandler.HandleRequest,
		),
	)
	router.Register(
		"POST", "/contacts/:userUUID/contacts",
		zanzibar.NewEndpoint(
			g,
			"contacts",
			"saveContacts",
			endpoints.SaveContactsHandler.HandleRequest,
		),
	)
	router.Register(
		"POST", "/googlenow/add-credentials",
		zanzibar.NewEndpoint(
			g,
			"googlenow",
			"addCredentials",
			endpoints.AddCredentialsHandler.HandleRequest,
		),
	)
	router.Register(
		"POST", "/googlenow/check-credentials",
		zanzibar.NewEndpoint(
			g,
			"googlenow",
			"checkCredentials",
			endpoints.CheckCredentialsHandler.HandleRequest,
		),
	)

}
