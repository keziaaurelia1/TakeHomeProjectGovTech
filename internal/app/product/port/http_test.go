package port_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	appmock "github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/app_mock"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/domain/product"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/port"
)

type MockObject struct {
	MockApp  *appmock.MockApplication
	MockHTTP *port.HTTP
}

func PrepareTest(mockCtrl *gomock.Controller) MockObject {
	mockApp := appmock.NewMockApplication(mockCtrl)
	mockHTTP := port.ProvideHTTP(mockApp)
	return MockObject{
		MockApp:  mockApp,
		MockHTTP: &mockHTTP,
	}
}

func TestGetIndex(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := PrepareTest(mockCtrl)

	t.Run("When success", func(t *testing.T) {
		responseWriter := httptest.NewRecorder()
		ctx, _ := http.NewRequest("GET", "/index", nil)
		mockObj.MockApp.EXPECT().GetAllProduct(gomock.Any()).Return([]*product.Product{{}}, nil)
		mockObj.MockHTTP.GetIndex(responseWriter, ctx)
		assert.Equal(t, http.StatusOK, responseWriter.Code)
	})
	t.Run("When error", func(t *testing.T) {
		responseWriter := httptest.NewRecorder()
		ctx, _ := http.NewRequest("GET", "/index", nil)
		mockObj.MockApp.EXPECT().GetAllProduct(gomock.Any()).Return(nil, errors.New("error"))
		mockObj.MockHTTP.GetIndex(responseWriter, ctx)
		assert.Equal(t, http.StatusInternalServerError, responseWriter.Code)
	})

}
func TestPostInsertReview(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := PrepareTest(mockCtrl)
	t.Run("post insert success", func(t *testing.T) {
		param := product.ReviewParam{
			ProductID:      1,
			ReviewID:       1,
			Rating:         2,
			ReviewComment:  "Comment 1",
			DateTimeReview: "2023-01-12",
		}
		json, _ := json.Marshal(param)
		reader := strings.NewReader(string(json))
		req, _ := http.NewRequest("POST", "/insert/review", reader)
		mockObj.MockApp.EXPECT().CreateReview(gomock.Any(), gomock.Any()).Return(nil)
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.PostInsertReview(resp, req)

		assert.Equal(t, resp.Code, http.StatusCreated)
	})
	t.Run("post insert error", func(t *testing.T) {
		param := product.ReviewParam{
			ProductID:      1,
			ReviewID:       1,
			Rating:         2,
			ReviewComment:  "Comment 1",
			DateTimeReview: "2023-01-12",
		}
		json, _ := json.Marshal(param)
		reader := strings.NewReader(string(json))
		req, _ := http.NewRequest("POST", "/insert/review", reader)
		mockObj.MockApp.EXPECT().CreateReview(gomock.Any(), gomock.Any()).Return(errors.New("error"))
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.PostInsertReview(resp, req)

		assert.Equal(t, resp.Code, http.StatusInternalServerError)
	})
}
func TestPostInsertProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := PrepareTest(mockCtrl)
	t.Run("post insert success", func(t *testing.T) {
		param := product.ProductParam{
			ProductID:   1,
			Sku:         "abc28373",
			Title:       "Product 1",
			Description: "Desc 1",
			Category:    "Cat 1",
			Etalase:     "Eta 1",
			Weight:      2,
		}
		json, _ := json.Marshal(param)
		reader := strings.NewReader(string(json))
		req, _ := http.NewRequest("POST", "/insert", reader)
		mockObj.MockApp.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(nil)
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.PostInsert(resp, req)

		assert.Equal(t, resp.Code, http.StatusCreated)
	})
	t.Run("post insert error", func(t *testing.T) {
		param := product.ProductParam{
			ProductID:   1,
			Sku:         "abc28373",
			Title:       "Product 1",
			Description: "Desc 1",
			Category:    "Cat 1",
			Etalase:     "Eta 1",
			Weight:      2,
		}
		json, _ := json.Marshal(param)
		reader := strings.NewReader(string(json))
		req, _ := http.NewRequest("POST", "/insert/review", reader)
		mockObj.MockApp.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(errors.New("error"))
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.PostInsert(resp, req)

		assert.Equal(t, resp.Code, http.StatusInternalServerError)
	})
}
func TestPostEdit(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := PrepareTest(mockCtrl)
	t.Run("post edit success", func(t *testing.T) {
		param := product.ProductParam{
			ProductID:   1,
			Sku:         "abc28373",
			Title:       "Product 1",
			Description: "Desc 1",
			Category:    "Cat 1",
			Etalase:     "Eta 1",
			Weight:      2,
		}
		json, _ := json.Marshal(param)
		reader := strings.NewReader(string(json))
		req, _ := http.NewRequest("POST", "/edit/{id}", reader)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "5")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		mockObj.MockApp.EXPECT().UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.PostEdit(resp, req)
		assert.Equal(t, resp.Code, http.StatusOK)
	})
	t.Run("post edit error", func(t *testing.T) {
		param := product.ProductParam{
			ProductID:   1,
			Sku:         "abc28373",
			Title:       "Product 1",
			Description: "Desc 1",
			Category:    "Cat 1",
			Etalase:     "Eta 1",
			Weight:      2,
		}
		json, _ := json.Marshal(param)
		reader := strings.NewReader(string(json))
		req, _ := http.NewRequest("POST", "/edit/{id}", reader)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "5")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		mockObj.MockApp.EXPECT().UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.PostEdit(resp, req)
		assert.Equal(t, resp.Code, http.StatusInternalServerError)
	})
}

