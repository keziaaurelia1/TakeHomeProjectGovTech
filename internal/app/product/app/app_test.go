package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/app"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/domain/product"
	productMock "github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/domain/product_mock"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/repo"
	"github.com/stretchr/testify/assert"
)

type MockObject struct {
	RepoMock *productMock.MockRepository
	App      app.Application
}

func PrepareTest(mockCtrl *gomock.Controller) *MockObject {
	repoMock := productMock.NewMockRepository(mockCtrl)
	appTest := app.ProvideApplication(repoMock)
	return &MockObject{
		RepoMock: repoMock,
		App:      appTest,
	}
}

func TestProvideApplication(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("When Success", func(t *testing.T) {
		testObj := PrepareTest(mockCtrl)
		testApp := app.ProvideApplication(testObj.RepoMock)
		assert.Equal(t, testObj.App, testApp)
	})
}

func TestCreateProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When Create Product have no error", func(t *testing.T) {
		param := &product.ProductParam{
			ProductID:   1,
			Sku:         "Asj12345",
			Title:       "Product1",
			Description: "Desc 1",
			Category:    "Cat 1",
			Etalase:     "Eta 1",
			Weight:      1.5,
		}
		testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(nil, repo.ErrProductNotFound)
		testObj.RepoMock.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).Return(nil)

		err := testObj.App.CreateProduct(context.Background(), param)
		assert.NoError(t, err)
	})

	t.Run("When Create have error", func(t *testing.T) {
		t.Run("Error in Find by Product ID", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

			err := testObj.App.CreateProduct(context.Background(), param)
			assert.Error(t, err)
		})
		t.Run("Error because Find by Product ID get something", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(&product.Product{}, nil)

			err := testObj.App.CreateProduct(context.Background(), param)
			assert.Error(t, err)
		})
		t.Run("Error in Insert", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(nil, repo.ErrProductNotFound)
			testObj.RepoMock.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).Return(errors.New("error"))

			err := testObj.App.CreateProduct(context.Background(), param)
			assert.Error(t, err)
		})
	})
}

func TestUpdateProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When UpdateProduct have no error", func(t *testing.T) {
		param := &product.ProductParam{
			ProductID:   2,
			Sku:         "Asj12345",
			Title:       "Product1",
			Description: "Desc 1",
			Category:    "Cat 1",
			Etalase:     "Eta 1",
			Weight:      1.5,
		}
		testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(&product.Product{}, nil)
		testObj.RepoMock.EXPECT().UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		err := testObj.App.UpdateProduct(context.Background(), 2, param)
		assert.NoError(t, err)
	})
	t.Run("When UpdateProduct have error", func(t *testing.T) {
		t.Run("Error in Find Product ID", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   2,
				Sku:         "Asj12345",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			err := testObj.App.UpdateProduct(context.Background(), 2, param)
			assert.Error(t, err)
		})
		t.Run("Error in Update Product", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   2,
				Sku:         "Asj12345",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(&product.Product{}, nil)
			testObj.RepoMock.EXPECT().UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			err := testObj.App.UpdateProduct(context.Background(), 2, param)
			assert.Error(t, err)
		})
		t.Run("Error in Updated product ID already exist ", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   2,
				Sku:         "Asj12345",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}

			testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(&product.Product{}, nil)
			testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(&product.Product{}, nil)
			err := testObj.App.UpdateProduct(context.Background(), 3, param)
			assert.Error(t, err)
		})
	})

}

func TestFindProductByProductID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When FindProductByProductID have no error", func(t *testing.T) {
		testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(&product.Product{}, nil)
		_, err := testObj.App.FindProductByProductID(context.Background(), 2)
		assert.NoError(t, err)
	})
	t.Run("When FindProductByProductID have error", func(t *testing.T) {
		t.Run("Error in Find Product ID", func(t *testing.T) {
			testObj.RepoMock.EXPECT().FindByProductID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := testObj.App.FindProductByProductID(context.Background(), 2)
			assert.Error(t, err)
		})
	})

}
func TestFindProductBySku(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When FindProductBySku have no error", func(t *testing.T) {
		testObj.RepoMock.EXPECT().FindBySku(gomock.Any(), gomock.Any()).Return(&product.Product{}, nil)
		_, err := testObj.App.FindProductBySku(context.Background(), "abc12383")
		assert.NoError(t, err)
	})
	t.Run("When FindProductBySku have error", func(t *testing.T) {
		t.Run("Error in FindBySku", func(t *testing.T) {
			testObj.RepoMock.EXPECT().FindBySku(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := testObj.App.FindProductBySku(context.Background(), "abc12383")
			assert.Error(t, err)
		})
	})
}
func TestFindProductByTitleCategoryEtalase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When FindProductByTitleCategoryEtalase have no error", func(t *testing.T) {
		search := &product.SearchParam{
			Title:    "a",
			Category: "b",
			Etalase:  "c",
		}
		testObj.RepoMock.EXPECT().FindByTitleCategoryEtalase(gomock.Any(), gomock.Any()).Return([]*product.Product{{}}, nil)
		_, err := testObj.App.FindProductByTitleCategoryEtalase(context.Background(), search)
		assert.NoError(t, err)
	})
	t.Run("When FindProductByTitleCategoryEtalase have error", func(t *testing.T) {
		t.Run("Error in FindByTitleCategoryEtalase", func(t *testing.T) {
			search := &product.SearchParam{
				Title:    "a",
				Category: "b",
				Etalase:  "c",
			}
			testObj.RepoMock.EXPECT().FindByTitleCategoryEtalase(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := testObj.App.FindProductByTitleCategoryEtalase(context.Background(), search)
			assert.Error(t, err)
		})
	})

}

