package product

import "context"

type Repository interface {
	FindByProductID(ctx context.Context, productID int64) (*Product, error)
	FindBySku(ctx context.Context, sku string) (*Product, error)
	FindByTitleCategoryEtalase(ctx context.Context, search *SearchParam) ([]*Product, error)
	InsertProduct(ctx context.Context, entProduct *Product) error
	UpdateProduct(ctx context.Context, prevProductID int64, entProduct *Product) error
	AllProduct(ctx context.Context) ([]*Product, error)

	InsertReview(ctx context.Context, entReview *Review) error
	FindByDate(ctx context.Context, tanggal string, reviewID int64) ([]*Review, error)
	FindByReviewID(ctx context.Context, reviewID int64) (*Review, error)
	AllReview(ctx context.Context, productID int64) ([]*Review, error)
}
