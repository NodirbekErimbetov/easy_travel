package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"easy_travel/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CountryRepo struct {
	db *pgxpool.Pool
}

func NewCountryRepo(db *pgxpool.Pool) *CountryRepo {
	return &CountryRepo{
		db: db,
	}
}

func (r *CountryRepo) Create(ctx context.Context, req *models.CreateCountry) (*models.Country, error) {

	var (
		countryId = uuid.New().String()
		query     = `
			INSERT INTO "country"(
				"id",
				"title",
				"code",
				"continent",
				"updated_at"
			) VALUES ($1, $2, $3, $4, NOW())`
	)

	_, err := r.db.Exec(
		ctx,
		query,
		countryId,
		req.Title,
		req.Code,
		req.Continent,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, &models.CountryPrimaryKey{Id: countryId})
}

func (r *CountryRepo) GetByID(ctx context.Context, req *models.CountryPrimaryKey) (*models.Country, error) {

	var (
		country models.Country
		query   = `
			SELECT
				"id",
				"title",
				"code",
				"continent",
				"created_at",
				"updated_at"	
			FROM "country"
			WHERE "id" = $1
		`
	)
	var (
		id         sql.NullString
		title      sql.NullString
		code       sql.NullString
		continent  sql.NullString
		updated_at sql.NullString
		created_at sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&title,
		&code,
		&continent,
		&updated_at,
		&created_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Country{
		Id: id.String,
		Title: title.String,
		Code: code.String,
		Continent: continent.String,
		UpdatedAt: updated_at.String,
		CreatedAt: country.CreatedAt,
	}, nil
}

func (r *CountryRepo) GetList(ctx context.Context, req *models.GetListCountryRequest) (*models.GetListCounrtyResponse, error) {
	var (
		resp  = &models.GetListCounrtyResponse{}
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if req.FromDate != "" || req.Todate != "" {
		where += " AND created_at BETWEEN $1 and $2 "
	}


	if len(req.Search) > 0 {
		where += " AND title ILIKE" + " '%" + req.Search + "%'"
	}

	if len(req.Query) > 0 {
		where += req.Query
	}

	var query = `
		SELECT
			"id",
			"title",
			"code",
			"continent",
			"created_at",
			"updated_at"
		FROM "country"
		
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var (
		id         sql.NullString
		title      sql.NullString
		code       sql.NullString
		continent  sql.NullString
		updated_at sql.NullString
		created_at sql.NullString
	)
	for rows.Next() {

		err = rows.Scan(
			&id,
			&title,
			&code,
			&continent,
			&updated_at,
			&created_at,			
			
		)
		if err != nil {
			return nil, err
		}

		resp.Countries = append(resp.Countries, &models.Country{
			Id: id.String,
			Title: title.String,
            Code: code.String,
            Continent: continent.String,
            UpdatedAt: updated_at.String,
            CreatedAt: created_at.String,
		})
	}

	return resp, nil
}

func (r *CountryRepo) UpdateCountry(ctx context.Context, req *models.UpdateCountry) (int64, error) {

	query := `
		UPDATE country
			SET
				title = $2,
				code = $3,
				continent = $4
		WHERE id = $1
	`
	result, err := r.db.Exec(
		ctx,
		query,
		req.Id,
		req.Title,
		req.Code,
		req.Continent,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *CountryRepo) Delete(ctx context.Context, req *models.CountryPrimaryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM country WHERE id = $1", req.Id)
	return err
}

