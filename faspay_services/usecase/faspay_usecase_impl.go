package usecase

import (
	"encoding/json"
	cfg "faspay/config"
	serviceFaspay "faspay/faspay_services"
	httpClient "faspay/faspay_services/http_client"
	vmFaspay "faspay/faspay_services/view_model"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	logging "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var config cfg.Config

type FaspayUsecaseImpl struct {
	//FaspayRegisterRepo repository.FaspayRepository
	logger logging.Logger
}

func NewFaspayUsecaseImpl(logger logging.Logger) FaspayUsecase {
	return &FaspayUsecaseImpl{logger}
}

func (u *FaspayUsecaseImpl) TokenHandler() (*vmFaspay.FaspayTokenResponse, error) {

	var time_now = time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println("time_now	=", time_now)

	// Authentication
	AuthorizationKey, err := u.AuthorizationKey()
	if err != nil {
		fmt.Println("AuthorizationKey err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl TokenHandler", "error", err)
		return nil, err
	}
	//fmt.Println("AuthorizationKey=", AuthorizationKey)

	TokenForSignature, err := u.TokenForSignature()
	if err != nil {
		fmt.Println("TokenForSignature err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl TokenHandler", "error", err)
		return nil, err
	}
	//fmt.Println("\n")
	//fmt.Println("TokenForSignature=", TokenForSignature)

	SignatureGetAccessToken, err := u.SignatureGetAccessToken(TokenForSignature, time_now)
	if err != nil {
		fmt.Println("GetAccessToken err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl TokenHandler", "error", err)
		return nil, err
	}
	//fmt.Println("\n")
	//fmt.Println("SignatureGetAccessToken=", SignatureGetAccessToken)

	TokenRequest := vmFaspay.TokenRequest{
		SignatureGetAccessToken: SignatureGetAccessToken,
		AuthorizationKey:        AuthorizationKey,
		TimeNow:                 time_now,
	}
	//fmt.Println("\n")
	//fmt.Println("TokenRequest=", TokenRequest)

	GetAccessToken, err := u.GetAccessToken(TokenRequest)
	if err != nil {
		fmt.Println("GetAccessToken err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl TokenHandler", "error", err)
		return nil, err
	}
	//fmt.Println("\n")
	//fmt.Println("GetAccessToken=", GetAccessToken)

	TokenForOtherService, err := u.TokenForOtherService(GetAccessToken)
	if err != nil {
		fmt.Println("ValidateToken err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl TokenHandler", "error", err)
		return nil, err
	}
	//fmt.Println("\n")
	//fmt.Println("TokenForOtherService=", TokenForOtherService)

	SignatureForOthers, err := u.SignatureRegister(TokenForOtherService, time_now)
	if err != nil {
		fmt.Println("ValidateToken err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl TokenHandler", "error", err)
		return nil, err
	}
	//fmt.Println("\n")
	//fmt.Println("SignatureForOthers=", SignatureForOthers)

	faspayTokenResponse := vmFaspay.FaspayTokenResponse{
		SignatureKey:     SignatureForOthers,
		AuthorizationKey: AuthorizationKey,
		TimeNow:          time_now,
	}

	/*	faspayResponse := vmFaspay.FaspayResponse{
			Status:   http.StatusText(http.StatusOK),
			Code:     http.StatusOK,
			CodeType: serviceFaspay.SuccessCodeType,
			Message:  serviceFaspay.MessageUpdatedSuccess,
			Data:     faspayTokenResponse,
		}
	*/
	return &faspayTokenResponse, nil
}

func (u *FaspayUsecaseImpl) RegisterHandler(req vmFaspay.FaspayRegisterRequest) (*vmFaspay.FaspayResponse, error) {

	TokenHandler, err := u.TokenHandler()

	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	req.SignatureKey = TokenHandler.SignatureKey
	req.AuthorizationKey = TokenHandler.AuthorizationKey
	req.TimeNow = TokenHandler.TimeNow

	resp, err := clientService.FaspayService.FaspayRegister(req)
	if err != nil {
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl RegisterHandler", "error", err)
		return nil, err
	}

	RegisterHandler, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("RegisterHandler	=", string(RegisterHandler))

	vmResHTTPJSON := string(RegisterHandler)
	faspayResponse := vmFaspay.FaspayResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceFaspay.SuccessCodeType,
		Message:  serviceFaspay.MessageRegisterSuccess,
		Data:     vmResHTTPJSON,
	}

	return &faspayResponse, nil
}

