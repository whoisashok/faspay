package viewmodel

type FaspayTransferRequest struct {
	VirtualAccount            string `json:"virtual_account"`
	BeneficiaryVirtualAccount string `json:"beneficiary_virtual_account"`
	BeneficiaryAccount        string `json:"beneficiary_account"`
	BeneficiaryName           string `json:"beneficiary_name"`
	BeneficiaryBankCode       string `json:"beneficiary_bank_code"`
	BeneficiaryRegionCode     string `json:"beneficiary_region_code"`
	BeneficiaryCountryCode    string `json:"beneficiary_country_code"`
	BeneficiaryPurposeCode    string `json:"beneficiary_purpose_code"`
	BeneficiaryPhone          string `json:"beneficiary_phone"`
	BeneficiaryEmail          string `json:"beneficiary_email"`
	TrxNo                     string `json:"trx_no"`
	TrxDate                   string `json:"trx_date"`
	InstructDate              string `json:"instruct_date"`
	TrxAmount                 string `json:"trx_amount"`
	TrxDesc                   string `json:"trx_desc"`
	CallBackURL               string `json:"callback_url" value:"https://dev2.faspay.co.id/account/api/callback`
	SignatureKey              string
	AuthorizationKey          string
	TimeNow                   string
}

type FaspayTransferRequestBody struct {
	VirtualAccount            string `json:"virtual_account"`
	BeneficiaryVirtualAccount string `json:"beneficiary_virtual_account"`
	BeneficiaryAccount        string `json:"beneficiary_account"`
	BeneficiaryName           string `json:"beneficiary_name"`
	BeneficiaryBankCode       string `json:"beneficiary_bank_code"`
	BeneficiaryRegionCode     string `json:"beneficiary_region_code"`
	BeneficiaryCountryCode    string `json:"beneficiary_country_code"`
	BeneficiaryPurposeCode    string `json:"beneficiary_purpose_code"`
	BeneficiaryPhone          string `json:"beneficiary_phone"`
	BeneficiaryEmail          string `json:"beneficiary_email"`
	TrxNo                     string `json:"trx_no"`
	TrxDate                   string `json:"trx_date"`
	InstructDate              string `json:"instruct_date"`
	TrxAmount                 string `json:"trx_amount"`
	TrxDesc                   string `json:"trx_desc"`
	CallBackURL               string `json:"callback_url" value:"https://dev2.faspay.co.id/account/api/callback`
}
