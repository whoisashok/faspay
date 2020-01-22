package usecase

import (
	vmFaspay "faspay/faspay_services/view_model"
)

type FaspayUsecase interface {
	//Token
	TokenHandler() (*vmFaspay.FaspayTokenResponse, error)
	//Register
	RegisterHandler(vmFaspay.FaspayRegisterRequest) (*vmFaspay.FaspayResponse, error)
	//Confirm Register
	ConfirmRegisterHandler(req vmFaspay.FaspayConfirmRegisterRequest) (*vmFaspay.FaspayResponse, error)
	//Transfer
	TransferHandler(req vmFaspay.FaspayTransferRequest) (*vmFaspay.FaspayResponse, error)
}
