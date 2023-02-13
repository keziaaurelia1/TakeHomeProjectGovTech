package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/domain/product"
)

type MySQLRepository struct {
	db sq.StdSqlCtx
}

var (
	ProductTableName = "product"
	ProductTable     = struct {
		ProductID   string
		Sku         string
		Title       string
		Description string
		Category    string
		Etalase     string
		Weight      string
		Price       string
		UpdatedAt   string
		CreatedAt   string
	}{
		ProductID:   "product_id",
		Sku:         "sku",
		Title:       "title",
		Description: "description",
		Category:    "category",
		Etalase:     "etalase",
		Weight:      "weight",
		Price:       "price",
		UpdatedAt:   "updated_at",
		CreatedAt:   "created_at",
	}

	ImageTableName = "image"
	ImageTable     = struct {
		ProductID   string
		ImageID     string
		Path        string
		Description string
		CreatedAt   string
	}{
		ProductID:   "product_id",
		ImageID:     "image_id",
		Path:        "path",
		Description: "description",
		CreatedAt:   "created_at",
	}

	ReviewTableName = "review"
	ReviewTable     = struct {
		ProductID      string
		ReviewID       string
		Rating         string
		ReviewComment  string
		DateTimeReview string
		CreatedAt      string
	}{
		ProductID:      "product_id",
		ReviewID:       "review_id",
		Rating:         "rating",
		ReviewComment:  "review_comment",
		DateTimeReview: "date_time_review",
		CreatedAt:      "created_at",
	}

	ErrProductNotFound = errors.New("Product tidak ditemukan")
	ErrReviewNotFound  = errors.New("Review tidak ditemukan")

	MsgProductNotFound = "Product yang dicari tidak dapat ditemukan..."
	MsgReviewNotFound  = "Review yang dicari tidak dapat ditemukan..."
)

func ProvideMySQLRepository(db *sql.DB) product.Repository {
	return &MySQLRepository{
		db: db,
	}
}

func (repo MySQLRepository) AllReview(ctx context.Context, productID int64) ([]*product.Review, error) {
	dbProduct, err := repo.orderByDate(ctx, productID)
	return dbProduct, err
}
func (repo MySQLRepository) AllProduct(ctx context.Context) ([]*product.Product, error) {
	dbProduct, err := repo.getAll(ctx)
	return dbProduct, err
}

func (repo MySQLRepository) orderByDate(ctx context.Context, productID int64) (list []*product.Review, err error) {
	builder := sq.Select(
		ReviewTable.ProductID,
		ReviewTable.ReviewID,
		ReviewTable.Rating,
		ReviewTable.ReviewComment,
		ReviewTable.DateTimeReview,
	).From(ReviewTableName).OrderBy(ReviewTable.DateTimeReview + " DESC")

	builder = builder.Where(sq.Eq{ReviewTable.ProductID: productID})

	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}

	rows, err := repo.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return
	}
	list = make([]*product.Review, 0)
	for rows.Next() {
		ent := new(product.Review)
		if err = rows.Scan(
			&ent.ProductID,
			&ent.ReviewID,
			&ent.Rating,
			&ent.ReviewComment,
			&ent.DateTimeReview,
		); err != nil {
			return
		}
		list = append(list, ent)
	}
	return
}

func (repo MySQLRepository) FindBySku(ctx context.Context, sku string) (*product.Product, error) {
	dbProduct, err := repo.findBySku(ctx, sku)
	if err != nil {
		return nil, err
	}
	if len(dbProduct) == 0 {
		return nil, ErrProductNotFound
	}
	return dbProduct[0], nil
}
func (repo MySQLRepository) FindByProductID(ctx context.Context, productID int64) (*product.Product, error) {
	dbProduct, err := repo.findByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	if len(dbProduct) == 0 {
		return nil, ErrProductNotFound
	}
	return dbProduct[0], nil
}

func (repo MySQLRepository) FindByReviewID(ctx context.Context, reviewID int64) (*product.Review, error) {
	dbReview, err := repo.findByReviewID(ctx, reviewID)
	if err != nil {
		return nil, err
	}
	if len(dbReview) == 0 {
		return nil, ErrReviewNotFound
	}
	return dbReview[0], nil
}

func (repo MySQLRepository) FindByTitleCategoryEtalase(ctx context.Context, search *product.SearchParam) ([]*product.Product, error) {
	dbProduct, err := repo.findByTitleCategoryEtalase(ctx, search)
	if err != nil {
		return nil, err
	}
	if len(dbProduct) == 0 {
		return nil, ErrProductNotFound
	}
	return dbProduct, nil
}
func (repo MySQLRepository) FindByDate(ctx context.Context, tanggal string, productID int64) ([]*product.Review, error) {
	dbReview, err := repo.findByDate(ctx, tanggal, productID)
	if err != nil {
		return nil, err
	}
	if len(dbReview) == 0 {
		return nil, ErrReviewNotFound
	}
	return dbReview, nil
}

