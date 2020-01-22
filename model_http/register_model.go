package modelhttp

import (
	"context"
	"encoding/json"
	helperInputValidateFaspay "faspay/faspay_services/helper/input_validate"
	modelFaspay "faspay/faspay_services/view_model"
	"net/http"
)

// Tokens
func DecodeFaspayTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t modelFaspay.EmptyVM
	return t, nil
}

// Register
func DecodeFaspayRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t modelFaspay.FaspayRegisterRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}
	return t, nil
}

// Confirm Register
func DecodeFaspayConfirmRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t modelFaspay.FaspayConfirmRegisterRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}
	return t, nil
}

// Transfer
func DecodeFaspayTransferRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var t modelFaspay.FaspayTransferRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return nil, ErrBadRouting
	}
	errv := helperInputValidateFaspay.InputValidateStruct(t)
	if errv != nil {
		return nil, errv
	}
	return t, nil
}
