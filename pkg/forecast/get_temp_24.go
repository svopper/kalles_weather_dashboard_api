package forecast

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/models"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/util"
)

func generateTemperatureUri(fromDate, toDate time.Time) string {
	// point := url.QueryEscape("POINT(12.585,55.662)")
	// uri := fmt.Sprintf(
	// 	"https://dmigw.govcloud.dk/v1/forecastedr/collections/harmonie_dini_sf/position?coords="+point+"&crs=crs84&parameter-name=temperature-2m&f=GeoJSON&datetime=2025-02-08T16:00:00Z/2025-02-09T15:00:00Z&api-key=%s",
	// 	// util.FormatDate(fromDate),
	// 	// util.FormatDate(toDate),
	// 	util.GetEnvVariable("DMI_FORECAST_EDR_API_KEY"),
	// )
	now := time.Now()
	rounded := now.Truncate(time.Hour)
	_, offset := now.Zone()
	utcAdjusted := rounded.Add(time.Duration(offset) * time.Second)
	startDate := utcAdjusted.UTC().Format(time.RFC3339)
	endDate := utcAdjusted.UTC().Add(23 * time.Hour).Format(time.RFC3339)

	return "https://dmigw.govcloud.dk/v1/forecastedr/collections/harmonie_dini_sf/position?coords=POINT%2812.585%2055.662%29&crs=crs84&parameter-name=temperature-2m&f=GeoJSON&datetime=" + startDate + "/" + endDate + "&api-key=" + util.GetEnvVariable("DMI_FORECAST_EDR_API_KEY")
	// return uri
}

func getTemperatureForecast(from, to time.Time) models.DMIObservation {
	uri := generateTemperatureUri(from, to)
	request := util.BuildRequest(uri)
	response := util.DoRequest(request)
	body, err := io.ReadAll(response.Body)

	if err != nil {

		panic(err)
	}

	temperatureForecast, err := util.UnmarshalDMIObservation(body)

	if err != nil {
		panic(err)
	}
	return temperatureForecast
}

func (h handler) GetTemperatureForecast(c *gin.Context) {
	forecast := getTemperatureForecast(time.Now(), time.Now().Add(24*time.Hour))
	viewModel := []models.TemperatureForecast{}

	for _, feature := range forecast.Features {
		viewModel = append(viewModel, models.TemperatureForecast{
			Date:        feature.Properties.Step,
			Temperature: feature.Properties.Temperature - 273.15,
		})
	}

	c.JSON(200, viewModel)
}
