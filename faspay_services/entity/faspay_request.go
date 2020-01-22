package entity

type MidtransChargeRequest struct {
	MemberCode      string `json:"member_code"`
	Amount          int    `json:"amount"`
	BankName        string `json:"bank_name"`
	MemberFirstName string `json:"member_first_name"`
	MemberLastName  string `json:"member_last_name"`
	MemberPhone     string `json:"member_phone"`
	MemberEmail     string `json:"member_email"`
}
