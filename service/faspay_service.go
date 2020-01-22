package service

import (
	vmFaspay "faspay/faspay_services/view_model"
)

type FaspayService interface {
	//Token
	//TokenHandler(req vmFaspay.EmptyVM) (*vmFaspay.FaspayResponse, error)
	//Register
	RegisterHandler(req vmFaspay.FaspayRegisterRequest) (*vmFaspay.FaspayResponse, error)
	//Confirm Register
	ConfirmRegisterHandler(req vmFaspay.FaspayConfirmRegisterRequest) (*vmFaspay.FaspayResponse, error)
	//Transfer
	TransferHandler(req vmFaspay.FaspayTransferRequest) (*vmFaspay.FaspayResponse, error)
}
