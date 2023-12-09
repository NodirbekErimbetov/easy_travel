package postgres

import (
	"context"
	"fmt"

	"easy_travel/config"
	"easy_travel/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	country  *CountryRepo
	city     *CityRepo
	aeroport *AeroportRepo
	upload   *UploadRepo
	user 	 *UserRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(

		fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s port=%s ",
			cfg.PostgresHost,
			cfg.PostgresUser,
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
			cfg.PostgresPort,
		),
	)
	if err != nil {
		return nil, err
	}
	config.MaxConns = cfg.PostgresMaxConnection
	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)

	return &Store{
		db: pgxpool,
	}, nil
}

func (s *Store) Country() storage.CountryRepoI {

	if s.country == nil {
		s.country = NewCountryRepo(s.db)
	}

	return s.country
}

func (s *Store) City() storage.CityRepoI {

	if s.city == nil {
		s.city = NewCityRepo(s.db)
	}

	return s.city
}

func (s *Store) Aeroport() storage.AeroportRepoI {

	if s.aeroport == nil {
		s.aeroport = NewAeroportRepo(s.db)
	}

	return s.aeroport
}

func (s *Store) Upload() storage.UploadRepoI {

	if s.upload == nil {
		s.upload = NewUploadRepo(s.db)
	}

	return s.upload
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}