func (repo MySQLRepository) findByDate(ctx context.Context, tanggal string, productID int64) (list []*product.Review, err error) {
	builder := sq.Select(
		ReviewTable.ProductID,
		ReviewTable.ReviewID,
		ReviewTable.Rating,
		ReviewTable.ReviewComment,
		ReviewTable.DateTimeReview,
	).From(ReviewTableName)
	builder = builder.Where(sq.Eq{ReviewTable.ProductID: productID})
	builder = builder.Where(sq.Eq{ReviewTable.DateTimeReview: tanggal})
	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}
	rows, err := repo.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return
	}
	list = make([]*product.Review, 0)
	for rows.Next() {
		ent := new(product.Review)
		if err = rows.Scan(
			&ent.ProductID,
			&ent.ReviewID,
			&ent.Rating,
			&ent.ReviewComment,
			&ent.DateTimeReview,
		); err != nil {
			return
		}
		list = append(list, ent)
	}
	return
}

func (repo MySQLRepository) findByTitleCategoryEtalase(ctx context.Context, search *product.SearchParam) (list []*product.Product, err error) {
	builder := sq.Select(
		ProductTable.ProductID,
		ProductTable.Sku,
		ProductTable.Title,
		ProductTable.Description,
		ProductTable.Category,
		ProductTable.Etalase,
		ProductTable.Weight,
		ProductTable.Price,
	).From(ProductTableName).OrderBy(ProductTable.Sku)
	if search.Title != "" {
		builder = builder.Where(sq.Eq{ProductTable.Title: search.Title})
	}
	if search.Category != "" {
		builder = builder.Where(sq.Eq{ProductTable.Category: search.Category})
	}
	if search.Etalase != "" {
		builder = builder.Where(sq.Eq{ProductTable.Etalase: search.Etalase})
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}
	rows, err := repo.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return
	}
	list = make([]*product.Product, 0)
	for rows.Next() {
		ent := new(product.Product)
		if err = rows.Scan(
			&ent.ProductID,
			&ent.Sku,
			&ent.Title,
			&ent.Description,
			&ent.Category,
			&ent.Etalase,
			&ent.Weight,
			&ent.Price,
		); err != nil {
			return
		}
		list = append(list, ent)
	}
	return
}

func (repo MySQLRepository) getAll(ctx context.Context) (list []*product.Product, err error) {
	builder := sq.Select(
		ProductTable.ProductID,
		ProductTable.Sku,
		ProductTable.Title,
		ProductTable.Description,
		ProductTable.Category,
		ProductTable.Etalase,
		ProductTable.Weight,
		ProductTable.Price,
	).From(ProductTableName).OrderBy(ProductTable.Sku)

	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}
	rows, err := repo.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return
	}
	list = make([]*product.Product, 0)
	for rows.Next() {
		ent := new(product.Product)
		if err = rows.Scan(
			&ent.ProductID,
			&ent.Sku,
			&ent.Title,
			&ent.Description,
			&ent.Category,
			&ent.Etalase,
			&ent.Weight,
			&ent.Price,
		); err != nil {
			return
		}
		list = append(list, ent)
	}
	return
}

func (repo MySQLRepository) InsertProduct(ctx context.Context, entProduct *product.Product) error {
	return repo.createProduct(ctx, entProduct)
}
func (repo MySQLRepository) InsertReview(ctx context.Context, entReview *product.Review) error {
	return repo.createReview(ctx, entReview)
}

func (repo MySQLRepository) UpdateProduct(ctx context.Context, prevProductID int64, entProduct *product.Product) error {
	_, err := repo.updateProduct(ctx, prevProductID, entProduct)
	return err
}

