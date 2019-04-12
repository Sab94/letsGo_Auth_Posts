package types

type APIErrors struct {
	Errors      []*APIError `json:"errors"`
}

func (errors *APIErrors) Status() int {
	return errors.Errors[0].Status
}

type APIError struct {
	Status      int         `json:"status"`
	Code        string      `json:"code"`
	Title       string      `json:"title"`
	Details     string      `json:"details"`
	Href        string      `json:"href"`
}

func newAPIError(status int, code string, title string, details string, href string) *APIError {
	return &APIError{
		Status:     status,
		Code:       code,
		Title:      title,
		Details:    details,
		Href:       href,
	}
}

var (
	ErrLogin     = newAPIError(500, "password_mismatch", "Password mismatch", "Password does not match.", "")
	ErrUnknown     = newAPIError(500, "something_went_wrong", "Something went wrong", "something went wrong.", "")
)