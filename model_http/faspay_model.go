package modelhttp

import (
	"context"
	"encoding/json"
	"errors"
	serviceFaspay "faspay/faspay_services"
	"net/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler")
)

// Errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type Errorer interface {
	error() error
}

// EncodeResponse is the common method to encode all response types to the client.
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(Errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	/*paymentResponse := vmPayment.PaymentResponse{
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
		CodeType: serviceFaspay.SuccessCodeType,
		Message:  http.StatusText(http.StatusOK),
		Data:     response,
	}
	return json.NewEncoder(w).Encode(paymentResponse)*/
	return json.NewEncoder(w).Encode(response)
}

// encodeError - Provide those as HTTP errors
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("EncodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case serviceFaspay.ErrNotFoundError:
		return http.StatusNotFound
	case serviceFaspay.ErrBadRouting,
		serviceFaspay.ErrRequiredAuthorizationToken,
		serviceFaspay.ErrRequiredPlayerID:
		return http.StatusBadRequest
	case serviceFaspay.ErrUnauthorizedError:
		return http.StatusUnauthorized
	case serviceFaspay.UserNameConflictError,
		serviceFaspay.PhoneConflictError,
		serviceFaspay.EmailConflictError,
		serviceFaspay.SocialMediaIDConflictError,
		serviceFaspay.ReferIDConflictError:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
