package paymentservice

import "errors"

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting                = errors.New("inconsistent mapping between route and handler")
	ErrRequiredNewPassword       = errors.New("Required new password")
	ErrRequiredConfirmedPassword = errors.New("Required confirmed password")
)

// errors : custom types
var (
	ErrRequiredAuthorizationToken        = errors.New("Required authorization token")
	ErrUnauthorizedError                 = errors.New("Unauthorized Error")
	ErrRequiredPlayerID                  = errors.New("Required player id")
	ErrInternalServerError               = errors.New("Internal Server Error")
	InvalidLoginError                    = errors.New("Invalid user name or password")
	ErrRequiredUserName                  = errors.New("Required user name")
	PwdDoesNotMatchError                 = errors.New("Password does not match")
	ConfirmPwdDoesNotMatchError          = errors.New("Confirmation Password does not match")
	ErrNotFoundError                     = errors.New("Your requested item is not found")
	ItemConflictError                    = errors.New("Your item already exists")
	NameConflictError                    = errors.New("Your name already exists")
	OrgNameConflictError                 = errors.New("Your org name already exists")
	OrgIDNumConflictError                = errors.New("Your org id number already exists")
	UserNameConflictError                = errors.New("Your user name already exists")
	MobileConflictError                  = errors.New("Your mobile already exists")
	PhoneConflictError                   = errors.New("Phone number already exists")
	EmailConflictError                   = errors.New("Your email already exists")
	SocialMediaIDConflictError           = errors.New("Your social media id already exists")
	ReferIDConflictError                 = errors.New("Your refer id already exists")
	IDNumConflictError                   = errors.New("Your id number already exists")
	ReferIDNotExistError                 = errors.New("Referral code is not valid")
	MobileRegisteredError                = errors.New("The phone number is registered")
	PhoneTokenNotExistError              = errors.New("Phone token is not valid")
	ForgotPasswordTokenNotExistError     = errors.New("Forgot Password token is not valid")
	ForgotPasswordSignatureNotExistError = errors.New("Forgot Password signature is not valid")
	MemberisnotactiveError               = errors.New("Member is not active")
	MemberisnotExistsError               = errors.New("Member does not exists")
	TransactionisduplicateError          = errors.New("Transaction is duplicate")
	CashierNotFoundError                 = errors.New("Cashier Not Found")
	MemberAllreadyExists                 = errors.New("Member already exists")
	CashierAlreadyExist                  = errors.New("Cashier already exists")
	ThisShipmentHasBeenRefund            = errors.New("This shipment has been refund")
	FailedToGetCourierMoney              = errors.New("Failed to get Courier Money")
	ShipmentIsPending                    = errors.New("Shipment is pending")
)

var (
	InternalServerError           = "Internal Server Error"
	StatusUnKnown                 = "UNKNOWN"
	StatusOk                      = "OK"
	SuccessCodeType               = "Success"
	BaseTypeUrl                   = "paxel.co/"
	MessageRegisterSuccess        = "Register Successfully"
	MessageRegisterConfirmSuccess = "Register Confirm Successfully"
	MessageTransferSuccess        = "Transfer Successfully"
	MessageLoginSuccess           = "Login Successfully"
	MessageCreatedSuccess         = "Created Successfully"
	MessageRetrievedSuccess       = "Retrieved Successfully"
	MessageRetrievedListSuccess   = "Retrieved List Successfully"
	MessageUpdatedSuccess         = "Updated Successfully"
	MessageDeletedSuccess         = "Deleted Successfully"
	MessageValidatedSuccess       = "Validated Successfully"
	MessageSentSuccess            = "Sent Successfully"
	MessageSuccess                = "Success"
)
