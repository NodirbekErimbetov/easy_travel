package postgres

import (
	"context"
	"easy_travel/models"
	"fmt"
	"testing"
)

func TestCreateCountry(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateCountry
		Output  *models.Country
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateCountry{
				Title:     "JP",
				Code:      "JP",
				Continent: "AI",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			data, err := countryTestRepo.Create(ctx, test.Input)
			if test.WantErr || err != nil {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}
			fmt.Println(">>>", data)
			if data.Code != test.Input.Code || data.Title != test.Input.Title {
				t.Errorf("%s: no mistmatch data got %v", test.Name, data)
				return
			}

		})
	}
}
func TestUpdateCountry(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateCountry
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateCountry{
				Id:        "b2cb6673-fd54-4eda-9724-103e6587a7c5",
				Title:     "Korea Kr",
				Code:      "KR",
				Continent: "KR",
			},
		},
		{
			Name: "Case 2",
			Input: &models.UpdateCountry{
				Id:        "1a9760d3-14d4-433a-9e12-37e73cc4db63",
				Title:     "Golland",
				Code:      "Gl",
				Continent: "Gl",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			_, err := countryTestRepo.UpdateCountry(ctx, test.Input)
			if test.WantErr || err != nil {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}
			if err != nil {
				t.Errorf("%s: unexpected error during update: %v", test.Name, err)
				return
			}
			_, err = countryTestRepo.GetByID(ctx, &models.CountryPrimaryKey{Id: test.Input.Id})
			if err != nil {
				t.Errorf("%s: error getting updated country: %v", test.Name, err)
				return
			}

		})
	}
}
func TestGetListCountry(t *testing.T) {
	tests := []struct {
		Name     string
		Input    *models.GetListCountryRequest
		Expected int
		WantErr  bool
	}{
		{
			Name: "Test with valid parameters",
			Input: &models.GetListCountryRequest{
				FromDate: "2023-12-04",
				Todate: "2023-12-05",
			},
			Expected: 2,
			WantErr:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			resp, err := countryTestRepo.GetList(ctx, test.Input)
			if test.WantErr && err == nil {
				t.Errorf("%s: expected an error, but got nil", test.Name)
				return
			}

			if !test.WantErr && err != nil {
				t.Errorf("%s: unexpected error: %v", test.Name, err)
				return
			}

			if len(resp.Countries) != test.Expected {
				t.Errorf("%s: expected %d countries, but got %d", test.Name, test.Expected, len(resp.Countries))
			}
		})
	}
}
func TestGetById(t *testing.T) {

	tests := []struct {
		Name     string
		Input    string
		Expected *models.Country
		WantErr  bool
	}{
		{
			Name:  "Case 1",
			Input: "b2cb6673-fd54-4eda-9724-103e6587a7c5",
			Expected: &models.Country{
				Id:        "b2cb6673-fd54-4eda-9724-103e6587a7c5",
				Title:     "Korea Kr",
				Code:      "KR",
				Continent: "KR",
				CreatedAt: "",
				UpdatedAt: "2023-12-05T19:27:04.191428Z",
			},
			WantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			resp, err := countryTestRepo.GetByID(ctx, &models.CountryPrimaryKey{Id: test.Input})
			if test.WantErr && err == nil {
				t.Errorf("%s: expected an error, but got nil", test.Name)
				return
			}

			if !test.WantErr && err != nil {
				t.Errorf("%s: unexpected error: %v", test.Name, err)
				return
			}
			if !test.WantErr && resp != nil {
				if resp.Id != test.Expected.Id || resp.Title != test.Expected.Title || resp.Code != test.Expected.Code || resp.Continent != test.Expected.Continent || resp.CreatedAt != test.Expected.CreatedAt || resp.UpdatedAt != test.Expected.UpdatedAt {
					t.Errorf("%s: expected %+v, but got %+v", test.Name, test.Expected, resp)
				}
			}

		})
	}

}

func TestDelete(t *testing.T) {
	tests := []struct {
		Name     string
		Input    string
		Expected int
		WantErr  bool
	}{
		{
			Name:     "Case 1",
			Input:    "6a722107-17ed-451a-9381-8a0d6816fe08",
			Expected: 1,
			WantErr:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			err := countryTestRepo.Delete(ctx, &models.CountryPrimaryKey{Id: test.Input})
			if test.WantErr && err == nil {
				t.Errorf("%s: expected an error, but got nil", test.Name)
				return
			}
		})
	}
}
