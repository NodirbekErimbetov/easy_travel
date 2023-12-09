package postgres

import (
	"easy_travel/config"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
)

var (
	countryTestRepo  *CountryRepo
	cityTestRepo     *CityRepo
	aeroportTestRepo *AeroportRepo
)

// func TestMain(m *testing.M) {
// 	cfg := config.Load()
  
// 	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresPort)
  
// 	poolConfig, err := pgxpool.ParseConfig(connStr)
// 	if err != nil {
// 	  panic(fmt.Sprintf("Error parsing connection string: %v", err))
// 	}
  
// 	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
// 	if err != nil {
// 	  panic(fmt.Sprintf("Error connecting to the database: %v", err))
// 	}
  
// 	countryTestRepo = NewCountryRepo(pool)
// 	cityTestRepo = NewCityRepo(pool)
// 	aeroportTestRepo = NewAeroportRepo(pool)
  
// 	os.Exit(m.Run())
//   }

func TestMain( m *testing.M) {
	cfg := config.Load()

	config, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s",
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		fmt.Println("Error while TestMain")
		return
	}
	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil{
		fmt.Println("panic")
		return
	}

	countryTestRepo = NewCountryRepo(pgxpool)
	cityTestRepo = NewCityRepo(pgxpool)
	aeroportTestRepo = NewAeroportRepo(pgxpool)

	os.Exit(m.Run())
}
