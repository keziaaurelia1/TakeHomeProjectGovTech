package app

import (
	"context"

	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/domain/product"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/repo"
)

type (
	Application interface {
		CreateProduct(context.Context, *product.ProductParam) error
		UpdateProduct(context.Context, int64, *product.ProductParam) error
		FindProductByProductID(context.Context, int64) (*product.Product, error)
		FindProductBySku(context.Context, string) (*product.Product, error)
		FindProductByTitleCategoryEtalase(context.Context, *product.SearchParam) ([]*product.Product, error)
		GetAllProduct(context.Context) ([]*product.Product, error)

		FindReviewByDate(context.Context, string, int64) ([]*product.Review, error)
		GetAllReview(context.Context, int64) ([]*product.Review, error)
		CreateReview(context.Context, *product.ReviewParam) error
	}
	ApplicationImpl struct {
		repo product.Repository
	}
)

func ProvideApplication(repo product.Repository) Application {
	return &ApplicationImpl{repo: repo}
}

func (impl *ApplicationImpl) CreateProduct(ctx context.Context, param *product.ProductParam) error {
	newProduct, err := product.NewProduct(param)
	if err != nil {
		return err
	}
	_, err = impl.repo.FindByProductID(ctx, param.ProductID)
	if err == nil {
		return ErrProductIDExist
	}
	if err != repo.ErrProductNotFound {
		return err
	}
	err = impl.repo.InsertProduct(ctx, newProduct)
	return err
}

func (impl *ApplicationImpl) CreateReview(ctx context.Context, param *product.ReviewParam) error {
	newReview, err := product.NewReview(param)
	if err != nil {
		return err
	}
	_, err = impl.repo.FindByReviewID(ctx, param.ReviewID)
	if err == nil {
		return ErrReviewIDExist
	}
	if err != repo.ErrReviewNotFound {
		return err
	}
	err = impl.repo.InsertReview(ctx, newReview)
	return err
}

func (impl *ApplicationImpl) GetAllProduct(ctx context.Context) ([]*product.Product, error) {
	listProduct, err := impl.repo.AllProduct(ctx)
	if err != nil {
		return nil, err
	}
	return listProduct, err
}

func (impl *ApplicationImpl) UpdateProduct(ctx context.Context, prevProductID int64, param *product.ProductParam) error {
	newProduct, err := product.NewProduct(param)
	if err != nil {
		return err
	}
	_, err = impl.repo.FindByProductID(ctx, prevProductID)
	if err != nil {
		return err
	}
	if prevProductID != param.ProductID {
		_, err = impl.repo.FindByProductID(ctx, param.ProductID)
		if err == nil {
			return ErrUpdatedProductIDExist
		}
		if err != repo.ErrProductNotFound {
			return err
		}
	}
	err = impl.repo.UpdateProduct(ctx, prevProductID, newProduct)
	return err
}

func (impl *ApplicationImpl) GetAllReview(ctx context.Context, productID int64) ([]*product.Review, error) {
	listReview, err := impl.repo.AllReview(ctx, productID)
	if err != nil {
		return nil, err
	}
	return listReview, err
}

func (impl *ApplicationImpl) FindProductByProductID(ctx context.Context, productID int64) (*product.Product, error) {
	resProduct, err := impl.repo.FindByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	return resProduct, err
}
func (impl *ApplicationImpl) FindProductBySku(ctx context.Context, sku string) (*product.Product, error) {
	resProduct, err := impl.repo.FindBySku(ctx, sku)
	if err != nil {
		return nil, err
	}
	return resProduct, err
}
func (impl *ApplicationImpl) FindProductByTitleCategoryEtalase(ctx context.Context, search *product.SearchParam) ([]*product.Product, error) {
	listProduct, err := impl.repo.FindByTitleCategoryEtalase(ctx, search)
	if err != nil {
		return nil, err
	}
	return listProduct, err
}
func (impl *ApplicationImpl) FindReviewByDate(ctx context.Context, tanggal string, reviewID int64) ([]*product.Review, error) {
	resReview, err := impl.repo.FindByDate(ctx, tanggal, reviewID)
	if err != nil {
		return nil, err
	}
	return resReview, err
}
