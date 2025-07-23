package data

import (
	"bufio"
	"testing"
	"world-pop/internal/common/logger"
)

func getMockLogger(t *testing.T) (*logger.Logger, *bufio.Scanner) {
	logger, reader, err := logger.MockLogger()
	if err != nil {
		t.Fatalf("Failed to create mock logger: %v", err)
	}
	return &logger, bufio.NewScanner(reader)
}

func getMockData() []Country {
	return []Country{
		{
			Name:           "Foo",
			Code:           "F00",
			PopulationRank: "3",
			Capital:        "Foo City",
			Continent:      "Foo Land",
			Population2022: 1000000,
			Population2020: 950000,
			Population2010: 900000,
			Population2000: 850000,
			Area:           10000,
		}, {
			Name:           "Bar",
			Code:           "B00",
			PopulationRank: "2",
			Capital:        "Bar City",
			Continent:      "Bar Land",
			Population2022: 2000000,
			Population2020: 1900000,
			Population2010: 1800000,
			Population2000: 1700000,
			Area:           20000,
		}, {
			Name:           "Baz",
			Code:           "B0Z",
			PopulationRank: "1",
			Capital:        "Bar City",
			Continent:      "Bar Land",
			Population2022: 3000000,
			Population2020: 2900000,
			Population2010: 2800000,
			Population2000: 2700000,
			Area:           30000,
		},
	}
}

func TestGetCountryData(t *testing.T) {
	//Arrange
	mockLogger, _ := getMockLogger(t)

	// Mocking the data manager
	dataManager := &CountryDataManager{
		PathToData:  "test_data.csv",
		CountryData: getMockData(),
		Logger:      *mockLogger,
	}

	var tests = []struct {
		name           string
		input          string
		shouldError    bool
		wantOutput     string
		wantPopulation int
	}{
		// the table itself
		{"Foo country name seach succeeds with population", "Foo", false, "Foo", 1000000},
		{"Baz country code seach succeeds with population", "B0Z", false, "Baz", 3000000},
		{"Qux country name seach fails with no data", "Qux", true, "country Qux not found", 0},
		{"Error when no input provided", "", true, "country argument is required", 0},
	}

	for _, testRun := range tests {
		t.Run(testRun.name, func(t *testing.T) {
			// Act
			got, err := dataManager.GetCountryData(testRun.input)
			if testRun.shouldError {
				// Assert
				if err == nil || err.Error() != testRun.wantOutput {
					t.Errorf("GetCountryData(%s) returned error: %v; expected error %s", testRun.input, err, testRun.wantOutput)
					return
				}
			} else {
				if got.Name != testRun.wantOutput {
					//Assert
					t.Errorf("got %s for country name, want %s", got.Name, testRun.wantOutput)
				}
				if got.Population2022 != testRun.wantPopulation {
					//Assert
					t.Errorf("got %d for country population, want %d", got.Population2022, testRun.wantPopulation)
				}
			}
		})
	}
}
