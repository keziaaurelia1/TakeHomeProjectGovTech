package port

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/app"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/domain/product"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/repo"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/common/errorutility"
)

type (
	HTTP struct {
		App app.Application
	}
)

func init() {
	errorutility.SetHTTPMapping(ErrMissingParam, errorutility.NewHTTPMapping(http.StatusBadRequest, MsgMissingParam))
	errorutility.SetHTTPMapping(repo.ErrProductNotFound, errorutility.NewHTTPMapping(http.StatusBadRequest, repo.MsgProductNotFound))
	errorutility.SetHTTPMapping(repo.ErrReviewNotFound, errorutility.NewHTTPMapping(http.StatusBadRequest, repo.MsgReviewNotFound))
	errorutility.SetHTTPMapping(app.ErrProductIDExist, errorutility.NewHTTPMapping(http.StatusBadRequest, app.MsgErrProductIDExist))
	errorutility.SetHTTPMapping(app.ErrReviewIDExist, errorutility.NewHTTPMapping(http.StatusBadRequest, app.MsgErrReviewIDExist))
	errorutility.SetHTTPMapping(app.ErrUpdatedProductIDExist, errorutility.NewHTTPMapping(http.StatusBadRequest, app.MsgErrUpdatedProductIDExist))
	errorutility.SetHTTPMapping(product.ErrInvalidSku, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrInvalidSku))
	errorutility.SetHTTPMapping(product.ErrTitleTooLong, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrTitleTooLong))
	errorutility.SetHTTPMapping(product.ErrDescriptionTooLong, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrDescriptionTooLong))
	errorutility.SetHTTPMapping(product.ErrCategoryTooLong, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrCategoryTooLong))
	errorutility.SetHTTPMapping(product.ErrEtalaseTooLong, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrEtalaseTooLong))
	errorutility.SetHTTPMapping(product.ErrWeightNotPositive, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrWeightNotPositive))
	errorutility.SetHTTPMapping(product.ErrInvalidImageID, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrInvalidImageID))
	errorutility.SetHTTPMapping(product.ErrInvalidPath, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrInvalidPath))
	errorutility.SetHTTPMapping(product.ErrInvalidReviewID, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrInvalidReviewID))
	errorutility.SetHTTPMapping(product.ErrInvalidRating, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrInvalidRating))
	errorutility.SetHTTPMapping(product.ErrReviewCommentTooLong, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrReviewCommentTooLong))
	errorutility.SetHTTPMapping(product.ErrInvalidProductID, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrInvalidProductID))
	errorutility.SetHTTPMapping(product.ErrInvalidDateTime, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrInvalidDateTime))
	errorutility.SetHTTPMapping(product.ErrInvalidPrice, errorutility.NewHTTPMapping(http.StatusBadRequest, product.MsgErrInvalidPrice))

}

func ProvideHTTP(app app.Application) HTTP {
	return HTTP{App: app}
}
func (h HTTP) RegisterRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/index", h.GetIndex)
	r.Post("/insert", h.PostInsert)
	r.Post("/edit/{id}", h.PostEdit)
	r.Get("/show/{id}", h.GetShowID)
	r.Post("/search", h.GetSearch)
	r.Route("/review", func(r chi.Router) {
		r.Post("/insert", h.PostInsertReview)
		r.Get("/search/{id}/{tanggal}", h.GetSearchReview)
		r.Get("/sort/{id}", h.GetSort)
	})
	return r
}

func (h HTTP) GetIndex(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	listProduct, err := h.App.GetAllProduct(c)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(listProduct)
	if _, err := io.Copy(w, &buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h HTTP) PostInsert(w http.ResponseWriter, r *http.Request) {
	param := &product.ProductParam{}
	c := r.Context()
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	err = json.Unmarshal(jsonData, &param)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	err = h.App.CreateProduct(c, param)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h HTTP) PostEdit(w http.ResponseWriter, r *http.Request) {
	prevID, err := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	param := &product.ProductParam{}
	c := r.Context()
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	err = json.Unmarshal(jsonData, &param)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	err = h.App.UpdateProduct(c, prevID, param)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h HTTP) GetShowID(w http.ResponseWriter, r *http.Request) {
	prevID, err := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)

	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	c := r.Context()
	tmpProduct, err := h.App.FindProductByProductID(c, prevID)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(tmpProduct)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	if _, err := io.Copy(w, &buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h HTTP) GetSort(w http.ResponseWriter, r *http.Request) {
	prevID, err := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	c := r.Context()
	tmpProduct, err := h.App.GetAllReview(c, prevID)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(tmpProduct)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	if _, err := io.Copy(w, &buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h HTTP) GetSearch(w http.ResponseWriter, r *http.Request) {
	param := &product.SearchParam{}
	c := r.Context()
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	err = json.Unmarshal(jsonData, &param)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	listProduct, err := h.App.FindProductByTitleCategoryEtalase(c, param)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(listProduct)
	if _, err := io.Copy(w, &buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h HTTP) PostInsertReview(w http.ResponseWriter, r *http.Request) {
	param := &product.ReviewParam{}
	c := r.Context()
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	err = json.Unmarshal(jsonData, &param)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	err = h.App.CreateReview(c, param)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h HTTP) GetSearchReview(w http.ResponseWriter, r *http.Request) {
	prevID, err := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)
	tanggal := chi.URLParam(r, "tanggal")
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	c := r.Context()
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	listReview, err := h.App.FindReviewByDate(c, tanggal, prevID)
	if err != nil {
		mapping := errorutility.GetHTTPMapping(err)
		http.Error(w, err.Error(), mapping.StatusCode)
		return
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(listReview)
	if _, err := io.Copy(w, &buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
