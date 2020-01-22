package viewmodel

type FaspayConfirmRegisterRequest struct {
	VirtualAccount         string `json:"virtual_account"`
	BeneficiaryAccount     string `json:"beneficiary_account"`
	BeneficiaryAccountName string `json:"beneficiary_account_name"`
	BeneficiaryVAName      string `json:"beneficiary_va_name"`
	BeneficiaryBankCode    string `json:"beneficiary_bank_code"`
	BeneficiaryBankName    string `json:"beneficiary_bank_name"`
	BeneficiaryBankBranch  string `json:"beneficiary_bank_branch"`
	BeneficiaryRegionCode  string `json:"beneficiary_region_code"`
	BeneficiaryCountryCode string `json:"beneficiary_country_code"`
	BeneficiaryPurposeCode string `json:"beneficiary_purpose_code"`
	BeneficiaryPhone       string `json:"beneficiary_phone"`
	BeneficiaryEmail       string `json:"beneficiary_email"`
	BankAccountNumber      string `json:"bank_account_number"`
	BankAccountName        string `json:"bank_account_name"`
	Confirm                string `json:"confirm"`
	SignatureKey           string
	AuthorizationKey       string
	TimeNow                string
}

type FaspayConfirmRegisterRequestBody struct {
	VirtualAccount         string `json:"virtual_account"`
	BeneficiaryAccount     string `json:"beneficiary_account"`
	BeneficiaryAccountName string `json:"beneficiary_account_name"`
	BeneficiaryVAName      string `json:"beneficiary_va_name"`
	BeneficiaryBankCode    string `json:"beneficiary_bank_code"`
	BeneficiaryBankName    string `json:"beneficiary_bank_name"`
	BeneficiaryBankBranch  string `json:"beneficiary_bank_branch"`
	BeneficiaryRegionCode  string `json:"beneficiary_region_code"`
	BeneficiaryCountryCode string `json:"beneficiary_country_code"`
	BeneficiaryPurposeCode string `json:"beneficiary_purpose_code"`
	BeneficiaryPhone       string `json:"beneficiary_phone"`
	BeneficiaryEmail       string `json:"beneficiary_email"`
	BankAccountNumber      string `json:"bank_account_number"`
	BankAccountName        string `json:"bank_account_name"`
	Confirm                string `json:"confirm"`
}
