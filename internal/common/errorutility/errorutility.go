package errorutility

import (
	"net/http"
)

type HTTPMapping struct {
	StatusCode int
	ErrMsg     string
}

var HTTPErrorMap = make(map[error]HTTPMapping)

// SetHTTPMapping set error to HTTPMapping
func SetHTTPMapping(err error, mapping HTTPMapping) {
	HTTPErrorMap[err] = mapping
}

// GetHTTPMapping get HTTPMapping for certain error
func GetHTTPMapping(err error) HTTPMapping {
	if _, ok := HTTPErrorMap[err]; !ok {
		return HTTPMapping{
			StatusCode: http.StatusInternalServerError,
			ErrMsg:     "Internal Server Error",
		}
	}
	return HTTPErrorMap[err]
}

// NewHTTPMapping ...
func NewHTTPMapping(statusCode int, errMsg string) HTTPMapping {
	return HTTPMapping{
		StatusCode: statusCode,
		ErrMsg:     errMsg,
	}
}
