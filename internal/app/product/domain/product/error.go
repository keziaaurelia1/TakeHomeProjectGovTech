package product

import "errors"

var (
	ErrInvalidSku         = errors.New("Nomor Sku tidak valid")
	ErrTitleTooLong       = errors.New("Title > 100")
	ErrDescriptionTooLong = errors.New("Description >1000")
	ErrCategoryTooLong    = errors.New("Category > 50")
	ErrEtalaseTooLong     = errors.New("Etalase > 50")
	ErrWeightNotPositive  = errors.New("Weight <= 0")

	ErrInvalidImageID = errors.New("Image ID tidak valid")
	ErrInvalidPath    = errors.New("Path tidak valid")

	ErrInvalidReviewID      = errors.New("Review ID tidak valid")
	ErrInvalidRating        = errors.New("Rating tidak valid")
	ErrReviewCommentTooLong = errors.New("Review comment > 1000")
	ErrInvalidProductID     = errors.New("Product ID tidak valid")

	ErrInvalidDateTime = errors.New("Date Time tidak valid")
	ErrInvalidPrice    = errors.New("Harga tidak valid")

	MsgErrInvalidPrice       = "Harga yang tertera tidak valid..."
	MsgErrInvalidDateTime    = "Tanggal tidak valid..."
	MsgErrInvalidProductID   = "Product ID tidak valid..."
	MsgErrInvalidSku         = "Nomor Sku tidak valid..."
	MsgErrTitleTooLong       = "Title melebihi batas maximum(100) character..."
	MsgErrDescriptionTooLong = "Description melebihi batas maximum(1000) character..."
	MsgErrCategoryTooLong    = "Category melebihi batas maximum(50) character..."
	MsgErrEtalaseTooLong     = "Etalase melebihi batas maximum(50) character..."
	MsgErrWeightNotPositive  = "Berat yang dimasukan kurang dari 0..."

	MsgErrInvalidImageID = "Image ID tidak valid..."
	MsgErrInvalidPath    = "Path tidak valid..."

	MsgErrInvalidReviewID      = "Review ID tidak valid..."
	MsgErrInvalidRating        = "Rating tidak valid..."
	MsgErrReviewCommentTooLong = "Review comment melebihi batas maximum(1000) character..."
)
