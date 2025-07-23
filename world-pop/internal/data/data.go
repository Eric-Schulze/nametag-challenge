package data

import (
	"embed"
	"fmt"
	"io"

	"world-pop/internal/common/logger"

	"github.com/jszwec/csvutil"
)

type CountryDataManager struct {
	PathToData  string
	CountryData []Country
	Logger      logger.Logger
}

type Country struct {
	Name                      string  `csv:"Country/Territory"`
	Code                      string  `csv:"CCA3"`
	PopulationRank            string  `csv:"Rank"`
	Capital                   string  `csv:"Capital"`
	Continent                 string  `csv:"Continent"`
	Population2022            int     `csv:"2022 Population"`
	Population2020            int     `csv:"2020 Population"`
	Population2015            int     `csv:"2015 Population"`
	Population2010            int     `csv:"2010 Population"`
	Population2000            int     `csv:"2000 Population"`
	Population1990            int     `csv:"1990 Population"`
	Population1980            int     `csv:"1980 Population"`
	Population1970            int     `csv:"1970 Population"`
	Area                      int     `csv:"Area (km²)"`
	PopulationDensity         float32 `csv:"Density (per km²)"`
	GrowthRate                float32 `csv:"Growth Rate"`
	WorldPopulationPercentage float32 `csv:"World Population Percentage"`
}

//go:embed world_population.csv
var f embed.FS

func NewCountryDataManager(pathToData string, log logger.Logger) (*CountryDataManager, error) {
	var data []Country

	fileBytes, err := f.ReadFile("world_population.csv")
	if err != nil {
		return &CountryDataManager{}, err
	}

	if err := csvutil.Unmarshal(fileBytes, &data); err != nil {
		return &CountryDataManager{}, err
	}

	return &CountryDataManager{
		PathToData:  pathToData,
		CountryData: data,
		Logger:      log,
	}, nil
}

func (manager *CountryDataManager) GetCountryData(countryName string) (Country, error) {
	if countryName == "" {
		return Country{}, fmt.Errorf("country argument is required")
	}

	manager.Logger.Info(fmt.Sprintf("Fetching population data for %s...\n", countryName))

	for _, c := range manager.CountryData {
		if c.Name == countryName {
			return c, nil
		}
	}

	for _, c := range manager.CountryData {
		if c.Code == countryName {
			return c, nil
		}
	}

	return Country{}, fmt.Errorf("country %s not found", countryName)
}

func (country Country) Print(w io.Writer) {
	fmt.Fprintf(w, "Country: %s\n", country.Name)
	fmt.Fprintf(w, "Code: %s\n", country.Code)
	fmt.Fprintf(w, "Population Rank: %s\n", country.PopulationRank)
	fmt.Fprintf(w, "Capital: %s\n", country.Capital)
	fmt.Fprintf(w, "Continent: %s\n", country.Continent)
	fmt.Fprintf(w, "Population (2022): %d\n", country.Population2022)
	fmt.Fprintf(w, "Population (2020): %d\n", country.Population2020)
	fmt.Fprintf(w, "Population (2010): %d\n", country.Population2010)
	fmt.Fprintf(w, "Population (2000): %d\n", country.Population2000)
	fmt.Fprintf(w, "Area (km²): %d\n", country.Area)
	fmt.Fprintf(w, "Population Density (per km²): %.2f\n", country.PopulationDensity)
}
