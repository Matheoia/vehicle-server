package vehicle

import "github.com/Matheoia/vehicle-server/pkg/httputil"

func newValidationError(issues []string) error {
	return &httputil.APIError{
		Code:    httputil.ErrCodeInvalidRequestPayload,
		Message: "The request payload is invalid",
		Details: issues,
	}
}