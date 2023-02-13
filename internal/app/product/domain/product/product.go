package product

import (
	"strconv"
	"strings"
	"time"
)

type (
	Product struct {
		ProductID   int64   `json:"product_id"`
		Sku         string  `json:"sku"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		Etalase     string  `json:"etalase"`
		Weight      float64 `json:"weight"`
		Price       int64   `json:"price"`
	}
	ProductParam struct {
		ProductID   int64   `json:"product_id"`
		Sku         string  `json:"sku"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		Etalase     string  `json:"etalase"`
		Weight      float64 `json:"weight"`
		Price       int64   `json:"price"`
	}
	Image struct {
		ProductID   int64  `json:"product_id"`
		ImageID     int64  `json:"image_id"`
		Path        string `json:"path"`
		Description string `json:"description"`
	}
	ImageParam struct {
		ProductID   int64  `json:"product_id"`
		ImageID     int64  `json:"image_id"`
		Path        string `json:"path"`
		Description string `json:"description"`
	}
	Review struct {
		ProductID      int64   `json:"product_id"`
		ReviewID       int64   `json:"review_id"`
		Rating         float32 `json:"rating"`
		ReviewComment  string  `json:"review_comment"`
		DateTimeReview string  `json:"date_time_review"`
	}
	ReviewParam struct {
		ProductID      int64   `json:"product_id"`
		ReviewID       int64   `json:"review_id"`
		Rating         float32 `json:"rating"`
		ReviewComment  string  `json:"review_comment"`
		DateTimeReview string  `json:"date_time_review"`
	}
	SearchParam struct {
		Title    string
		Category string
		Etalase  string
	}
)

func ConvertToSearch(param *Product) (*SearchParam, error) {
	return &SearchParam{
		Title:    param.Title,
		Category: param.Category,
		Etalase:  param.Etalase,
	}, nil
}

func NewProduct(param *ProductParam) (*Product, error) {
	err := validateProductParam(param)
	if err != nil {
		return nil, err
	}
	return &Product{
		ProductID:   param.ProductID,
		Sku:         param.Sku,
		Title:       param.Title,
		Description: param.Description,
		Category:    param.Category,
		Etalase:     param.Etalase,
		Weight:      param.Weight,
		Price:       param.Price,
	}, nil
}
func NewImage(param *ImageParam) (*Image, error) {
	err := validateImageParam(param)
	if err != nil {
		return nil, err
	}
	return &Image{
		ProductID:   param.ProductID,
		ImageID:     param.ImageID,
		Path:        param.Path,
		Description: param.Description,
	}, nil
}
func NewReview(param *ReviewParam) (*Review, error) {
	err := validateReviewParam(param)
	if err != nil {
		return nil, err
	}
	return &Review{
		ProductID:      param.ProductID,
		ReviewID:       param.ReviewID,
		Rating:         param.Rating,
		ReviewComment:  param.ReviewComment,
		DateTimeReview: param.DateTimeReview,
	}, nil
}

func validateProductParam(param *ProductParam) error {
	err := validateProductID(param.ProductID)
	if err != nil {
		return err
	}
	err = validateSku(param.Sku)
	if err != nil {
		return err
	}
	err = validateTitle(param.Title)
	if err != nil {
		return err
	}
	err = validateDescription(param.Description)
	if err != nil {
		return err
	}
	err = validateCategory(param.Category)
	if err != nil {
		return err
	}
	err = validateEtalase(param.Etalase)
	if err != nil {
		return err
	}
	err = validateWeight(param.Weight)
	if err != nil {
		return err
	}
	err = validatePrice(param.Price)
	if err != nil {
		return err
	}
	return nil
}
func validatePrice(price int64) error {
	if price <= 0 {
		return ErrInvalidPrice
	}
	return nil
}
func validateProductID(productID int64) error {
	if productID < 0 {
		return ErrInvalidProductID
	}
	return nil
}
func validateImageParam(param *ImageParam) error {
	err := validateProductID(param.ProductID)
	if err != nil {
		return err
	}
	err = validateImageID(param.ImageID)
	if err != nil {
		return err
	}
	err = validatePath(param.Path)
	if err != nil {
		return err
	}
	err = validateDescription(param.Description)
	if err != nil {
		return err
	}
	return nil
}
func validateImageID(imageID int64) error {
	if imageID < 0 {
		return ErrInvalidImageID
	}
	return nil
}

func validatePath(path string) error {
	if len(path) == 0 {
		return ErrInvalidPath
	}
	return nil
}

func validateReviewParam(param *ReviewParam) error {
	err := validateProductID(param.ProductID)
	if err != nil {
		return err
	}
	err = validateReviewID(param.ReviewID)
	if err != nil {
		return err
	}
	err = validateRating(param.Rating)
	if err != nil {
		return err
	}
	err = validateReviewComment(param.ReviewComment)
	if err != nil {
		return err
	}
	err = validateDateTime(param.DateTimeReview)
	if err != nil {
		return err
	}
	return nil
}
func validateDateTime(dateTimeReview string) error {
	if len(dateTimeReview) != 10 {
		return ErrInvalidDateTime
	}
	dateTimeSplit := strings.Split(dateTimeReview, "-")
	if len(dateTimeSplit) != 3 {
		return ErrInvalidDateTime
	}
	year, err := toNumber(dateTimeSplit[0])
	if err != nil {
		return ErrInvalidDateTime
	}
	month, err := toNumber(dateTimeSplit[1])
	if err != nil || month > 12 {
		return ErrInvalidDateTime
	}
	day, err := toNumber(dateTimeSplit[2])
	if err != nil {
		return ErrInvalidDateTime
	}
	tmpDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if tmpDate.Year() != year || tmpDate.Month() != time.Month(month) || tmpDate.Day() != day {
		return ErrInvalidDateTime
	}
	return nil
}
func toNumber(x string) (int, error) {
	num, err := strconv.ParseInt(x, 10, 64)
	return int(num), err
}
func validateReviewID(reviewID int64) error {
	if reviewID < 0 {
		return ErrInvalidReviewID
	}
	return nil
}
func validateRating(rating float32) error {
	if rating > 10 {
		return ErrInvalidRating
	}
	return nil
}
func validateReviewComment(reviewComment string) error {
	if len(reviewComment) > 1000 {
		return ErrReviewCommentTooLong
	}
	return nil
}

func validateSku(sku string) error {
	if len(sku) != 8 {
		return ErrInvalidSku
	}
	return nil
}

func validateTitle(title string) error {
	if len(title) > 100 {
		return ErrTitleTooLong
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) > 1000 {
		return ErrDescriptionTooLong
	}
	return nil
}
func validateCategory(category string) error {
	if len(category) > 50 {
		return ErrCategoryTooLong
	}
	return nil
}
func validateEtalase(etalase string) error {
	if len(etalase) > 50 {
		return ErrEtalaseTooLong
	}
	return nil
}

func validateImageIDs(imageIDs []Image) error {
	return nil
}
func validateWeight(weight float64) error {
	if weight <= 0 {
		return ErrWeightNotPositive
	}
	return nil
}
func validateReviewIDs(reviewIDs []Review) error {
	return nil
}
