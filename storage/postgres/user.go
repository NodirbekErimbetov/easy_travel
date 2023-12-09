package postgres

import (
	"database/sql"
	"easy_travel/config"
	"easy_travel/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(ctx context.Context, req *models.CreateUser) (*models.User, error) {

	var (
		userId    = uuid.New().String()
		expiredAt = time.Now().Add(config.ApiKeyExpiredAt)
		query     = `
		INSERT INTO "user"(
			"id",
			"name",
			"login",
			"password",
			"expired_at",
			"updated_at"
		) VALUES($1,$2,$3,$4,$5,NOW()) `
	)

	_, err := u.db.Exec(
		ctx,
		query,
		userId,
		req.Name,
		req.Login,
		req.Password,
		expiredAt,
	)
	if err != nil {
		return nil, err
	}
	return u.GetByID(ctx, &models.UserPrimaryKey{Id: userId})

}

func (u *UserRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query = `
		SELECT 
			"id",
			"name",
			"login",
			"password",
			"created_at",
			"expired_at",
			"updated_at"
		FROM "user"
		WHERE "id" = $1`
	)
	var (
		id         sql.NullString
		name       sql.NullString
		login      sql.NullString
		password   sql.NullString
		created_at sql.NullString
		expired_at sql.NullTime
		updated_at sql.NullString
	)

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&login,
		&password,
		&created_at,
		&expired_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        id.String,
		Name:      name.String,
		Login:     login.String,
		Password:  password.String,
		CreatedAt: created_at.String,
		ExpiredAt: expired_at.Time,
		UpdatedAt: updated_at.String,
	}, nil
}

func (u *UserRepo) LoginUser(ctx context.Context, req *models.LoginUser) (*models.User, error) {

	var (
		query = `SELECT 
						"id",
						"name",
						"login",
						"password",
						"created_at",
						"expired_at",
						"updated_at"
				FROM "user"
				WHERE login = $1 OR password = $2
	`
	)
	var (
		id         sql.NullString
		name       sql.NullString
		login      sql.NullString
		password   sql.NullString
		created_at sql.NullString
		expired_at sql.NullTime
		updated_at sql.NullString
	)

	err := u.db.QueryRow(ctx, query, req.Login,req.Password).Scan(
		&id,
		&name,
		&login,
		&password,
		&created_at,
		&expired_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Id:        id.String,
		Name:      name.String,
		Login:     login.String,
		Password:  password.String,
		CreatedAt: created_at.String,
		ExpiredAt: expired_at.Time,
		UpdatedAt: updated_at.String,
	}, nil

}


func (u *UserRepo) UpdateExpireTime(ctx context.Context, req *models.UserPrimaryKey)(int64,error) {

	expireTime := time.Now().Add(config.ApiKeyExpiredAt)
	update := time.Now()
	query := `Update "user" SET expired_at = $1, updated_at = $3 WHERE id = $2`

	result,err := u.db.Exec(
		ctx,
		query,
		expireTime,
		req.Id,
		update,

    )
	if err!= nil {
        return 0,err
    }
	rowsAffected := result.RowsAffected()
	if err!= nil {
        return 0,err
    }
	return rowsAffected,nil

}