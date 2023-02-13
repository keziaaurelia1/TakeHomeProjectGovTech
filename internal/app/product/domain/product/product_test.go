package product_test

import (
	"testing"

	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/domain/product"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	t.Run("When NewProduct Has No Error", func(t *testing.T) {
		t.Run("Common Case", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			pro, err := product.NewProduct(param)
			assert.NoError(t, err)
			assert.Equal(t, pro.ProductID, int64(1))
			assert.Equal(t, pro.Sku, "Asj12345")
			assert.Equal(t, pro.Title, "Product1")
			assert.Equal(t, pro.Description, "Desc 1")
			assert.Equal(t, pro.Category, "Cat 1")
			assert.Equal(t, pro.Etalase, "Eta 1")
			assert.Equal(t, pro.Weight, 1.5)

		})
	})
	t.Run("When NewProduct Has error", func(t *testing.T) {
		t.Run("Corner Case 1", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   -5,
				Sku:         "Asj12345",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			_, err := product.NewProduct(param)
			assert.Error(t, err)

		})
		t.Run("Corner Case 2", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj123452",
				Title:       "Product1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			_, err := product.NewProduct(param)
			assert.Error(t, err)

		})
		t.Run("Corner Case 3", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       gen(101),
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			_, err := product.NewProduct(param)
			assert.Error(t, err)

		})
		t.Run("Corner Case 4", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "Title 1",
				Description: gen(1001),
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			_, err := product.NewProduct(param)
			assert.Error(t, err)

		})
		t.Run("Corner Case 5", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "title 1",
				Description: gen(1001),
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			_, err := product.NewProduct(param)
			assert.Error(t, err)

		})
		t.Run("Corner Case 5", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "Title 1",
				Description: "Desc 1",
				Category:    gen(51),
				Etalase:     "Eta 1",
				Weight:      1.5,
			}
			_, err := product.NewProduct(param)
			assert.Error(t, err)

		})
		t.Run("Corner Case 6", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "title 1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     gen(51),
				Weight:      1.5,
			}
			_, err := product.NewProduct(param)
			assert.Error(t, err)

		})
		t.Run("Corner Case 7", func(t *testing.T) {
			param := &product.ProductParam{
				ProductID:   1,
				Sku:         "Asj12345",
				Title:       "title 1",
				Description: "Desc 1",
				Category:    "Cat 1",
				Etalase:     "Eta 1",
				Weight:      0,
			}
			_, err := product.NewProduct(param)
			assert.Error(t, err)

		})

	})
}

func gen(cnt int) string {
	a := ""
	for i := 0; i < cnt; i++ {
		a += "a"
	}
	return a
}

func TestNewImage(t *testing.T) {
	t.Run("When NewImage Has No Error", func(t *testing.T) {
		t.Run("Common Case", func(t *testing.T) {
			param := &product.ImageParam{
				ProductID:   1,
				ImageID:     1,
				Path:        "Path 1",
				Description: "Desc 1",
			}
			im, err := product.NewImage(param)
			assert.NoError(t, err)
			assert.Equal(t, im.ProductID, int64(1))
			assert.Equal(t, im.ImageID, int64(1))
			assert.Equal(t, im.Path, "Path 1")
			assert.Equal(t, im.Description, "Desc 1")
		})

	})
	t.Run("When NewImage Has Error", func(t *testing.T) {
		t.Run("Corner Case", func(t *testing.T) {
			param := &product.ImageParam{
				ProductID:   -1,
				ImageID:     1,
				Path:        "Path 1",
				Description: "Desc 1",
			}
			_, err := product.NewImage(param)
			assert.Error(t, err)
		})
		t.Run("Corner Case", func(t *testing.T) {
			param := &product.ImageParam{
				ProductID:   1,
				ImageID:     -1,
				Path:        "Path 1",
				Description: "Desc 1",
			}
			_, err := product.NewImage(param)
			assert.Error(t, err)
		})
		t.Run("Corner Case", func(t *testing.T) {
			param := &product.ImageParam{
				ProductID:   1,
				ImageID:     1,
				Path:        "",
				Description: "Desc 1",
			}
			_, err := product.NewImage(param)
			assert.Error(t, err)
		})
		t.Run("Corner Case", func(t *testing.T) {
			param := &product.ImageParam{
				ProductID:   1,
				ImageID:     1,
				Path:        "Path 1",
				Description: gen(1001),
			}
			_, err := product.NewImage(param)
			assert.Error(t, err)
		})

	})
}
func TestNewReview(t *testing.T) {
	t.Run("When NewReview Has No Error", func(t *testing.T) {
		t.Run("Common Case", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      1,
				ReviewID:       1,
				Rating:         1,
				ReviewComment:  "Comment 1",
				DateTimeReview: "2023-02-12",
			}
			rev, err := product.NewReview(param)
			assert.NoError(t, err)
			assert.Equal(t, rev.ProductID, int64(1))
			assert.Equal(t, rev.ReviewID, int64(1))
			assert.Equal(t, rev.Rating, float32(1))
			assert.Equal(t, rev.ReviewComment, "Comment 1")
			assert.Equal(t, rev.DateTimeReview, "2023-02-12")

		})
	})
	t.Run("When NewReview Has Error", func(t *testing.T) {
		t.Run("Common Case", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      -1,
				ReviewID:       1,
				Rating:         1,
				ReviewComment:  "Comment 1",
				DateTimeReview: "2023-02-12",
			}
			_, err := product.NewReview(param)
			assert.Error(t, err)

		})
		t.Run("Common Case", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      1,
				ReviewID:       -1,
				Rating:         1,
				ReviewComment:  "Comment 1",
				DateTimeReview: "2023-02-12",
			}
			_, err := product.NewReview(param)
			assert.Error(t, err)

		})
		t.Run("Common Case", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      1,
				ReviewID:       1,
				Rating:         10.1,
				ReviewComment:  "Comment 1",
				DateTimeReview: "2023-02-12",
			}
			_, err := product.NewReview(param)
			assert.Error(t, err)

		})
		t.Run("Common Case", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      1,
				ReviewID:       1,
				Rating:         1,
				ReviewComment:  gen(1001),
				DateTimeReview: "2023-02-12",
			}
			_, err := product.NewReview(param)
			assert.Error(t, err)

		})
		t.Run("Common Case", func(t *testing.T) {
			param := &product.ReviewParam{
				ProductID:      1,
				ReviewID:       1,
				Rating:         1,
				ReviewComment:  "Comment 1",
				DateTimeReview: "2023-02-12sss",
			}
			_, err := product.NewReview(param)
			assert.Error(t, err)

		})
	})
}