func TestFindReviewByDate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When FindReviewByDate have no error", func(t *testing.T) {
		testObj.RepoMock.EXPECT().FindByDate(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*product.Review{{}}, nil)

		_, err := testObj.App.FindReviewByDate(context.Background(), "2023-02-12", 1)
		assert.NoError(t, err)
	})
	t.Run("When FindReviewByDate have error", func(t *testing.T) {
		t.Run("Error in FindReviewByDate", func(t *testing.T) {
			testObj.RepoMock.EXPECT().FindByDate(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

			_, err := testObj.App.FindReviewByDate(context.Background(), "2023-02-12", 1)
			assert.Error(t, err)
		})
	})
}

func TestGetAllProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When GetAllProduct have no error", func(t *testing.T) {
		testObj.RepoMock.EXPECT().AllProduct(gomock.Any()).Return([]*product.Product{{}}, nil)

		_, err := testObj.App.GetAllProduct(context.Background())
		assert.NoError(t, err)
	})
	t.Run("When GetAllProduct have error", func(t *testing.T) {
		t.Run("Error in AllProduct", func(t *testing.T) {
			testObj.RepoMock.EXPECT().AllProduct(gomock.Any()).Return(nil, errors.New("error"))
			_, err := testObj.App.GetAllProduct(context.Background())
			assert.Error(t, err)
		})
	})
}

func TestGetAllReview(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When GetAllReview have no error", func(t *testing.T) {
		testObj.RepoMock.EXPECT().AllReview(gomock.Any(), gomock.Any()).Return([]*product.Review{{}}, nil)

		_, err := testObj.App.GetAllReview(context.Background(), 1)
		assert.NoError(t, err)
	})
	t.Run("When GetAllReview have error", func(t *testing.T) {
		t.Run("Error in AllReview", func(t *testing.T) {
			testObj.RepoMock.EXPECT().AllReview(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := testObj.App.GetAllReview(context.Background(), 1)
			assert.Error(t, err)
		})
	})
}

func TestCreateReview(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testObj := PrepareTest(mockCtrl)
	t.Run("When Create Review have no error", func(t *testing.T) {
		param := &product.ReviewParam{
			ProductID:      1,
			ReviewID:       1,
			Rating:         2,
			ReviewComment:  "Comment 1",
			DateTimeReview: "2023-02-12",
		}
		testObj.RepoMock.EXPECT().FindByReviewID(gomock.Any(), gomock.Any()).Return(nil, repo.ErrReviewNotFound)
		testObj.RepoMock.EXPECT().InsertReview(gomock.Any(), gomock.Any()).Return(nil)

		err := testObj.App.CreateReview(context.Background(), param)
		assert.NoError(t, err)
	})

	t.Run("When Create Review have error", func(t *testing.T) {
		t.Run("Error in Find by Review ID", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      1,
				ReviewID:       1,
				Rating:         2,
				ReviewComment:  "Comment 1",
				DateTimeReview: "2023-02-12",
			}
			testObj.RepoMock.EXPECT().FindByReviewID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			err := testObj.App.CreateReview(context.Background(), param)
			assert.Error(t, err)
		})
		t.Run("Error because Find by Review ID get something", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      1,
				ReviewID:       1,
				Rating:         2,
				ReviewComment:  "Comment 1",
				DateTimeReview: "2023-02-12",
			}
			testObj.RepoMock.EXPECT().FindByReviewID(gomock.Any(), gomock.Any()).Return(&product.Review{}, nil)

			err := testObj.App.CreateReview(context.Background(), param)
			assert.Error(t, err)
		})
		t.Run("Error in Insert", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      1,
				ReviewID:       1,
				Rating:         2,
				ReviewComment:  "Comment 1",
				DateTimeReview: "2023-02-12",
			}
			testObj.RepoMock.EXPECT().FindByReviewID(gomock.Any(), gomock.Any()).Return(nil, repo.ErrReviewNotFound)
			testObj.RepoMock.EXPECT().InsertReview(gomock.Any(), gomock.Any()).Return(errors.New("error"))

			err := testObj.App.CreateReview(context.Background(), param)
			assert.Error(t, err)
		})
	})
}