//Confirm Register
func (u *FaspayUsecaseImpl) ConfirmRegisterHandler(req vmFaspay.FaspayConfirmRegisterRequest) (*vmFaspay.FaspayResponse, error) {

	TokenHandler, err := u.TokenHandler()

	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	req.SignatureKey = TokenHandler.SignatureKey
	req.AuthorizationKey = TokenHandler.AuthorizationKey
	req.TimeNow = TokenHandler.TimeNow

	resp, err := clientService.FaspayService.FaspayConfirmRegister(req)
	if err != nil {
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl ConfirmRegisterHandler", "error", err)
		return nil, err
	}

	ConfirmRegisterHandler, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("ConfirmRegisterHandler=", string(ConfirmRegisterHandler))

	vmResHTTPJSON := string(ConfirmRegisterHandler)
	faspayResponse := vmFaspay.FaspayResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceFaspay.SuccessCodeType,
		Message:  serviceFaspay.MessageRegisterConfirmSuccess,
		Data:     vmResHTTPJSON,
	}
	return &faspayResponse, nil
}

//Transfer
func (u *FaspayUsecaseImpl) TransferHandler(req vmFaspay.FaspayTransferRequest) (*vmFaspay.FaspayResponse, error) {

	TokenHandler, err := u.TokenHandler()

	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	req.SignatureKey = TokenHandler.SignatureKey
	req.AuthorizationKey = TokenHandler.AuthorizationKey
	req.TimeNow = TokenHandler.TimeNow
	resp, err := clientService.FaspayService.FaspayTransfer(req)
	if err != nil {
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl TransferHandler", "error", err)
		return nil, err
	}

	FaspayTransfer, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("FaspayTransfer=", string(FaspayTransfer))

	vmResHTTPJSON := string(FaspayTransfer)
	faspayResponse := vmFaspay.FaspayResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceFaspay.SuccessCodeType,
		Message:  serviceFaspay.MessageTransferSuccess,
		Data:     vmResHTTPJSON,
	}

	return &faspayResponse, nil
}

// AuthorizationKey
func (u *FaspayUsecaseImpl) AuthorizationKey() (string, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	resp, err := clientService.FaspayService.AuthorizationKey()
	if err != nil {
		fmt.Println("AuthorizationKey err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl AuthorizationKey", "error", err)
		return "", err
	}
	AuthorizationKey, _ := ioutil.ReadAll(resp.Body)
	return string(AuthorizationKey), nil
}

// Token For signature get token
func (u *FaspayUsecaseImpl) TokenForSignature() (string, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	resp, err := clientService.FaspayService.TokenForSignature()
	if err != nil {
		fmt.Println("TokenForSignature err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl TokenForSignature", "error", err)
		return "", err
	}
	GetAccessToken, _ := ioutil.ReadAll(resp.Body)
	return string(GetAccessToken), nil
}

// Signature for get token
func (u *FaspayUsecaseImpl) SignatureGetAccessToken(SignatureForTokenRequest, time_now string) (string, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	resp, err := clientService.FaspayService.SignatureGetAccessToken(SignatureForTokenRequest, time_now)
	if err != nil {
		fmt.Println("AuthorizationKey err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl SignatureGetAccessToken", "error", err)
		return "", err
	}
	SignatureGetAccessToken, _ := ioutil.ReadAll(resp.Body)
	return string(SignatureGetAccessToken), nil
}

type jsonFormater struct {
	Access_token string `json:"access_token"`
	Message      string `json:"message"`
}

// Get Access Token
func (u *FaspayUsecaseImpl) GetAccessToken(req vmFaspay.TokenRequest) (string, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	resp, err := clientService.FaspayService.GetAccessToken(req)
	if err != nil {
		fmt.Println("resp err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl GetAccessToken", "error", err)
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("body	=", string(body))

	var data jsonFormater
	//var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	return data.Access_token, nil
}

// Token for other services
func (u *FaspayUsecaseImpl) TokenForOtherService(GetAccessToken string) (string, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	resp, err := clientService.FaspayService.TokenForOtherService(GetAccessToken)
	if err != nil {
		fmt.Println("AuthorizationKey err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl GetAccessToken", "error", err)
		return "", err
	}
	TokenForOtherService, _ := ioutil.ReadAll(resp.Body)
	return string(TokenForOtherService), nil
}

// Signature Register
func (u *FaspayUsecaseImpl) SignatureRegister(TokenForOtherService, time_now string) (string, error) {
	// HTTP Client Service to call other service
	clientService := httpClient.NewClient(nil)

	resp, err := clientService.FaspayService.SignatureRegister(TokenForOtherService, time_now)
	if err != nil {
		fmt.Println("AuthorizationKey err=", err)
		level.Error(u.logger).Log("function", "FaspayUsecaseImpl SignatureRegister", "error", err)
		return "", err
	}
	SignatureRegister, _ := ioutil.ReadAll(resp.Body)
	return string(SignatureRegister), nil
}
