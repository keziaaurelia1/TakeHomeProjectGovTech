package port

import "errors"

var (
	ErrMissingParam = errors.New("Error missing param")

	MsgMissingParam           = "Salah satu field tidak terisi, silakan coba lagi..."
	MsgErrInternalServerError = "There is an issue in our system, please retry in a few seconds..."
)
