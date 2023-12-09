package storage

import (
	"context"
	"easy_travel/models"
)

type StorageI interface {
	Country() CountryRepoI
	City() CityRepoI
	Aeroport() AeroportRepoI
	Upload() UploadRepoI
	User() UserRepoI
}

type CountryRepoI interface {
	Create(ctx context.Context, req *models.CreateCountry) (*models.Country, error)
	GetByID(ctx context.Context, req *models.CountryPrimaryKey) (*models.Country, error)
	GetList(ctx context.Context, req *models.GetListCountryRequest) (*models.GetListCounrtyResponse, error)
	UpdateCountry(ctx context.Context, req *models.UpdateCountry) (int64, error)
	Delete(ctx context.Context, req *models.CountryPrimaryKey) error
	// UploadCity(ctx context.Context,  req *models.Country) (*models.GetListCounrtyResponse, error)
}

type CityRepoI interface {
	Create(ctx context.Context, req *models.CreateCity) (*models.City, error)
	GetByID(ctx context.Context, req *models.CityPrimaryKey) (*models.City, error)
	GetList(ctx context.Context, req *models.GetListCityRequest) (*models.GetListCityResponse, error)
	Update(ctx context.Context, req *models.UpdateCity) (int64, error)
	Delete(ctx context.Context, req *models.CityPrimaryKey) error
}

type AeroportRepoI interface {
	Create(ctx context.Context, req *models.CreateAeroport) (*models.Aeroport, error)
	GetByID(ctx context.Context, req *models.AeroportPrimaryKey) (*models.Aeroport, error)
	GetList(ctx context.Context, req *models.GetListAeroportRequest) (*models.GetListAeroportResponse, error)
	Update(ctx context.Context, req *models.UpdateAeroport) (int64, error)
	Delete(ctx context.Context, req *models.AeroportPrimaryKey) error
}

type UploadRepoI interface {
	Upload(ctx context.Context, req *models.UploadRequest) error
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (*models.User, error)
	GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	LoginUser(ctx context.Context, req *models.LoginUser) (*models.User, error)
	UpdateExpireTime(ctx context.Context, req *models.UserPrimaryKey) (int64, error)
}
