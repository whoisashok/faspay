package viewmodel

type FaspayRegisterRequest struct {
	VirtualAccount         string `json:"virtual_account"`
	BeneficiaryAccount     string `json:"beneficiary_account"`
	BeneficiaryAccountName string `json:"beneficiary_account_name"`
	BeneficiaryVAName      string `json:"beneficiary_va_name"`
	BeneficiaryBankCode    string `json:"beneficiary_bank_code"`
	BeneficiaryBankBranch  string `json:"beneficiary_bank_branch"`
	BeneficiaryRegionCode  string `json:"beneficiary_region_code"`
	BeneficiaryCountryCode string `json:"beneficiary_country_code"`
	BeneficiaryPurposeCode string `json:"beneficiary_purpose_code"`
	SignatureKey           string
	AuthorizationKey       string
	TimeNow                string
}

type FaspayRegisterRequestBody struct {
	VirtualAccount         string `json:"virtual_account"`
	BeneficiaryAccount     string `json:"beneficiary_account"`
	BeneficiaryAccountName string `json:"beneficiary_account_name"`
	BeneficiaryVAName      string `json:"beneficiary_va_name"`
	BeneficiaryBankCode    string `json:"beneficiary_bank_code"`
	BeneficiaryBankBranch  string `json:"beneficiary_bank_branch"`
	BeneficiaryRegionCode  string `json:"beneficiary_region_code"`
	BeneficiaryCountryCode string `json:"beneficiary_country_code"`
	BeneficiaryPurposeCode string `json:"beneficiary_purpose_code"`
}