func (repo MySQLRepository) findBySku(ctx context.Context, sku string) (list []*product.Product, err error) {
	builder := sq.Select(
		ProductTable.ProductID,
		ProductTable.Sku,
		ProductTable.Title,
		ProductTable.Description,
		ProductTable.Category,
		ProductTable.Etalase,
		ProductTable.Weight,
		ProductTable.Price,
	).From(ProductTableName)
	if sku != "" {
		builder = builder.Where(sq.Eq{ProductTable.Sku: sku})
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}
	rows, err := repo.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return
	}
	list = make([]*product.Product, 0)
	for rows.Next() {
		ent := new(product.Product)
		if err = rows.Scan(
			&ent.ProductID,
			&ent.Sku,
			&ent.Title,
			&ent.Description,
			&ent.Category,
			&ent.Etalase,
			&ent.Weight,
			&ent.Price,
		); err != nil {
			return
		}
		list = append(list, ent)
	}
	return
}
func (repo MySQLRepository) findByProductID(ctx context.Context, productID int64) (list []*product.Product, err error) {
	builder := sq.Select(
		ProductTable.ProductID,
		ProductTable.Sku,
		ProductTable.Title,
		ProductTable.Description,
		ProductTable.Category,
		ProductTable.Etalase,
		ProductTable.Weight,
		ProductTable.Price,
	).From(ProductTableName)

	builder = builder.Where(sq.Eq{ProductTable.ProductID: productID})
	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}
	rows, err := repo.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return
	}
	list = make([]*product.Product, 0)
	for rows.Next() {
		ent := new(product.Product)
		if err = rows.Scan(
			&ent.ProductID,
			&ent.Sku,
			&ent.Title,
			&ent.Description,
			&ent.Category,
			&ent.Etalase,
			&ent.Weight,
			&ent.Price,
		); err != nil {
			return
		}
		list = append(list, ent)
	}
	return
}

func (repo MySQLRepository) findByReviewID(ctx context.Context, reviewID int64) (list []*product.Review, err error) {
	builder := sq.Select(
		ReviewTable.ProductID,
		ReviewTable.ReviewID,
		ReviewTable.Rating,
		ReviewTable.ReviewComment,
		ReviewTable.DateTimeReview,
	).From(ReviewTableName)

	builder = builder.Where(sq.Eq{ReviewTable.ReviewID: reviewID})
	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}
	rows, err := repo.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return
	}
	list = make([]*product.Review, 0)
	for rows.Next() {
		ent := new(product.Review)
		if err = rows.Scan(
			&ent.ProductID,
			&ent.ReviewID,
			&ent.Rating,
			&ent.ReviewComment,
			&ent.DateTimeReview,
		); err != nil {
			return
		}
		list = append(list, ent)
	}
	return
}

func (repo MySQLRepository) createProduct(ctx context.Context, ent *product.Product) (err error) {
	now := time.Now().UTC()
	query, args, err := sq.Insert(ProductTableName).
		Columns(ProductTable.ProductID,
			ProductTable.Sku,
			ProductTable.Title,
			ProductTable.Description,
			ProductTable.Category,
			ProductTable.Etalase,
			ProductTable.Weight,
			ProductTable.Price,
			ProductTable.UpdatedAt,
			ProductTable.CreatedAt).
		Values(ent.ProductID,
			ent.Sku,
			ent.Title,
			ent.Description,
			ent.Category,
			ent.Etalase,
			ent.Weight,
			ent.Price,
			now,
			now).ToSql()
	if err != nil {
		return
	}
	_, err = repo.db.ExecContext(ctx, query, args...)
	return
}

func (repo MySQLRepository) createReview(ctx context.Context, ent *product.Review) (err error) {
	now := time.Now().UTC()
	query, args, err := sq.Insert(ReviewTableName).
		Columns(ReviewTable.ProductID,
			ReviewTable.ReviewID,
			ReviewTable.Rating,
			ReviewTable.ReviewComment,
			ReviewTable.DateTimeReview,
			ProductTable.CreatedAt).
		Values(ent.ProductID,
			ent.ReviewID,
			ent.Rating,
			ent.ReviewComment,
			ent.DateTimeReview,
			now).ToSql()
	if err != nil {
		return
	}
	_, err = repo.db.ExecContext(ctx, query, args...)
	return
}

func (repo MySQLRepository) updateProduct(ctx context.Context, prevProjectID int64, ent *product.Product) (affectedRow int64, err error) {
	builder := sq.Update(ProductTableName)
	builder = builder.Set(ProductTable.ProductID, ent.ProductID)
	builder = builder.Set(ProductTable.Sku, ent.Sku)
	builder = builder.Set(ProductTable.Title, ent.Title)
	builder = builder.Set(ProductTable.Description, ent.Description)
	builder = builder.Set(ProductTable.Category, ent.Category)
	builder = builder.Set(ProductTable.Etalase, ent.Etalase)
	builder = builder.Set(ProductTable.Weight, ent.Weight)
	builder = builder.Set(ProductTable.Price, ent.Price)
	builder = builder.Set(ProductTable.UpdatedAt, time.Now().UTC())
	builder = builder.Where(sq.Eq{ProductTable.ProductID: prevProjectID})

	query, args, err := builder.ToSql()
	if err != nil {
		return
	}
	res, err := repo.db.ExecContext(ctx, query, args...)
	if err != nil {
		return
	}
	affectedRow, err = res.RowsAffected()
	return

}
