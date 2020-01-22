package httpclient

import (
	"bytes"
	"encoding/json"
	vmFaspay "faspay/faspay_services/view_model"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// methods of the HTTP REST API.
type FaspayService struct {
	client *ClientService
}

// Authentication - authorization
func (s *FaspayService) AuthorizationKey() (*http.Response, error) {
	// build the URL

	u, err := url.Parse(config.GetString(`faspay_key_gen_urls.base_url`) + config.GetString(`faspay_key_gen_urls.authorization`))
	if err != nil {
		fmt.Println("u err  =", err)
		return nil, err
	}

	hc := http.Client{}
	form := url.Values{}
	form.Add("type", "get_token")
	form.Add("faspay_secret", config.GetString(`faspay_keys.Faspay_Secret`))
	form.Add("app_key", config.GetString(`faspay_keys.APP_Key`))
	form.Add("app_secret", config.GetString(`faspay_keys.APP_Secret`))
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Token For signature get token
func (s *FaspayService) TokenForSignature() (*http.Response, error) {
	// build the URL

	u, err := url.Parse(config.GetString(`faspay_key_gen_urls.base_url`) + config.GetString(`faspay_key_gen_urls.access_token`))
	if err != nil {
		fmt.Println("u err  =", err)
		return nil, err
	}

	hc := http.Client{}
	form := url.Values{}
	form.Add("type", "get_token")
	form.Add("client_key", config.GetString(`faspay_keys.Client_Key`))
	form.Add("client_secret", config.GetString(`faspay_keys.Client_Secret`))
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//fmt.Println(form)
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Signature for get token
func (s *FaspayService) SignatureGetAccessToken(SignatureForTokenRequest, time_now string) (*http.Response, error) {
	// build the URL

	u, err := url.Parse(config.GetString(`faspay_key_gen_urls.base_url`) + config.GetString(`faspay_key_gen_urls.signature`))
	if err != nil {
		fmt.Println("u err  =", err)
		return nil, err
	}

	hc := http.Client{}
	form := url.Values{}
	form.Add("type", "gettoken")
	form.Add("method", "GET")
	form.Add("url", "/account/api/register")
	form.Add("timestamp", time_now)
	form.Add("request_body", "")
	form.Add("token", SignatureForTokenRequest)
	form.Add("faspay_secret", config.GetString(`faspay_keys.Faspay_Secret`))
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//fmt.Println(form)
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Get Access Token
func (s *FaspayService) GetAccessToken(opt vmFaspay.TokenRequest) (*http.Response, error) {
	// build the URL

	u, err := url.Parse(config.GetString(`faspay.base_url`) + config.GetString(`faspay.tokens`))
	if err != nil {
		fmt.Println("u err  =", err)
		return nil, err
	}

	//fmt.Println("opt	=", opt)

	hc := http.Client{}
	form := url.Values{}

	req, err := http.NewRequest("GET", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("faspay-key", config.GetString(`faspay_keys.Faspay_Key`))
	req.Header.Add("faspay-timestamp", opt.TimeNow)
	req.Header.Add("faspay-signature", opt.SignatureGetAccessToken)
	req.Header.Add("faspay-authorization", opt.AuthorizationKey)

	//fmt.Println(form)
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Token for other services
func (s *FaspayService) TokenForOtherService(GetAccessToken string) (*http.Response, error) {
	// build the URL

	u, err := url.Parse(config.GetString(`faspay_key_gen_urls.base_url`) + config.GetString(`faspay_key_gen_urls.access_token`))
	if err != nil {
		fmt.Println("u err  =", err)
		return nil, err
	}

	hc := http.Client{}
	form := url.Values{}
	form.Add("type", "validate_token")
	form.Add("client_key", config.GetString(`faspay_keys.Client_Key`))
	form.Add("client_secret", config.GetString(`faspay_keys.Client_Secret`))
	form.Add("validate_token", GetAccessToken)

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//fmt.Println(form)
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Signature Register
func (s *FaspayService) SignatureRegister(TokenForOtherService, time_now string) (*http.Response, error) {
	// build the URL

	u, err := url.Parse(config.GetString(`faspay_key_gen_urls.base_url`) + config.GetString(`faspay_key_gen_urls.signature`))
	if err != nil {
		fmt.Println("u err  =", err)
		return nil, err
	}

	hc := http.Client{}
	form := url.Values{}
	form.Add("type", "ValidateToken")
	form.Add("method", "POST")
	form.Add("url", "/account/api/register")
	form.Add("timestamp", time_now)
	form.Add("request_body", "")
	form.Add("token", TokenForOtherService)
	form.Add("faspay_secret", config.GetString(`faspay_keys.Faspay_Secret`))
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//fmt.Println(form)
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Register
func (s *FaspayService) FaspayRegister(opt vmFaspay.FaspayRegisterRequest) (*http.Response, error) {
	// build the URL
	u, err := url.Parse(config.GetString(`faspay.base_url`) + config.GetString(`faspay.register`))
	if err != nil {
		return nil, err
	}
	hc := http.Client{}

	faspayRegisterRequestBody := vmFaspay.FaspayRegisterRequestBody{

		VirtualAccount:         opt.VirtualAccount,
		BeneficiaryAccount:     opt.BeneficiaryAccount,
		BeneficiaryAccountName: opt.BeneficiaryAccountName,
		BeneficiaryVAName:      opt.BeneficiaryVAName,
		BeneficiaryBankCode:    opt.BeneficiaryBankCode,
		BeneficiaryBankBranch:  opt.BeneficiaryBankBranch,
		BeneficiaryRegionCode:  opt.BeneficiaryRegionCode,
		BeneficiaryCountryCode: opt.BeneficiaryCountryCode,
		BeneficiaryPurposeCode: opt.BeneficiaryPurposeCode,
	}

	json_body, _ := json.Marshal(faspayRegisterRequestBody)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(json_body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("faspay-key", config.GetString(`faspay_keys.Faspay_Key`))
	req.Header.Add("faspay-timestamp", opt.TimeNow)
	req.Header.Add("faspay-signature", opt.SignatureKey)
	req.Header.Add("faspay-authorization", opt.AuthorizationKey)

	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Confirm Register
func (s *FaspayService) FaspayConfirmRegister(opt vmFaspay.FaspayConfirmRegisterRequest) (*http.Response, error) {
	// build the URL
	u, err := url.Parse(config.GetString(`faspay.base_url`) + config.GetString(`faspay.confirm`))
	if err != nil {
		return nil, err
	}

	hc := http.Client{}

	faspayConfirmRegisterRequestBody := vmFaspay.FaspayConfirmRegisterRequestBody{

		VirtualAccount:         opt.VirtualAccount,
		BeneficiaryAccount:     opt.BeneficiaryAccount,
		BeneficiaryAccountName: opt.BeneficiaryAccountName,
		BeneficiaryVAName:      opt.BeneficiaryVAName,
		BeneficiaryBankCode:    opt.BeneficiaryBankCode,
		BeneficiaryBankBranch:  opt.BeneficiaryBankBranch,
		BeneficiaryRegionCode:  opt.BeneficiaryRegionCode,
		BeneficiaryCountryCode: opt.BeneficiaryCountryCode,
		BeneficiaryPurposeCode: opt.BeneficiaryPurposeCode,
		BeneficiaryPhone:       opt.BeneficiaryPhone,
		BeneficiaryEmail:       opt.BeneficiaryEmail,
		BankAccountNumber:      opt.BankAccountNumber,
		BankAccountName:        opt.BankAccountName,
		Confirm:                opt.Confirm,
	}

	json_body, _ := json.Marshal(faspayConfirmRegisterRequestBody)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(json_body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("faspay-key", config.GetString(`faspay_keys.Faspay_Key`))
	req.Header.Add("faspay-timestamp", opt.TimeNow)
	req.Header.Add("faspay-signature", opt.SignatureKey)
	req.Header.Add("faspay-authorization", opt.AuthorizationKey)

	//fmt.Println(form)
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}

// Transfer
func (s *FaspayService) FaspayTransfer(opt vmFaspay.FaspayTransferRequest) (*http.Response, error) {
	// build the URL
	u, err := url.Parse(config.GetString(`faspay.base_url`) + config.GetString(`faspay.transfer`))
	if err != nil {
		return nil, err
	}

	hc := http.Client{}
	faspayTransferRequestBody := vmFaspay.FaspayTransferRequestBody{

		VirtualAccount:            opt.VirtualAccount,
		BeneficiaryVirtualAccount: opt.BeneficiaryVirtualAccount,
		BeneficiaryAccount:        opt.BeneficiaryAccount,
		BeneficiaryName:           opt.BeneficiaryName,
		BeneficiaryBankCode:       opt.BeneficiaryBankCode,
		BeneficiaryRegionCode:     opt.BeneficiaryRegionCode,
		BeneficiaryCountryCode:    opt.BeneficiaryCountryCode,
		BeneficiaryPurposeCode:    opt.BeneficiaryPurposeCode,
		BeneficiaryPhone:          opt.BeneficiaryPhone,
		BeneficiaryEmail:          opt.BeneficiaryEmail,
		TrxNo:                     opt.TrxNo,
		TrxDate:                   opt.TrxDate,
		InstructDate:              opt.InstructDate,
		TrxAmount:                 opt.TrxAmount,
		TrxDesc:                   opt.TrxDesc,
		CallBackURL:               opt.CallBackURL,
	}

	json_body, _ := json.Marshal(faspayTransferRequestBody)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(json_body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("faspay-key", config.GetString(`faspay_keys.Faspay_Key`))
	req.Header.Add("faspay-timestamp", opt.TimeNow)
	req.Header.Add("faspay-signature", opt.SignatureKey)
	req.Header.Add("faspay-authorization", opt.AuthorizationKey)

	//fmt.Println(form)
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("resp err =", err)
		return resp, err
	}
	return resp, err
}
