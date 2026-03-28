package characterize

import (
	"testing"

	"github.com/kaiiorg/nws-demo-project/internal/config"

	"github.com/test-go/testify/require"
)

type testCharacterizeCase struct {
	Name           string
	Temp           int
	ForecastConfig *config.Forecast
	ExpectPanic    bool
	ExpectedResult string
}

var testCharacterizeCases = []testCharacterizeCase{{
	Name:           "Config is nil",
	ForecastConfig: nil,
	ExpectPanic:    true,
}, {
	Name: "0 degrees is cold when expected",
	Temp: 0,
	ForecastConfig: &config.Forecast{
		Hot:  90,
		Cold: 60,
	},
	ExpectedResult: "cold",
}, {
	Name: "59 degrees is cold when expected",
	Temp: 59,
	ForecastConfig: &config.Forecast{
		Hot:  90,
		Cold: 60,
	},
	ExpectedResult: "cold",
}, {
	Name: "60 degrees is moderate when expected",
	Temp: 60,
	ForecastConfig: &config.Forecast{
		Hot:  90,
		Cold: 60,
	},
	ExpectedResult: "moderate",
}, {
	Name: "90 degrees is moderate when expected",
	Temp: 90,
	ForecastConfig: &config.Forecast{
		Hot:  90,
		Cold: 60,
	},
	ExpectedResult: "moderate",
}, {
	Name: "91 degrees is hot when expected",
	Temp: 91,
	ForecastConfig: &config.Forecast{
		Hot:  90,
		Cold: 60,
	},
	ExpectedResult: "hot",
}, {
	Name: "100 degrees is hot when expected",
	Temp: 100,
	ForecastConfig: &config.Forecast{
		Hot:  90,
		Cold: 60,
	},
	ExpectedResult: "hot",
},
}

func TestCharacterize(t *testing.T) {
	ch := &Characterize{}

	for _, testCase := range testCharacterizeCases {
		result := ""

		if testCase.ExpectPanic {
			require.Panics(
				t,
				func() {
					result = ch.Characterize(testCase.Temp, testCase.ForecastConfig)
				},
			)
		} else {
			require.NotPanics(
				t,
				func() {
					result = ch.Characterize(testCase.Temp, testCase.ForecastConfig)
				},
			)
		}

		require.Equal(t, testCase.ExpectedResult, result)
	}
}
