package endpoint

import (
	"context"

	endpointGRPC "github.com/go-kit/kit/endpoint"

	vmFaspay "faspay/faspay_services/view_model"
	serviceFaspay "faspay/service"
)

/*
// Tokens
func MakeFaspayTokenEndpoint(s serviceFaspay.FaspayService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(vmFaspay.EmptyVM)
		r, err := s.TokenHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
*/

// Register
func MakeFaspayRegisterEndpoint(s serviceFaspay.FaspayService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(vmFaspay.FaspayRegisterRequest)
		r, err := s.RegisterHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

// Confirm Register
func MakeFaspayConfirmRegisterEndpoint(s serviceFaspay.FaspayService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(vmFaspay.FaspayConfirmRegisterRequest)
		r, err := s.ConfirmRegisterHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

// Transfer
func MakeFaspayTransferEndpoint(s serviceFaspay.FaspayService) endpointGRPC.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(vmFaspay.FaspayTransferRequest)
		r, err := s.TransferHandler(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
