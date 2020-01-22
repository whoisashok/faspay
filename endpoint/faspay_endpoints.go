package endpoint

import (
	endpointHttp "github.com/go-kit/kit/endpoint"
)

type FaspayEndpoints struct {
	FaspayTokenEndpoint           endpointHttp.Endpoint
	FaspayRegisterEndpoint        endpointHttp.Endpoint
	FaspayConfirmRegisterEndpoint endpointHttp.Endpoint
	FaspayTransferEndpoint        endpointHttp.Endpoint
}
