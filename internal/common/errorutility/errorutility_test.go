package errorutility_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/common/errorutility"
	"github.com/stretchr/testify/assert"
)

func TestSetHTTPMapping(t *testing.T) {
	t.Run("Error exist", func(t *testing.T) {
		err := errors.New("Error 2")
		statusCode := 111
		errMsg := "Random Number Generated 2"
		mapping := errorutility.NewHTTPMapping(statusCode, errMsg)

		befMapping := errorutility.GetHTTPMapping(err)
		assert.Equal(t, befMapping.StatusCode, http.StatusInternalServerError)

		errorutility.SetHTTPMapping(err, mapping)
		afterMapping := errorutility.GetHTTPMapping(err)
		assert.Equal(t, mapping, afterMapping)
	})
}

func TestGetHTTPMapping(t *testing.T) {
	t.Run("Error not exist in mapping", func(t *testing.T) {
		err := errors.New("non existing error")
		mapping := errorutility.GetHTTPMapping(err)

		assert.Equal(t, mapping.StatusCode, http.StatusInternalServerError)
		assert.Equal(t, mapping.ErrMsg, "Internal Server Error")
	})
	t.Run("Error exist", func(t *testing.T) {
		err := errors.New("Error")
		statusCode := 222
		errMsg := "Random Number Generated"
		mapping := errorutility.NewHTTPMapping(statusCode, errMsg)
		errorutility.SetHTTPMapping(err, mapping)
		retMapping := errorutility.GetHTTPMapping(err)

		assert.Equal(t, mapping, retMapping)
	})
}
func TestNewHTTPMapping(t *testing.T) {
	statusCode := 404
	errMsg := "Not Found"
	mapping := errorutility.NewHTTPMapping(statusCode, errMsg)

	assert.Equal(t, mapping.StatusCode, statusCode)
	assert.Equal(t, mapping.ErrMsg, errMsg)
}
