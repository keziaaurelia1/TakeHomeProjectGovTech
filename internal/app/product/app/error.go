package app

import "errors"

var (
	ErrProductIDExist        = errors.New("Error Product ID sudah ada")
	ErrReviewIDExist         = errors.New("Error Review ID sudah ada")
	ErrUpdatedProductIDExist = errors.New("Updated Product ID sudah ada")

	MsgErrProductIDExist        = "Product ID sudah pernah ada..."
	MsgErrReviewIDExist         = "Review ID sudah pernah ada..."
	MsgErrUpdatedProductIDExist = "Product ID yang mau diganti sudah ada..."
)
