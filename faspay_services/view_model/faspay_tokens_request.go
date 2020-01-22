package viewmodel

type EmptyVM struct {
}

type TokenRequest struct {
	SignatureGetAccessToken string
	AuthorizationKey        string
	TimeNow                 string
}

type FaspayTokenResponse struct {
	SignatureKey     string
	AuthorizationKey string
	TimeNow          string
}
