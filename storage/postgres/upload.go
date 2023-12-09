package postgres

import (
	"context"
	"easy_travel/models"
	"easy_travel/pkg/helpers"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/cast"
)

type UploadRepo struct {
	db *pgxpool.Pool
}

func NewUploadRepo(db *pgxpool.Pool) *UploadRepo {
	return &UploadRepo{
		db: db,
	}
}

func (c UploadRepo) Upload(ctx context.Context, req *models.UploadRequest) error {

	if _, ok := models.TableEntity[req.TableSlug]; !ok {
		return errors.New("not found table")
	}

	var (
		entity = models.TableEntity[req.TableSlug]
		query  = "INSERT INTO " + req.TableSlug + "("
		column = strings.Join(entity, ",")
	)
	query += column + ") VALUES"

	if helpers.Contains(entity, `"offset"`) {
		index := helpers.FindIndex(entity, `"offset"`)
		entity[index] = "offset"
	}

	for _, data := range req.Data {
		var obj = cast.ToStringMap(data)
		query += "("
		for ind, key := range entity {
			if _, ok := obj[key]; ok {
				if ind == len(entity)-1 {
					query += fmt.Sprintf(`'%s'`, cast.ToString(obj[key]))
					break
				}

				query += fmt.Sprintf(`'%s',`, cast.ToString(obj[key]))
			} else {
				if ind == len(entity)-1 {
					query += `''`
					break
				}

				query += `'',`
			}
		}
		query += "),"
	}
	query = query[:len(query)-1]
	_, err := c.db.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