func TestShowID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := PrepareTest(mockCtrl)
	t.Run("Get Show ID success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/show/{id}", nil)
		mockObj.MockApp.EXPECT().FindProductByProductID(gomock.Any(), gomock.Any()).Return(&product.Product{}, nil)
		resp := httptest.NewRecorder()

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "5")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		mockObj.MockHTTP.GetShowID(resp, req)
		assert.Equal(t, resp.Code, http.StatusOK)
	})
	t.Run("Get Show ID error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/show/{id}", nil)
		mockObj.MockApp.EXPECT().FindProductByProductID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		resp := httptest.NewRecorder()

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "5")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		mockObj.MockHTTP.GetShowID(resp, req)
		assert.Equal(t, resp.Code, http.StatusInternalServerError)
	})
}

func TestSort(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := PrepareTest(mockCtrl)
	t.Run("get sort success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/sort/{id}", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "5")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		mockObj.MockApp.EXPECT().GetAllReview(gomock.Any(), gomock.Any()).Return([]*product.Review{{}}, nil)
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.GetSort(resp, req)
		assert.Equal(t, resp.Code, http.StatusOK)
	})
	t.Run("get sort error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/sort/{id}", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "5")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		mockObj.MockApp.EXPECT().GetAllReview(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.GetSort(resp, req)
		assert.Equal(t, resp.Code, http.StatusInternalServerError)
	})
}

func TestSearchReview(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := PrepareTest(mockCtrl)
	t.Run("get search review success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/search/{id}/{tanggal}", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "5")
		rctx.URLParams.Add("tanggal", "2023-01-12")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		mockObj.MockApp.EXPECT().FindReviewByDate(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*product.Review{{}}, nil)
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.GetSearchReview(resp, req)
		assert.Equal(t, resp.Code, http.StatusOK)
	})
	t.Run("get sort error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/search/{id}/{tanggal}", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "5")
		rctx.URLParams.Add("tanggal", "2023-01-12")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		mockObj.MockApp.EXPECT().FindReviewByDate(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		resp := httptest.NewRecorder()
		mockObj.MockHTTP.GetSearchReview(resp, req)
		assert.Equal(t, resp.Code, http.StatusInternalServerError)
	})
}

func TestGetSearch(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := PrepareTest(mockCtrl)

	t.Run("When search success", func(t *testing.T) {
		responseWriter := httptest.NewRecorder()
		param := product.SearchParam{
			Title:    "Product 1",
			Category: "Cat 1",
			Etalase:  "Eta 1",
		}
		json, _ := json.Marshal(param)
		reader := strings.NewReader(string(json))
		ctx, _ := http.NewRequest("POST", "/search", reader)
		mockObj.MockApp.EXPECT().FindProductByTitleCategoryEtalase(gomock.Any(), gomock.Any()).Return([]*product.Product{{}}, nil)
		mockObj.MockHTTP.GetSearch(responseWriter, ctx)
		assert.Equal(t, http.StatusOK, responseWriter.Code)
	})
	t.Run("When search error", func(t *testing.T) {
		responseWriter := httptest.NewRecorder()
		param := product.SearchParam{
			Title:    "Product 1",
			Category: "Cat 1",
			Etalase:  "Eta 1",
		}
		json, _ := json.Marshal(param)
		reader := strings.NewReader(string(json))
		ctx, _ := http.NewRequest("POST", "/search", reader)
		mockObj.MockApp.EXPECT().FindProductByTitleCategoryEtalase(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		mockObj.MockHTTP.GetSearch(responseWriter, ctx)
		assert.Equal(t, http.StatusInternalServerError, responseWriter.Code)
	})

}
