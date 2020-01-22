package transporthttp

import (
	"net/http"

	"context"

	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	endpointFaspay "faspay/endpoint"
	modelHttpFaspay "faspay/model_http"
)

// MakeExtMemberHttpHandler - make http handler
func MakeFaspayHttpHandler(ctx context.Context, endpointFaspay endpointFaspay.FaspayEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(modelHttpFaspay.EncodeError),
	}

	apiV1 := r.PathPrefix("/account/api").Subrouter()

	//1. GET /tokens
	apiV1.Methods("GET").Path("/tokens").Handler(httpTransport.NewServer(
		endpointFaspay.FaspayTokenEndpoint,
		modelHttpFaspay.DecodeFaspayTokenRequest,
		modelHttpFaspay.EncodeResponse,
		options...,
	))

	//2. POST /register
	apiV1.Methods("POST").Path("/register").Handler(httpTransport.NewServer(
		endpointFaspay.FaspayRegisterEndpoint,
		modelHttpFaspay.DecodeFaspayRegisterRequest,
		modelHttpFaspay.EncodeResponse,
		options...,
	))

	//3. POST /Confirm Register
	apiV1.Methods("POST").Path("/register/confirm").Handler(httpTransport.NewServer(
		endpointFaspay.FaspayConfirmRegisterEndpoint,
		modelHttpFaspay.DecodeFaspayConfirmRegisterRequest,
		modelHttpFaspay.EncodeResponse,
		options...,
	))

	//4. POST /Transfer
	apiV1.Methods("POST").Path("/transfer").Handler(httpTransport.NewServer(
		endpointFaspay.FaspayTransferEndpoint,
		modelHttpFaspay.DecodeFaspayTransferRequest,
		modelHttpFaspay.EncodeResponse,
		options...,
	))

	return apiV1
}
