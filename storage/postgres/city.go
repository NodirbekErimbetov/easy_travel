package postgres

import (
	"context"
	"fmt"

	"easy_travel/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CityRepo struct {
	db *pgxpool.Pool
}

func NewCityRepo(db *pgxpool.Pool) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (r *CityRepo) Create(ctx context.Context, req *models.CreateCity) (*models.City, error) {

	var (
		cityId = uuid.New().String()
		query  = `
			INSERT INTO "city"(
				"id",
				"title",
				"country_id",
				"city_code",
				"latitude",
				"longitude",
				"offset",
				"country_name",
				"updated_at"
			) VALUES ($1, $2, $3, $4,$5, $6, $7, $8, NOW())`
	)

	_, err := r.db.Exec(
		ctx,
		query,
		cityId,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.CountryName,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx,&models.CityPrimaryKey{Id: cityId})
}

func (r *CityRepo) GetByID(ctx context.Context,req *models.CityPrimaryKey) (*models.City, error) {

	var (
		city  models.City
		query = `
			SELECT
					"id",
					"title",
					"country_id",
					"city_code",
					"latitude",
					"longitude",
					"offset",
					"country_name",
					"created_at",
					"updated_at"
			FROM "city"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(ctx,query, req.Id).Scan(
		&city.Id,
		&city.Title,
		&city.CountryId,
		&city.CityCode,
		&city.Latitude,
		&city.Longitude,
		&city.Offset,
		&city.CountryName,
		&city.CreatedAt,
		&city.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &city, nil
}

func (r *CityRepo) GetList(ctx context.Context, req *models.GetListCityRequest) (*models.GetListCityResponse, error) {
	var (
		resp   models.GetListCityResponse
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

	if len(req.Search) > 0 {
		where += " AND title ILIKE" + " '%" + req.Search + "%'"
	}

	if len(req.Query) > 0 {
		where += req.Query
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"country_name",
			"created_at",
			"updated_at"
		FROM "city"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(ctx,query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			city models.City
		)

		err = rows.Scan(
			&resp.Count,
			&city.Id,
			&city.Title,
			&city.CountryId,
			&city.CityCode,
			&city.Latitude,
			&city.Longitude,
			&city.Offset,
			&city.CountryName,
			&city.CreatedAt,
			&city.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Cities = append(resp.Cities, &city)
	}

	return &resp, nil
}

func (r *CityRepo) Update(ctx context.Context, req *models.UpdateCity) (int64, error) {

	query := `
		UPDATE city
			SET
				"title" =$2
				"country_id" =$3
				"city_code" =$4
				"latitude" =$5
				"longitude" =$6
				"offset" =$7
				"country_name" =$8
		WHERE id = $1
	`
	result, err := r.db.Exec(
		ctx,
		query,
		req.Id,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.CountryName,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, nil
}

func (r *CityRepo) Delete(ctx context.Context, req *models.CityPrimaryKey) error {
	_, err := r.db.Exec(ctx,"DELETE FROM city WHERE id = $1", req.Id)
	return err
}
